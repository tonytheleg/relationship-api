package data

import (
	"context"

	"github.com/tonytheleg/relationship-api/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type relationshipRepo struct {
	data *Data
	log  *log.Helper
}

// NewRelationshipRepo .
func NewRelationshipRepo(data *Data, logger log.Logger) biz.RelationshipRepo {
	return &relationshipRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *relationshipRepo) Save(ctx context.Context, g *biz.Relationship) (*biz.Relationship, error) {
	return g, nil
}

func (r *relationshipRepo) Update(ctx context.Context, g *biz.Relationship) (*biz.Relationship, error) {
	return g, nil
}

func (r *relationshipRepo) FindByID(context.Context, int64) (*biz.Relationship, error) {
	return nil, nil
}

func (r *relationshipRepo) ListByHello(context.Context, string) ([]*biz.Relationship, error) {
	return nil, nil
}

func (r *relationshipRepo) ListAll(context.Context) ([]*biz.Relationship, error) {
	return nil, nil
}
