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
	port := "50055"
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

func (s *Server) SendNotification(ctx context.Context, req *notifypb.NotifyRequest) (*notifypb.NotifyResponse, error) {
	log.Printf("SendNotification function was invoked with %v \n", req)

	from := os.Getenv("GMAIL_LOGIN")
	pass := os.Getenv("GMAIL_PASSWORD")

	recipient := req.GetEmail().GetAddress()

	body := " - Title: " + req.GetEmail().GetTitle() + "\n" +
		" - When: " + req.GetEmail().GetTime() + "\n\n\n" +
		" Enjoy reading!"

	message := "From: " + from + "\n" +
		"To: " + recipient + "\n" +
		"Subject: New Article is Published!\n\n" +
		body

	to := []string{
		recipient,
	}
	msg := []byte(message)

	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)

	res := &notifypb.NotifyResponse{
		Result: "Success",
	}

	if err != nil {
		res.Result = "Fail"
		log.Fatalf("Error while processing SendNotification: %v", err.Error())
	}

	return res, err
}
