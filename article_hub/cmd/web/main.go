// In terminal: go run ./cmd/web -addr=":4000"
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golangcollege/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ol-ilyassov/final/article_hub/articlepb"
	"github.com/ol-ilyassov/final/article_hub/notifypb"
	"github.com/ol-ilyassov/final/article_hub/pkg/models/postgres"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type contextKey string

const contextKeyIsAuthenticated = contextKey("isAuthenticated")

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      articlepb.ArticlesServiceClient
	notifier      notifypb.NotifierServiceClient
	templateCache map[string]*template.Template
	users         *postgres.UserModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dns := flag.String("dns", "postgres://postgres:1@localhost:5432/snippetbox", "Postgre data source name")
	db, err := openDB(*dns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	port := "50051" // Port of Article_DB Microservice
	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	articlesDBService := articlepb.NewArticlesServiceClient(conn)

	// Notifier Service
	port2 := "50055"
	conn2, err := grpc.Dial("localhost:"+port2, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	notifierService := notifypb.NewNotifierServiceClient(conn2)

	templateCache, err := newTemplateCache("ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	//session.Secure = true
	//session.SameSite = http.SameSiteStrictMode

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		snippets:      articlesDBService,
		notifier:      notifierService,
		templateCache: templateCache,
		users:         &postgres.UserModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting a server on port%v", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dns string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(context.Background(), dns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
		return nil, err
	}
	return conn, nil
}
