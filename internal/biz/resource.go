package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Resource is a Resource model.
type Resource struct {
	Hello string
}

// ResourceRepo is a Greater repo.
type ResourceRepo interface {
	Save(context.Context, *Resource) (*Resource, error)
	Update(context.Context, *Resource) (*Resource, error)
	FindByID(context.Context, int64) (*Resource, error)
	ListByHello(context.Context, string) ([]*Resource, error)
	ListAll(context.Context) ([]*Resource, error)
}

// ResourceUsecase is a Resource usecase.
type ResourceUsecase struct {
	repo ResourceRepo
	log  *log.Helper
}

// NewResourceUsecase new a Resource usecase.
func NewResourceUsecase(repo ResourceRepo, logger log.Logger) *ResourceUsecase {
	return &ResourceUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateResource creates a Resource, and returns the new Resource.
func (uc *ResourceUsecase) CreateResource(ctx context.Context, g *Resource) (*Resource, error) {
	uc.log.WithContext(ctx).Infof("CreateResource: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
