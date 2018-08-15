package tasks

import (
	"context"
	"os"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"

	"github.com/ljesparis/micro-todo/services/tasks/client"
	proto "github.com/ljesparis/micro-todo/services/tasks/proto"
)

const (
	clientName      = "tasks.client"
	envVariableName = "REGISTRY_ADDRESS"
)

var (
	taskServiceClient proto.TasksService
	defaultContext    context.Context
)

func init() {
	registryAddress := registry.NewRegistry(
		registry.Addrs(os.Getenv(envVariableName)),
	)

	service := micro.NewService(
		micro.Name(clientName),
		micro.Registry(registryAddress),
	)

	taskServiceClient = client.NewTasksClient("tasks", service.Client())
	defaultContext = context.Background()
}
