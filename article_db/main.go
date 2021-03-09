package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ol-ilyassov/final/article_db/articlepb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type Server struct {
	articlepb.UnimplementedArticlesServiceServer
}

func main() {
	// Connection to DB
	dns := flag.String("dns", "postgres://postgres:1@localhost:5432/articleshub", "Postgre data source name")
	db, err := openDB(*dns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	m = &ArticleModel{
		DB: db,
	}

	// Server Creation
	port := "60051" // Port of Article_DB Microservice
	l, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	articlepb.RegisterArticlesServiceServer(s, &Server{})
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

func (s *Server) GetArticle(ctx context.Context, req *articlepb.GetArticleRequest) (*articlepb.GetArticleResponse, error) {
	log.Printf("GetArticle function was invoked with %v \n", req)

	articleRes, err := m.Get(int(req.GetId()))

	res := &articlepb.GetArticleResponse{
		Article: &articlepb.Article{
			Id:       int32(articleRes.ID),
			Authorid: int32(articleRes.AuthorID),
			Title:    articleRes.Title,
			Content:  articleRes.Content,
			Created:  articleRes.Created.String(),
			Expires:  articleRes.Expires.String(),
		},
		Result: "Success",
	}

	return res, err
}

func (s *Server) GetArticles(req *articlepb.GetArticlesRequest, stream articlepb.ArticlesService_GetArticlesServer) error {
	log.Printf("GetArticles function was invoked\n")

	arr, err := m.Latest()
	if err != nil {
		return err
	}

	for _, articleRes := range arr {
		res := &articlepb.GetArticlesResponse{
			Article: &articlepb.Article{
				Id:       int32(articleRes.ID),
				Authorid: int32(articleRes.AuthorID),
				Title:    articleRes.Title,
				Content:  articleRes.Content,
				Created:  articleRes.Created.String(),
				Expires:  articleRes.Expires.String(),
			},
		}

		if err := stream.Send(res); err != nil {
			log.Fatalf("Error while sending GetArticles responses: %v", err.Error())
		}
	}

	return nil
}

func (s *Server) InsertArticle(ctx context.Context, req *articlepb.InsertArticleRequest) (*articlepb.InsertArticleResponse, error) {
	log.Printf("InsertArticle function was invoked with %v \n", req)

	id, err := m.Insert(req.GetArticle().GetTitle(), req.GetArticle().GetContent(), req.GetArticle().GetExpires(), int(req.GetArticle().GetAuthorid()))

	res := &articlepb.InsertArticleResponse{
		Id:     int32(id),
		Result: "Success",
	}

	if err != nil {
		res.Result = "Fail"
		log.Fatalf("Error while processing InsertArticle: %v", err.Error())
	}

	return res, nil
}

func (s *Server) DeleteArticle(ctx context.Context, req *articlepb.DeleteArticleRequest) (*articlepb.DeleteArticleResponse, error) {
	log.Printf("DeleteArticle function was invoked with %v \n", req)

	res := &articlepb.DeleteArticleResponse{
		Result: "Success",
	}

	err := m.Delete(int(req.GetId()))
	if err != nil {
		res.Result = "Fail"
		log.Fatalf("Error while processing DeleteArticle: %v", err.Error())
	}

	return res, nil
}

func (s *Server) SearchArticles(req *articlepb.SearchArticlesRequest, stream articlepb.ArticlesService_SearchArticlesServer) error {
	log.Printf("SearchArticles function was invoked\n")

	arr, err := m.Search(req.GetTitle())
	if err != nil {
		return err
	}

	for _, articleRes := range arr {
		res := &articlepb.SearchArticlesResponse{
			Article: &articlepb.Article{
				Id:       int32(articleRes.ID),
				Authorid: int32(articleRes.AuthorID),
				Title:    articleRes.Title,
				Content:  articleRes.Content,
				Created:  articleRes.Created.String(),
				Expires:  articleRes.Expires.String(),
			},
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("Error while sending SearchArticles responses: %v", err.Error())
		}
	}

	return nil
}

func (s *Server) EditArticle(ctx context.Context, req *articlepb.EditArticleRequest) (*articlepb.EditArticleResponse, error) {
	log.Printf("EditArticle function was invoked with %v \n", req)
	res := &articlepb.EditArticleResponse{
		Result: "Success",
	}

	err := m.Edit(req.GetArticle().GetTitle(), req.GetArticle().GetContent(), int(req.GetArticle().GetId()))
	if err != nil {
		res.Result = "Fail"
		log.Fatalf("Error while processing EditArticle: %v", err.Error())
	}

	return res, nil
}
