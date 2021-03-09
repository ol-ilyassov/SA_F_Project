package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/ol-ilyassov/final/notifier/notifypb"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/smtp"
	"os"
)

type Server struct {
	notifypb.UnimplementedNotifierServiceServer
}

func main() {
	_ = godotenv.Load("globals.env")
	port := "60055"
	l, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	notifypb.RegisterNotifierServiceServer(s, &Server{})
	log.Println("Server is running on port: " + port)
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

func (s *Server) ArticleCreationNotify(ctx context.Context, req *notifypb.ArticleCreationRequest) (*notifypb.ArticleCreationResponse, error) {
	log.Printf("ArticleCreationNotify function was invoked with %v \n", req)

	from := os.Getenv("GMAIL_LOGIN")
	pass := os.Getenv("GMAIL_PASSWORD")

	recipient := req.GetAddress()

	body := " - Title: " + req.GetTitle() + "\n" +
		" - When: " + req.GetTime() + "\n\n\n" +
		" Enjoy reading!"
	message := "From: " + from + "\n" +
		"To: " + recipient + "\n" +
		"Subject: New Article is Published!\n\n" + body

	to := []string{
		recipient,
	}
	msg := []byte(message)
	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)

	res := &notifypb.ArticleCreationResponse{
		Result: "Success",
	}
	if err != nil {
		res.Result = "Fail"
		log.Fatalf("Error while processing SendNotification: %v", err.Error())
	}

	return res, err
}

func (s *Server) UserCreationNotify(ctx context.Context, req *notifypb.UserCreationRequest) (*notifypb.UserCreationResponse, error) {
	log.Printf("UserCreationNotify function was invoked with %v \n", req)

	from := os.Getenv("GMAIL_LOGIN")
	pass := os.Getenv("GMAIL_PASSWORD")

	recipient := req.GetAddress()

	body := " - New Account connected with " + recipient + " email have been created!\n" +
		" - When: " + req.GetTime() + "\n\n\n" +
		" Enjoy your Account on our Web Portal - Articles' Hub!"
	message := "From: " + from + "\n" +
		"To: " + recipient + "\n" +
		"Subject: New Account is Created!\n\n" + body

	to := []string{
		recipient,
	}
	msg := []byte(message)
	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)

	res := &notifypb.UserCreationResponse{
		Result: "Success",
	}
	if err != nil {
		res.Result = "Fail"
		log.Fatalf("Error while processing SendNotification: %v", err.Error())
	}

	return res, err

	return res, err
}
