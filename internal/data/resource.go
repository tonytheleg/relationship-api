package data

import (
	"context"

	"github.com/tonytheleg/relationship-api/internal/biz"

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

func (r *resourceRepo) Save(ctx context.Context, g *biz.Relationship) (*biz.Relationship, error) {
	return g, nil
}

func (r *resourceRepo) Update(ctx context.Context, g *biz.Relationship) (*biz.Relationship, error) {
	return g, nil
}

func (r *resourceRepo) FindByID(context.Context, int64) (*biz.Relationship, error) {
	return nil, nil
}

func (r *resourceRepo) ListByHello(context.Context, string) ([]*biz.Relationship, error) {
	return nil, nil
}

func (r *resourceRepo) ListAll(context.Context) ([]*biz.Relationship, error) {
	return nil, nil
}
