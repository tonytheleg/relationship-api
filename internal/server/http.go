package server

import (
	"github.com/bufbuild/protovalidate-go"
	v1 "github.com/tonytheleg/relationship-api/api/relationships/v1"

	"github.com/tonytheleg/relationship-api/internal/conf"
	m "github.com/tonytheleg/relationship-api/internal/middleware"
	"github.com/tonytheleg/relationship-api/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, rs *service.KesselResourceServiceService, logger log.Logger) *http.Server {
	validator, _ := protovalidate.New()

	var opts = []http.ServerOption{
		http.Middleware(
			logging.Client(logger),
			recovery.Recovery(),
			m.Validation(validator),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterKesselResourceServiceHTTPServer(srv, rs)
	return srv
}
