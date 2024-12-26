package model

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	pb "grpcservice/proto/model" // 更新为提供的路径
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

// 初始化 bufconn
func initListener() {
	lis = bufconn.Listen(bufSize)
}

// mock 实现的服务
type mockModelServiceServer struct {
	pb.UnimplementedModelServiceServer
}

func (s *mockModelServiceServer) GetModels(ctx context.Context, req *pb.EmptyRequest) (*pb.ModelListResponse, error) {
	return &pb.ModelListResponse{
		Models: []*pb.Model{
			{Id: "1", Name: "Model1"},
			{Id: "2", Name: "Model2"},
		},
	}, nil
}

func (s *mockModelServiceServer) AddToNetwork(ctx context.Context, req *pb.NetworkRequest) (*pb.Response, error) {
	return &pb.Response{Message: "Model added to network successfully"}, nil
}

func (s *mockModelServiceServer) RemoveFromNetwork(ctx context.Context, req *pb.RemoveRequest) (*pb.Response, error) {
	return &pb.Response{Message: "Model removed from network successfully"}, nil
}

type mockClient struct {
	conn   *grpc.ClientConn
	client pb.ModelServiceClient
}

// Close 关闭连接
func (c *mockClient) Close() {
	if c.conn != nil {
		err := c.conn.Close()
		if err != nil {
			log.Println("Failed to close connection:", err)
		}
	}
}
func (c *mockClient) GetModels(ctx context.Context) (*pb.ModelListResponse, error) {
	return c.client.GetModels(ctx, &pb.EmptyRequest{})
}

// 启动 mock gRPC 服务
func startMockServer() {
	server := grpc.NewServer()
	pb.RegisterModelServiceServer(server, &mockModelServiceServer{})
	go func() {
		if err := server.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

// 自定义拨号器，用于 bufconn
func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func NewTestClient(bufnet string) (*mockClient, error) {
	opts := []grpc.DialOption{
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	resolver.SetDefaultScheme("passthrough")
	conn, err := grpc.NewClient(bufnet, opts...)
	if err != nil {
		return nil, err
	}
	client := pb.NewModelServiceClient(conn)
	return &mockClient{conn: conn, client: client}, nil
}
func TestClient_GetModels(t *testing.T) {
	initListener()
	startMockServer()
	// 创建客户端

	client, err := NewTestClient("bufnet")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// 调用 GetModels
	ctx := context.Background()
	resp, err := client.GetModels(ctx)
	if err != nil {
		t.Fatalf("GetModels failed: %v", err)
	}

	// 验证结果
	if len(resp.Models) != 2 {
		t.Errorf("Expected 2 models, got %d", len(resp.Models))
	}
	if resp.Models[0].Name != "Model1" || resp.Models[1].Name != "Model2" {
		t.Errorf("Unexpected models: %+v", resp.Models)
	}
}
