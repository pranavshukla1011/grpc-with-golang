package greet_client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-with-golang/greetpb"
	"log"
)

func UseConn() error {
	log.Println("Establishing connection to server...")

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Could not connect", err)
		return err
	}
	defer conn.Close()
	defer log.Println("Client Connection SUCCESSFULLY Closed")
	c := greetpb.NewGreetServiceClient(conn)
	log.Println("GreetServiceClient : ", c)

	req := &greetpb.GreetRequest{
		Request: &greetpb.Greeting{
			FirstName: "Pranav",
			LastName:  "Shukla",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling greet rpc:", err)
	}
	log.Println(res)
	return nil
}
