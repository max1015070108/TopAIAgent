package transaction

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpcservice/proto/transaction"
	"log"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.TransactionServiceClient
}

// NewClient 创建新的交易服务客户端
func NewClient(address string) (*Client, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewTransactionServiceClient(conn)
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

// PerformTransaction 执行交易
func (c *Client) PerformTransaction(ctx context.Context, modelID string, amount string) (*pb.Response, error) {
	req := &pb.TransactionRequest{ModelId: modelID, Amount: amount}
	return c.client.PerformTransaction(ctx, req)
}
