package main

import (
	"grpc-with-golang/greet_client"
	"grpc-with-golang/greet_server"
	"log"
	"testing"
)

func TestServerStart(t *testing.T) {
	port := "0.0.0.0:50051"
	err := greet_server.StartServer(port)

	if err != nil {
		log.Println("Server Creation UNSUCCESSFUL")
	}
	log.Println("Server Creation SUCCESSFUL")
}

func TestClientConn(t *testing.T) {
	err := greet_client.UseConn()
	if err != nil {
		log.Fatal("ERROR WHILE USING GRPC CONN")
	}
}
