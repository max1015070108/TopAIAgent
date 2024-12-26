package model

import (
	"context"
	"grpcservice/database"
	pb "grpcservice/proto/model"
)

type Server struct {
	pb.UnimplementedModelServiceServer
}

func (s *Server) GetModels(ctx context.Context, req *pb.EmptyRequest) (*pb.ModelListResponse, error) {
	var dbModels []database.Model
	// 直接使用全局 DB
	if err := database.DB.Find(&dbModels).Error; err != nil {
		return nil, err
	}

	models := make([]*pb.Model, len(dbModels))
	for i, dbModel := range dbModels {
		models[i] = &pb.Model{
			Id:   dbModel.ID,
			Name: dbModel.Name,
		}
	}

	return &pb.ModelListResponse{Models: models}, nil
}

func (s *Server) AddToNetwork(ctx context.Context, req *pb.NetworkRequest) (*pb.Response, error) {
	networkModel := database.NetworkModel{
		ModelID:   req.ModelId,
		NetworkID: req.NetworkId,
	}

	err := database.DB.Create(&networkModel).Error
	if err != nil {
		return &pb.Response{
			Success: false,
			Message: "Failed to add model to network: " + err.Error(),
		}, nil
	}

	return &pb.Response{
		Success: true,
		Message: "Model successfully added to network",
	}, nil
}

func (s *Server) RemoveFromNetwork(ctx context.Context, req *pb.RemoveRequest) (*pb.Response, error) {
	return &pb.Response{Message: "Successfully removed from network", Success: true}, nil
}
