// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/meeting.proto

/*
Package meeting is a generated protocol buffer package.

It is generated from these files:
	proto/meeting.proto

It has these top-level messages:
	Meeting
	AddRequest
	AddResponse
	GetRequest
	GetResponse
	ListRequest
	ListResponse
	DeleteRequest
	DeleteResponse
*/
package meeting

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

// Client API for MeetingService service

type MeetingService interface {
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type meetingService struct {
	c    client.Client
	name string
}

func NewMeetingService(name string, c client.Client) MeetingService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "meeting"
	}
	return &meetingService{
		c:    c,
		name: name,
	}
}

func (c *meetingService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "MeetingService.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "MeetingService.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "MeetingService.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "MeetingService.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MeetingService service

type MeetingServiceHandler interface {
	Add(context.Context, *AddRequest, *AddResponse) error
	Get(context.Context, *GetRequest, *GetResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterMeetingServiceHandler(s server.Server, hdlr MeetingServiceHandler, opts ...server.HandlerOption) {
	type meetingService interface {
		Add(ctx context.Context, in *AddRequest, out *AddResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type MeetingService struct {
		meetingService
	}
	h := &meetingServiceHandler{hdlr}
	s.Handle(s.NewHandler(&MeetingService{h}, opts...))
}

type meetingServiceHandler struct {
	MeetingServiceHandler
}

func (h *meetingServiceHandler) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.MeetingServiceHandler.Add(ctx, in, out)
}

func (h *meetingServiceHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.MeetingServiceHandler.Get(ctx, in, out)
}

func (h *meetingServiceHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.MeetingServiceHandler.List(ctx, in, out)
}

func (h *meetingServiceHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.MeetingServiceHandler.Delete(ctx, in, out)
}
