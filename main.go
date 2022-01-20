package main

import (
	"grpc-with-golang/greet_server"
	"log"
)

func main() {
	port := "0.0.0.0:50051"
	err := greet_server.StartServer(port)

	if err != nil {
		log.Println("Server Creation UNSUCCESSFUL")
	}
	log.Println("Server Creation SUCCESSFUL")
}
