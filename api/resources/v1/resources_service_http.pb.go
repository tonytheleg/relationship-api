// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v5.27.3
// source: api/resources/v1/resources_service.proto

package resources

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationKesselResourceServiceCreateResource = "/resources.v1.KesselResourceService/CreateResource"
const OperationKesselResourceServiceDeleteResource = "/resources.v1.KesselResourceService/DeleteResource"
const OperationKesselResourceServiceUpdateResource = "/resources.v1.KesselResourceService/UpdateResource"

type KesselResourceServiceHTTPServer interface {
	CreateResource(context.Context, *CreateResourceRequest) (*CreateResourceResponse, error)
	DeleteResource(context.Context, *DeleteResourceRequest) (*DeleteResourceResponse, error)
	UpdateResource(context.Context, *UpdateResourceRequest) (*UpdateResourceResponse, error)
}

func RegisterKesselResourceServiceHTTPServer(s *http.Server, srv KesselResourceServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/resources", _KesselResourceService_CreateResource0_HTTP_Handler(srv))
	r.PUT("/api/v1/resources", _KesselResourceService_UpdateResource0_HTTP_Handler(srv))
	r.DELETE("/api/v1/resources", _KesselResourceService_DeleteResource0_HTTP_Handler(srv))
}

func _KesselResourceService_CreateResource0_HTTP_Handler(srv KesselResourceServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateResourceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKesselResourceServiceCreateResource)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateResource(ctx, req.(*CreateResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateResourceResponse)
		return ctx.Result(200, reply)
	}
}

func _KesselResourceService_UpdateResource0_HTTP_Handler(srv KesselResourceServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateResourceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKesselResourceServiceUpdateResource)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateResource(ctx, req.(*UpdateResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateResourceResponse)
		return ctx.Result(200, reply)
	}
}

func _KesselResourceService_DeleteResource0_HTTP_Handler(srv KesselResourceServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteResourceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKesselResourceServiceDeleteResource)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteResource(ctx, req.(*DeleteResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteResourceResponse)
		return ctx.Result(200, reply)
	}
}

type KesselResourceServiceHTTPClient interface {
	CreateResource(ctx context.Context, req *CreateResourceRequest, opts ...http.CallOption) (rsp *CreateResourceResponse, err error)
	DeleteResource(ctx context.Context, req *DeleteResourceRequest, opts ...http.CallOption) (rsp *DeleteResourceResponse, err error)
	UpdateResource(ctx context.Context, req *UpdateResourceRequest, opts ...http.CallOption) (rsp *UpdateResourceResponse, err error)
}

type KesselResourceServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewKesselResourceServiceHTTPClient(client *http.Client) KesselResourceServiceHTTPClient {
	return &KesselResourceServiceHTTPClientImpl{client}
}

func (c *KesselResourceServiceHTTPClientImpl) CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...http.CallOption) (*CreateResourceResponse, error) {
	var out CreateResourceResponse
	pattern := "/api/v1/resources"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKesselResourceServiceCreateResource))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *KesselResourceServiceHTTPClientImpl) DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...http.CallOption) (*DeleteResourceResponse, error) {
	var out DeleteResourceResponse
	pattern := "/api/v1/resources"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKesselResourceServiceDeleteResource))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *KesselResourceServiceHTTPClientImpl) UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...http.CallOption) (*UpdateResourceResponse, error) {
	var out UpdateResourceResponse
	pattern := "/api/v1/resources"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKesselResourceServiceUpdateResource))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}