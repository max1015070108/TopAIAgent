package network

import (
	"context"
	pb "grpcservice/proto/network"
)

type Server struct {
	pb.UnimplementedNetworkServiceServer
}

func (s *Server) GetNetworks(ctx context.Context, req *pb.EmptyRequest) (*pb.NetworkListResponse, error) {
	networks := []*pb.Network{
		{Id: "1", Name: "Network A"},
		{Id: "2", Name: "Network B"},
	}
	return &pb.NetworkListResponse{Networks: networks}, nil
}
