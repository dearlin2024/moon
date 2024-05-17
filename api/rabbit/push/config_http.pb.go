// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.19.4
// source: rabbit/push/config.proto

package push

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

const OperationConfigNotifyObject = "/api.rabbit.push.Config/NotifyObject"

type ConfigHTTPServer interface {
	NotifyObject(context.Context, *NotifyObjectRequest) (*NotifyObjectReply, error)
}

func RegisterConfigHTTPServer(s *http.Server, srv ConfigHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/rabbit/push/config", _Config_NotifyObject0_HTTP_Handler(srv))
}

func _Config_NotifyObject0_HTTP_Handler(srv ConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NotifyObjectRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationConfigNotifyObject)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NotifyObject(ctx, req.(*NotifyObjectRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NotifyObjectReply)
		return ctx.Result(200, reply)
	}
}

type ConfigHTTPClient interface {
	NotifyObject(ctx context.Context, req *NotifyObjectRequest, opts ...http.CallOption) (rsp *NotifyObjectReply, err error)
}

type ConfigHTTPClientImpl struct {
	cc *http.Client
}

func NewConfigHTTPClient(client *http.Client) ConfigHTTPClient {
	return &ConfigHTTPClientImpl{client}
}

func (c *ConfigHTTPClientImpl) NotifyObject(ctx context.Context, in *NotifyObjectRequest, opts ...http.CallOption) (*NotifyObjectReply, error) {
	var out NotifyObjectReply
	pattern := "/v1/rabbit/push/config"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationConfigNotifyObject))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
