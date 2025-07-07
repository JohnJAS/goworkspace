package main

import (
	"fmt"
	"log"
	"os"

	"github.com/user/demo/internal/client"
)

func main() {
	os.Args = []string{"main_client.go", "127.0.0.1:50051"}
	if len(os.Args) < 2 {
		fmt.Println("Usage: client <server_address>")
		return
	}

	serverAddr := os.Args[1]
	cli, err := client.NewDemoClient(serverAddr)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer cli.Close()

	// 测试SimpleRPC
	resp, err := cli.SimpleRPC("test simple")
	if err != nil {
		log.Printf("SimpleRPC failed: %v", err)
	} else {
		log.Printf("SimpleRPC response: %s", resp)
	}

	// 测试ServerStreamRPC
	streamResp, err := cli.ServerStreamRPC("test stream")
	if err != nil {
		log.Printf("ServerStreamRPC failed: %v", err)
	} else {
		log.Println("ServerStreamRPC responses:")
		for _, r := range streamResp {
			log.Printf("- %s", r)
		}
	}

	// 测试ClientStreamRPC
	clientStreamResp, err := cli.ClientStreamRPC([]string{"msg1", "msg2", "msg3"})
	if err != nil {
		log.Printf("ClientStreamRPC failed: %v", err)
	} else {
		log.Printf("ClientStreamRPC response: %s", clientStreamResp)
	}

	// 测试BidirectionalStreamRPC
	biResp, err := cli.BidirectionalStreamRPC([]string{"msgA", "msgB", "msgC"})
	if err != nil {
		log.Printf("BidirectionalStreamRPC failed: %v", err)
	} else {
		log.Println("BidirectionalStreamRPC responses:")
		for _, r := range biResp {
			log.Printf("- %s", r)
		}
	}
}
