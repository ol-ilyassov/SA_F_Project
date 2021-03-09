// In terminal: go run ./cmd/web -addr=":4000"
package main

import (
	"flag"
	"github.com/golangcollege/sessions"
	"github.com/ol-ilyassov/final/article_hub/articlepb"
	"github.com/ol-ilyassov/final/article_hub/authpb"
	"github.com/ol-ilyassov/final/article_hub/notifypb"
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
	articles      articlepb.ArticlesServiceClient
	notifier      notifypb.NotifierServiceClient
	auth          authpb.AuthServiceClient
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()

	// Loggers init
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Article_DB Microservice
	port := "60051"
	conn1, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn1.Close()
	articlesDBService := articlepb.NewArticlesServiceClient(conn1)

	// Notifier Service
	port2 := "60055"
	conn2, err := grpc.Dial("localhost:"+port2, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn2.Close()
	notifierService := notifypb.NewNotifierServiceClient(conn2)

	// Auth Service
	port3 := "60059"
	conn3, err := grpc.Dial("localhost:"+port3, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn3.Close()
	authService := authpb.NewAuthServiceClient(conn3)

	// Template Cache init
	templateCache, err := newTemplateCache("ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Session init
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		articles:      articlesDBService,
		notifier:      notifierService,
		auth:          authService,
		templateCache: templateCache,
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
