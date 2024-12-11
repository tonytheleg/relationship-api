package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Relationship is a Relationship model.
type Relationship struct {
	Hello string
}

// RelationshipRepo is a Greater repo.
type RelationshipRepo interface {
	Save(context.Context, *Relationship) (*Relationship, error)
	Update(context.Context, *Relationship) (*Relationship, error)
	FindByID(context.Context, int64) (*Relationship, error)
	ListByHello(context.Context, string) ([]*Relationship, error)
	ListAll(context.Context) ([]*Relationship, error)
}

// RelationshipUsecase is a Relationship usecase.
type RelationshipUsecase struct {
	repo RelationshipRepo
	log  *log.Helper
}

// NewRelationshipUsecase new a Relationship usecase.
func NewRelationshipUsecase(repo RelationshipRepo, logger log.Logger) *RelationshipUsecase {
	return &RelationshipUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateRelationship creates a Relationship, and returns the new Relationship.
func (uc *RelationshipUsecase) CreateRelationship(ctx context.Context, g *Relationship) (*Relationship, error) {
	uc.log.WithContext(ctx).Infof("CreateRelationship: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
