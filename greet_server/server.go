package greet_server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-with-golang/greetpb"
	"log"
	"net"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (c *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetRequest().FirstName
	result := "Hello" + firstName
	response := &greetpb.GreetResponse{
		Result: result,
	}
	return response, nil
}

func StartServer(port string) error {
	fmt.Println("Starting port ...")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error while starting connection", err)
		return err
	}

	log.Println("Listener successfully created for port ", port)

	log.Println("Creating grpc server...")
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve", err)
	}
	return nil
}
