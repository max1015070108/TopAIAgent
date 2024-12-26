package network

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	pb "grpcservice/proto/network"
)

var bufConnNetwork *bufconn.Listener

func init() {
	bufConnNetwork = bufconn.Listen(1024 * 1024)
}

type mockNetworkServiceServer struct {
	pb.UnimplementedNetworkServiceServer
}

func NewMockNetworkServer(s *mockNetworkServiceServer) {
	go func() {
		grpcServer := grpc.NewServer()
		pb.RegisterNetworkServiceServer(grpcServer, s)
		if err := grpcServer.Serve(bufConnNetwork); err != nil {
			panic(err)
		}
	}()
}

// Add test methods for the client...
