// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/scraper.proto

/*
Package scraper is a generated protocol buffer package.

It is generated from these files:
	proto/scraper.proto

It has these top-level messages:
	ScrapeItem
	ScrapeHistoryItem
	GetWorkQueueRequest
	GetWorkQueueResponse
	GetStatusRequest
	GetStatusResponse
	GetWorkHistoryRequest
	GetWorkHistoryResponse
*/
package scraper

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ScraperService service

type ScraperService interface {
	GetWorkQueue(ctx context.Context, in *GetWorkQueueRequest, opts ...client.CallOption) (*GetWorkQueueResponse, error)
	GetWorkHistory(ctx context.Context, in *GetWorkHistoryRequest, opts ...client.CallOption) (*GetWorkHistoryResponse, error)
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...client.CallOption) (*GetStatusResponse, error)
}

type scraperService struct {
	c    client.Client
	name string
}

func NewScraperService(name string, c client.Client) ScraperService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "scraper"
	}
	return &scraperService{
		c:    c,
		name: name,
	}
}

func (c *scraperService) GetWorkQueue(ctx context.Context, in *GetWorkQueueRequest, opts ...client.CallOption) (*GetWorkQueueResponse, error) {
	req := c.c.NewRequest(c.name, "ScraperService.GetWorkQueue", in)
	out := new(GetWorkQueueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scraperService) GetWorkHistory(ctx context.Context, in *GetWorkHistoryRequest, opts ...client.CallOption) (*GetWorkHistoryResponse, error) {
	req := c.c.NewRequest(c.name, "ScraperService.GetWorkHistory", in)
	out := new(GetWorkHistoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scraperService) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...client.CallOption) (*GetStatusResponse, error) {
	req := c.c.NewRequest(c.name, "ScraperService.GetStatus", in)
	out := new(GetStatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ScraperService service

type ScraperServiceHandler interface {
	GetWorkQueue(context.Context, *GetWorkQueueRequest, *GetWorkQueueResponse) error
	GetWorkHistory(context.Context, *GetWorkHistoryRequest, *GetWorkHistoryResponse) error
	GetStatus(context.Context, *GetStatusRequest, *GetStatusResponse) error
}

func RegisterScraperServiceHandler(s server.Server, hdlr ScraperServiceHandler, opts ...server.HandlerOption) {
	type scraperService interface {
		GetWorkQueue(ctx context.Context, in *GetWorkQueueRequest, out *GetWorkQueueResponse) error
		GetWorkHistory(ctx context.Context, in *GetWorkHistoryRequest, out *GetWorkHistoryResponse) error
		GetStatus(ctx context.Context, in *GetStatusRequest, out *GetStatusResponse) error
	}
	type ScraperService struct {
		scraperService
	}
	h := &scraperServiceHandler{hdlr}
	s.Handle(s.NewHandler(&ScraperService{h}, opts...))
}

type scraperServiceHandler struct {
	ScraperServiceHandler
}

func (h *scraperServiceHandler) GetWorkQueue(ctx context.Context, in *GetWorkQueueRequest, out *GetWorkQueueResponse) error {
	return h.ScraperServiceHandler.GetWorkQueue(ctx, in, out)
}

func (h *scraperServiceHandler) GetWorkHistory(ctx context.Context, in *GetWorkHistoryRequest, out *GetWorkHistoryResponse) error {
	return h.ScraperServiceHandler.GetWorkHistory(ctx, in, out)
}

func (h *scraperServiceHandler) GetStatus(ctx context.Context, in *GetStatusRequest, out *GetStatusResponse) error {
	return h.ScraperServiceHandler.GetStatus(ctx, in, out)
}
