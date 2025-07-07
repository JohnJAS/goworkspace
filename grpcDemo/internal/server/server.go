package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/user/demo/proto"
	"google.golang.org/grpc"
)

type demoServer struct {
	proto.UnimplementedDemoServiceServer
}

// SimpleRPC实现
func (s *demoServer) SimpleRPC(ctx context.Context, req *proto.SimpleRequest) (*proto.SimpleResponse, error) {
	log.Printf("Received SimpleRPC request: %v", req.GetRequestData())
	return &proto.SimpleResponse{
		ResponseData: fmt.Sprintf("Processed: %s", req.GetRequestData()),
	}, nil
}

// ServerStreamRPC实现
func (s *demoServer) ServerStreamRPC(req *proto.StreamRequest, stream proto.DemoService_ServerStreamRPCServer) error {
	log.Printf("Received ServerStreamRPC request: %v", req.GetRequestData())

	for i := 0; i < 5; i++ {
		resp := &proto.StreamResponse{
			ResponseData: fmt.Sprintf("Stream %d: %s", i+1, req.GetRequestData()),
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}

// ClientStreamRPC实现
func (s *demoServer) ClientStreamRPC(stream proto.DemoService_ClientStreamRPCServer) error {
	var combinedData string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.StreamResponse{
				ResponseData: fmt.Sprintf("Processed: %s", combinedData),
			})
		}
		if err != nil {
			return err
		}
		combinedData += req.GetRequestData() + " "
	}
}

// BidirectionalStreamRPC实现
func (s *demoServer) BidirectionalStreamRPC(stream proto.DemoService_BidirectionalStreamRPCServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		resp := &proto.StreamResponse{
			ResponseData: fmt.Sprintf("Echo: %s", req.GetRequestData()),
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}

func StartServer(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterDemoServiceServer(grpcServer, &demoServer{})

	log.Printf("Server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
