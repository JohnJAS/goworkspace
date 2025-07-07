package client

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/user/demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type mockDemoServiceClient struct {
	mock.Mock
	proto.DemoServiceClient
}

func (m *mockDemoServiceClient) SimpleRPC(ctx context.Context, in *proto.SimpleRequest, opts ...grpc.CallOption) (*proto.SimpleResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*proto.SimpleResponse), args.Error(1)
}

func TestSimpleRPC(t *testing.T) {
	mockClient := new(mockDemoServiceClient)
	client := &DemoClient{
		client: mockClient,
	}

	expectedResponse := &proto.SimpleResponse{ResponseData: "test response"}
	mockClient.On("SimpleRPC", mock.Anything, &proto.SimpleRequest{RequestData: "test"}).Return(expectedResponse, nil)

	resp, err := client.SimpleRPC("test")
	assert.NoError(t, err)
	assert.Equal(t, "test response", resp)
	mockClient.AssertExpectations(t)
}

func (m *mockDemoServiceClient) ServerStreamRPC(ctx context.Context, in *proto.StreamRequest, opts ...grpc.CallOption) (proto.DemoService_ServerStreamRPCClient, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(proto.DemoService_ServerStreamRPCClient), args.Error(1)
}

func (m *mockDemoServiceClient) ClientStreamRPC(ctx context.Context, opts ...grpc.CallOption) (proto.DemoService_ClientStreamRPCClient, error) {
	args := m.Called(ctx)
	return args.Get(0).(proto.DemoService_ClientStreamRPCClient), args.Error(1)
}

func (m *mockDemoServiceClient) BidirectionalStreamRPC(ctx context.Context, opts ...grpc.CallOption) (proto.DemoService_BidirectionalStreamRPCClient, error) {
	args := m.Called(ctx)
	return args.Get(0).(proto.DemoService_BidirectionalStreamRPCClient), args.Error(1)
}

type mockStreamClient struct {
	mock.Mock
	grpc.ClientStream
}

func (m *mockStreamClient) Recv() (*proto.StreamResponse, error) {
	args := m.Called()
	return args.Get(0).(*proto.StreamResponse), args.Error(1)
}

func (m *mockStreamClient) Send(request *proto.StreamRequest) error {
	args := m.Called(request)
	return args.Error(0)
}

func (m *mockStreamClient) CloseSend() error {
	args := m.Called()
	return args.Error(0)
}

func (m *mockStreamClient) CloseAndRecv() (*proto.StreamResponse, error) {
	args := m.Called()
	return args.Get(0).(*proto.StreamResponse), args.Error(1)
}

func TestServerStreamRPC(t *testing.T) {
	mockClient := new(mockDemoServiceClient)
	mockStream := new(mockStreamClient)
	client := &DemoClient{
		client: mockClient,
	}

	responses := []*proto.StreamResponse{
		{ResponseData: "response 1"},
		{ResponseData: "response 2"},
	}

	mockClient.On("ServerStreamRPC", mock.Anything, &proto.StreamRequest{RequestData: "test"}).
		Return(mockStream, nil)

	mockStream.On("Recv").Return(responses[0], nil).Once()
	mockStream.On("Recv").Return(responses[1], nil).Once()
	mockStream.On("Recv").Return(&proto.StreamResponse{}, io.EOF).Once()

	result, err := client.ServerStreamRPC("test")
	assert.NoError(t, err)
	assert.Equal(t, []string{"response 1", "response 2"}, result)

	mockClient.AssertExpectations(t)
	mockStream.AssertExpectations(t)
}

func TestClientStreamRPC(t *testing.T) {
	mockClient := new(mockDemoServiceClient)
	mockStream := new(mockStreamClient)
	client := &DemoClient{
		client: mockClient,
	}

	response := &proto.StreamResponse{ResponseData: "final response"}

	mockClient.On("ClientStreamRPC", mock.Anything).
		Return(mockStream, nil)

	mockStream.On("Send", &proto.StreamRequest{RequestData: "msg1"}).Return(nil).Once()
	mockStream.On("Send", &proto.StreamRequest{RequestData: "msg2"}).Return(nil).Once()
	mockStream.On("CloseAndRecv").Return(response, nil).Once()

	result, err := client.ClientStreamRPC([]string{"msg1", "msg2"})
	assert.NoError(t, err)
	assert.Equal(t, "final response", result)

	mockClient.AssertExpectations(t)
	mockStream.AssertExpectations(t)
}

func TestBidirectionalStreamRPC(t *testing.T) {
	mockClient := new(mockDemoServiceClient)
	mockStream := new(mockStreamClient)
	client := &DemoClient{
		client: mockClient,
	}

	responses := []*proto.StreamResponse{
		{ResponseData: "response A"},
		{ResponseData: "response B"},
	}

	mockClient.On("BidirectionalStreamRPC", mock.Anything).
		Return(mockStream, nil)

	mockStream.On("Send", &proto.StreamRequest{RequestData: "msgA"}).Return(nil).Once()
	mockStream.On("Send", &proto.StreamRequest{RequestData: "msgB"}).Return(nil).Once()
	mockStream.On("CloseSend").Return(nil).Once()

	mockStream.On("Recv").Return(responses[0], nil).Once()
	mockStream.On("Recv").Return(responses[1], nil).Once()
	mockStream.On("Recv").Return(&proto.StreamResponse{}, io.EOF).Once()

	result, err := client.BidirectionalStreamRPC([]string{"msgA", "msgB"})
	assert.NoError(t, err)
	assert.Equal(t, []string{"response A", "response B"}, result)

	mockClient.AssertExpectations(t)
	mockStream.AssertExpectations(t)
}

func TestNewDemoClient(t *testing.T) {
	// 测试连接创建
	client, err := NewDemoClient("localhost:50051")
	assert.NoError(t, err)
	assert.NotNil(t, client)
	defer client.Close()
}

func TestClose(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	assert.NoError(t, err)

	client := &DemoClient{
		conn: conn,
	}

	// 测试Close方法
	client.Close()

	// 验证连接是否已关闭
	state := conn.GetState()
	assert.Equal(t, state, connectivity.Shutdown) // 使用connectivity包中的状态常量
}
