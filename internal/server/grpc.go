package server

import (
	"github.com/bufbuild/protovalidate-go"
	v1 "github.com/tonytheleg/resource-api/api/resources/v1"

	"github.com/tonytheleg/resource-api/internal/conf"
	"github.com/tonytheleg/resource-api/internal/middleware"
	"github.com/tonytheleg/resource-api/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, rs *service.KesselResourceServiceService, logger log.Logger) *grpc.Server {
	validator, _ := protovalidate.New()

	var opts = []grpc.ServerOption{
		grpc.Middleware(
			logging.Server(logger),
			recovery.Recovery(),
			middleware.Validation(validator),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterKesselResourceServiceServer(srv, rs)
	return srv
}
