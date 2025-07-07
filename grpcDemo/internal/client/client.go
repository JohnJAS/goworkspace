package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/user/demo/proto"
	"google.golang.org/grpc"
)

type DemoClient struct {
	conn   *grpc.ClientConn
	client proto.DemoServiceClient
}

func NewDemoClient(serverAddr string) (*DemoClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %v", err)
	}

	return &DemoClient{
		conn:   conn,
		client: proto.NewDemoServiceClient(conn),
	}, nil
}

func (c *DemoClient) Close() {
	c.conn.Close()
}

// SimpleRPC调用
func (c *DemoClient) SimpleRPC(data string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.client.SimpleRPC(ctx, &proto.SimpleRequest{
		RequestData: data,
	})
	if err != nil {
		return "", fmt.Errorf("SimpleRPC failed: %v", err)
	}

	return resp.GetResponseData(), nil
}

// ServerStreamRPC调用
func (c *DemoClient) ServerStreamRPC(data string) ([]string, error) {
	stream, err := c.client.ServerStreamRPC(context.Background(), &proto.StreamRequest{
		RequestData: data,
	})
	if err != nil {
		return nil, fmt.Errorf("ServerStreamRPC failed: %v", err)
	}

	var responses []string
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error receiving stream: %v", err)
		}
		responses = append(responses, resp.GetResponseData())
	}

	return responses, nil
}

// ClientStreamRPC调用
func (c *DemoClient) ClientStreamRPC(data []string) (string, error) {
	stream, err := c.client.ClientStreamRPC(context.Background())
	if err != nil {
		return "", fmt.Errorf("ClientStreamRPC failed: %v", err)
	}

	for _, d := range data {
		if err := stream.Send(&proto.StreamRequest{
			RequestData: d,
		}); err != nil {
			return "", fmt.Errorf("error sending stream: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return "", fmt.Errorf("error closing stream: %v", err)
	}

	return resp.GetResponseData(), nil
}

// BidirectionalStreamRPC调用
func (c *DemoClient) BidirectionalStreamRPC(data []string) ([]string, error) {
	stream, err := c.client.BidirectionalStreamRPC(context.Background())
	if err != nil {
		return nil, fmt.Errorf("BidirectionalStreamRPC failed: %v", err)
	}

	waitc := make(chan struct{})
	var responses []string

	// 接收响应
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Printf("Error receiving response: %v", err)
				return
			}
			responses = append(responses, resp.GetResponseData())
		}
	}()

	// 发送请求
	for _, d := range data {
		if err := stream.Send(&proto.StreamRequest{
			RequestData: d,
		}); err != nil {
			return nil, fmt.Errorf("error sending request: %v", err)
		}
	}

	stream.CloseSend()
	<-waitc

	return responses, nil
}
