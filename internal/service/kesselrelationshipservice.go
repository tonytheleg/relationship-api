package service

import (
	"context"

	pb "github.com/tonytheleg/relationship-api/api/relationships/v1"

	"github.com/tonytheleg/relationship-api/internal/biz"
)

type KesselRelationshipserviceService struct {
	pb.UnimplementedKesselRelationshipserviceServer
	uc *biz.RelationshipUsecase
}

func NewKesselRelationshipserviceService(uc *biz.RelationshipUsecase) *KesselRelationshipserviceService {
	return &KesselRelationshipserviceService{uc: uc}
}

func (s *KesselRelationshipserviceService) CreateRelationship(ctx context.Context, req *pb.CreateRelationshipRequest) (*pb.CreateRelationshipResponse, error) {
	return &pb.CreateRelationshipResponse{}, nil
}
func (s *KesselRelationshipserviceService) UpdateRelationship(ctx context.Context, req *pb.UpdateRelationshipRequest) (*pb.UpdateRelationshipResponse, error) {
	return &pb.UpdateRelationshipResponse{}, nil
}
func (s *KesselRelationshipserviceService) DeleteRelationship(ctx context.Context, req *pb.DeleteRelationshipRequest) (*pb.DeleteRelationshipResponse, error) {
	return &pb.DeleteRelationshipResponse{}, nil
}
