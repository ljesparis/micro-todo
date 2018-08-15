// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: tasks.proto

/*
Package tasks is a generated protocol buffer package.

It is generated from these files:
	tasks.proto

It has these top-level messages:
	Task
	CommonRequest
	GetTasksRequest
	CreateTaskRequest
	UpdateTaskRequest
	CommonResponse
	GetTaskResponse
	GetTasksResponse
*/
package tasks

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for Tasks service

type TasksService interface {
	GetOpenedTask(ctx context.Context, in *CommonRequest, opts ...client.CallOption) (*GetTaskResponse, error)
	GetOpenedTasks(ctx context.Context, in *GetTasksRequest, opts ...client.CallOption) (*GetTasksResponse, error)
	GetClosedTask(ctx context.Context, in *CommonRequest, opts ...client.CallOption) (*GetTaskResponse, error)
	GetClosedTasks(ctx context.Context, in *GetTasksRequest, opts ...client.CallOption) (*GetTasksResponse, error)
	DeleteTask(ctx context.Context, in *CommonRequest, opts ...client.CallOption) (*CommonResponse, error)
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...client.CallOption) (*CommonResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...client.CallOption) (*CommonResponse, error)
}

type tasksService struct {
	c    client.Client
	name string
}

func NewTasksService(name string, c client.Client) TasksService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "tasks"
	}
	return &tasksService{
		c:    c,
		name: name,
	}
}

func (c *tasksService) GetOpenedTask(ctx context.Context, in *CommonRequest, opts ...client.CallOption) (*GetTaskResponse, error) {
	req := c.c.NewRequest(c.name, "Tasks.GetOpenedTask", in)
	out := new(GetTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksService) GetOpenedTasks(ctx context.Context, in *GetTasksRequest, opts ...client.CallOption) (*GetTasksResponse, error) {
	req := c.c.NewRequest(c.name, "Tasks.GetOpenedTasks", in)
	out := new(GetTasksResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksService) GetClosedTask(ctx context.Context, in *CommonRequest, opts ...client.CallOption) (*GetTaskResponse, error) {
	req := c.c.NewRequest(c.name, "Tasks.GetClosedTask", in)
	out := new(GetTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksService) GetClosedTasks(ctx context.Context, in *GetTasksRequest, opts ...client.CallOption) (*GetTasksResponse, error) {
	req := c.c.NewRequest(c.name, "Tasks.GetClosedTasks", in)
	out := new(GetTasksResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksService) DeleteTask(ctx context.Context, in *CommonRequest, opts ...client.CallOption) (*CommonResponse, error) {
	req := c.c.NewRequest(c.name, "Tasks.DeleteTask", in)
	out := new(CommonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksService) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...client.CallOption) (*CommonResponse, error) {
	req := c.c.NewRequest(c.name, "Tasks.CreateTask", in)
	out := new(CommonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksService) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...client.CallOption) (*CommonResponse, error) {
	req := c.c.NewRequest(c.name, "Tasks.UpdateTask", in)
	out := new(CommonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tasks service

type TasksHandler interface {
	GetOpenedTask(context.Context, *CommonRequest, *GetTaskResponse) error
	GetOpenedTasks(context.Context, *GetTasksRequest, *GetTasksResponse) error
	GetClosedTask(context.Context, *CommonRequest, *GetTaskResponse) error
	GetClosedTasks(context.Context, *GetTasksRequest, *GetTasksResponse) error
	DeleteTask(context.Context, *CommonRequest, *CommonResponse) error
	CreateTask(context.Context, *CreateTaskRequest, *CommonResponse) error
	UpdateTask(context.Context, *UpdateTaskRequest, *CommonResponse) error
}

func RegisterTasksHandler(s server.Server, hdlr TasksHandler, opts ...server.HandlerOption) error {
	type tasks interface {
		GetOpenedTask(ctx context.Context, in *CommonRequest, out *GetTaskResponse) error
		GetOpenedTasks(ctx context.Context, in *GetTasksRequest, out *GetTasksResponse) error
		GetClosedTask(ctx context.Context, in *CommonRequest, out *GetTaskResponse) error
		GetClosedTasks(ctx context.Context, in *GetTasksRequest, out *GetTasksResponse) error
		DeleteTask(ctx context.Context, in *CommonRequest, out *CommonResponse) error
		CreateTask(ctx context.Context, in *CreateTaskRequest, out *CommonResponse) error
		UpdateTask(ctx context.Context, in *UpdateTaskRequest, out *CommonResponse) error
	}
	type Tasks struct {
		tasks
	}
	h := &tasksHandler{hdlr}
	return s.Handle(s.NewHandler(&Tasks{h}, opts...))
}

type tasksHandler struct {
	TasksHandler
}

func (h *tasksHandler) GetOpenedTask(ctx context.Context, in *CommonRequest, out *GetTaskResponse) error {
	return h.TasksHandler.GetOpenedTask(ctx, in, out)
}

func (h *tasksHandler) GetOpenedTasks(ctx context.Context, in *GetTasksRequest, out *GetTasksResponse) error {
	return h.TasksHandler.GetOpenedTasks(ctx, in, out)
}

func (h *tasksHandler) GetClosedTask(ctx context.Context, in *CommonRequest, out *GetTaskResponse) error {
	return h.TasksHandler.GetClosedTask(ctx, in, out)
}

func (h *tasksHandler) GetClosedTasks(ctx context.Context, in *GetTasksRequest, out *GetTasksResponse) error {
	return h.TasksHandler.GetClosedTasks(ctx, in, out)
}

func (h *tasksHandler) DeleteTask(ctx context.Context, in *CommonRequest, out *CommonResponse) error {
	return h.TasksHandler.DeleteTask(ctx, in, out)
}

func (h *tasksHandler) CreateTask(ctx context.Context, in *CreateTaskRequest, out *CommonResponse) error {
	return h.TasksHandler.CreateTask(ctx, in, out)
}

func (h *tasksHandler) UpdateTask(ctx context.Context, in *UpdateTaskRequest, out *CommonResponse) error {
	return h.TasksHandler.UpdateTask(ctx, in, out)
}