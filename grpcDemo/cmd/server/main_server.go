package main

import (
	"log"

	"github.com/user/demo/internal/server"
)

func main() {
	log.Println("Starting gRPC server...")
	server.StartServer("50051")
}
