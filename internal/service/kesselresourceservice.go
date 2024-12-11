package service

import (
	"context"

	pb "github.com/tonytheleg/relationship-api/api/relationships/v1"

	"github.com/tonytheleg/relationship-api/internal/biz"
)

type KesselResourceServiceService struct {
	pb.UnimplementedKesselResourceServiceServer
	uc *biz.ResourceUsecase
}

func NewKesselResourceServiceService(uc *biz.ResourceUsecase) *KesselResourceServiceService {
	return &KesselResourceServiceService{uc: uc}
}

func (s *KesselResourceServiceService) CreateResource(ctx context.Context, req *pb.CreateResourceRequest) (*pb.CreateResourceResponse, error) {
	return &pb.CreateResourceResponse{}, nil
}
func (s *KesselResourceServiceService) UpdateResource(ctx context.Context, req *pb.UpdateResourceRequest) (*pb.UpdateResourceResponse, error) {
	return &pb.UpdateResourceResponse{}, nil
}
func (s *KesselResourceServiceService) DeleteResource(ctx context.Context, req *pb.DeleteResourceRequest) (*pb.DeleteResourceResponse, error) {
	return &pb.DeleteResourceResponse{}, nil
}
