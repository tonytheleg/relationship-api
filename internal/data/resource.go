package data

import (
	"context"

	"github.com/tonytheleg/resource-api/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type resourceRepo struct {
	data *Data
	log  *log.Helper
}

// NewResourceRepo .
func NewResourceRepo(data *Data, logger log.Logger) biz.ResourceRepo {
	return &resourceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *resourceRepo) Save(ctx context.Context, g *biz.Resource) (*biz.Resource, error) {
	return g, nil
}

func (r *resourceRepo) Update(ctx context.Context, g *biz.Resource) (*biz.Resource, error) {
	return g, nil
}

func (r *resourceRepo) FindByID(context.Context, int64) (*biz.Resource, error) {
	return nil, nil
}

func (r *resourceRepo) ListByHello(context.Context, string) ([]*biz.Resource, error) {
	return nil, nil
}

func (r *resourceRepo) ListAll(context.Context) ([]*biz.Resource, error) {
	return nil, nil
}
