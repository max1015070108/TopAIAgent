package transaction

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	pb "grpcservice/proto/transaction"
)

var bufConnTransaction *bufconn.Listener

func init() {
	bufConnTransaction = bufconn.Listen(1024 * 1024)
}

type mockTransactionServiceServer struct {
	pb.UnimplementedTransactionServiceServer
}

func NewMockTransactionServer(s *mockTransactionServiceServer) {
	go func() {
		grpcServer := grpc.NewServer()
		pb.RegisterTransactionServiceServer(grpcServer, s)
		if err := grpcServer.Serve(bufConnTransaction); err != nil {
			panic(err)
		}
	}()
}

// Add test methods for the client...
