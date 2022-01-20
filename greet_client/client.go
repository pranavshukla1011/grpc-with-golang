package greet_client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-with-golang/greetpb"
	"log"
)

func CreateConn() error {
	log.Println("Establishing connection to server...")

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Could not connect", err)
		return err
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	log.Println("Client connection to grpc sever SUCCESSFUL")
	log.Println("GreetServiceClient : ", c)
	return nil
}
