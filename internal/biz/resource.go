package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Relationship is a Relationship model.
type Relationship struct {
	Hello string
}

// ResourceRepo is a Greater repo.
type ResourceRepo interface {
	Save(context.Context, *Relationship) (*Relationship, error)
	Update(context.Context, *Relationship) (*Relationship, error)
	FindByID(context.Context, int64) (*Relationship, error)
	ListByHello(context.Context, string) ([]*Relationship, error)
	ListAll(context.Context) ([]*Relationship, error)
}

// ResourceUsecase is a Relationship usecase.
type ResourceUsecase struct {
	repo ResourceRepo
	log  *log.Helper
}

// NewResourceUsecase new a Relationship usecase.
func NewResourceUsecase(repo ResourceRepo, logger log.Logger) *ResourceUsecase {
	return &ResourceUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateResource creates a Relationship, and returns the new Relationship.
func (uc *ResourceUsecase) CreateResource(ctx context.Context, g *Relationship) (*Relationship, error) {
	uc.log.WithContext(ctx).Infof("CreateResource: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
