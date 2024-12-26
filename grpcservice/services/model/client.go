package model

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpcservice/proto/model" // 更新为提供的路径
	"log"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.ModelServiceClient
}

// NewClient 创建一个新的客户端连接到指定的 gRPC 服务器
func NewClient(address string) (*Client, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewModelServiceClient(conn)
	return &Client{conn: conn, client: client}, nil
}

// Close 关闭连接
func (c *Client) Close() {
	if c.conn != nil {
		err := c.conn.Close()
		if err != nil {
			log.Println("Failed to close connection:", err)
		}
	}
}

// GetModels 查询模型列表
func (c *Client) GetModels(ctx context.Context) (*pb.ModelListResponse, error) {
	return c.client.GetModels(ctx, &pb.EmptyRequest{})
}

// AddToNetwork 加入网络
func (c *Client) AddToNetwork(ctx context.Context, modelID string, networkID string) (*pb.Response, error) {
	req := &pb.NetworkRequest{ModelId: modelID, NetworkId: networkID}
	return c.client.AddToNetwork(ctx, req)
}

// RemoveFromNetwork 退网
func (c *Client) RemoveFromNetwork(ctx context.Context, modelID string) (*pb.Response, error) {
	req := &pb.RemoveRequest{ModelId: modelID}
	return c.client.RemoveFromNetwork(ctx, req)
}
