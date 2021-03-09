// In terminal: go run ./cmd/web -addr=":4000"
package main

import (
	"context"
	"errors"
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
	port := "60059" // Port of Article_DB Microservice
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

func (s *Server) CreateUser(ctx context.Context, req *authpb.CreateUserRequest) (*authpb.CreateUserResponse, error) {
	log.Printf("CreateUser function was invoked with %v \n", req.GetUser().GetEmail())

	errSql := m.Insert(req.GetUser().GetName(), req.GetUser().GetEmail(), req.GetUser().GetPassword())

	res := &authpb.CreateUserResponse{
		Result: "Success",
		Status: true,
	}

	if errSql != nil {
		res.Status = false
		if errors.Is(errSql, ErrDuplicateEmail) {
			res.Result = ErrDuplicateEmail.Error()
		} else {
			res.Result = errSql.Error()
		}
	}

	return res, nil
}

func (s *Server) AuthUser(ctx context.Context, req *authpb.AuthUserRequest) (*authpb.AuthUserResponse, error) {
	log.Printf("AuthUser function was invoked with %s \n", req.GetUser().GetEmail())

	id, errSql := m.Authenticate(req.GetUser().GetEmail(), req.GetUser().GetPassword())

	res := &authpb.AuthUserResponse{
		Result: "Success",
		Status: true,
	}

	if errSql != nil {
		res.Id = 0
		res.Status = false
		if errors.Is(errSql, ErrInvalidCredentials) {
			res.Result = ErrInvalidCredentials.Error()
		} else if errors.Is(errSql, ErrNoRecord) {
			res.Result = ErrNoRecord.Error()
		}
	} else {
		res.Id = int32(id)
	}

	log.Printf("Result: %s", res.GetResult())

	return res, nil
}

func (s *Server) GetUser(ctx context.Context, req *authpb.GetUserRequest) (*authpb.GetUserResponse, error) {
	log.Printf("GetUser function was invoked with id: %s \n", req.GetId())

	resSql, errSql := m.Get(int(req.GetId()))

	res := &authpb.GetUserResponse{
		Result: "Success",
		Status: true,
		User:   &authpb.User{},
	}

	if errSql != nil {
		res.Status = false
		if errors.Is(errSql, ErrNoRecord) {
			res.Result = ErrNoRecord.Error()
		}
	} else {
		res.GetUser().Id = int32(resSql.ID)
		res.GetUser().Name = resSql.Name
		res.GetUser().Email = resSql.Email
		res.GetUser().Created = resSql.Created.String()
		res.GetUser().Active = resSql.Active
	}

	return res, nil
}
