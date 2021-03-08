// In terminal: go run ./cmd/web -addr=":4000"
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ol-ilyassov/final/auth/authpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type Server struct {
	authpb.UnimplementedAuthServiceServer
}

func main() {
	//secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()

	// Connection to DB
	dns := flag.String("dns", "postgres://postgres:1@localhost:5432/snippetbox", "Postgre data source name")
	db, err := openDB(*dns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	m = &UserModel{
		DB: db,
	}

	// Server Creation
	port := "50051" // Port of Article_DB Microservice
	l, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &Server{})
	log.Println("Server is running on port: " + port)
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
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

//func (s *Server) CreateUser(ctx context.Context, req *authpb.AuthUserRequest) (*authpb.AuthUserResponse, error) {
//	log.Printf("GetArticle function was invoked with %v \n", req)
//
//	articleRes, err := m.Get(int(req.GetId()))
//
//	res := &articlepb.GetArticleResponse{
//		Article: &articlepb.Article{
//			Id:      int32(articleRes.ID),
//			Title:   articleRes.Title,
//			Content: articleRes.Content,
//			Created: articleRes.Created.String(),
//			Expires: articleRes.Expires.String(),
//		},
//		Result: "Success",
//	}
//
//	return res, err
//}
