package transaction

import (
	"context"
	pb "grpcservice/proto/transaction"
)

type Server struct {
	pb.UnimplementedTransactionServiceServer
}

func (s *Server) PerformTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.Response, error) {
	// Logic to perform the transaction
	return &pb.Response{Message: "Transaction successful", Success: true}, nil
}
