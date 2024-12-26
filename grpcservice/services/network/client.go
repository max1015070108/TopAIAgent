package network

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpcservice/proto/network"
	"log"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.NetworkServiceClient
}

// NewClient 创建新的网络服务客户端
func NewClient(address string) (*Client, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewNetworkServiceClient(conn)
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

// GetNetworks 查询网络列表
func (c *Client) GetNetworks(ctx context.Context) (*pb.NetworkListResponse, error) {
	return c.client.GetNetworks(ctx, &pb.EmptyRequest{})
}
