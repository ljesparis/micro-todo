package main

import (
	"github.com/micro/go-micro"

	_ "github.com/ljesparis/micro-todo/services/tasks/db"
	"github.com/ljesparis/micro-todo/services/tasks/handler"
	proto "github.com/ljesparis/micro-todo/services/tasks/proto"
	"github.com/micro/go-micro/registry"
	"os"
)

const (
	envVariableName = "REGISTRY_ADDRESS"
)

func main() {
	registryAddress := registry.NewRegistry(
		registry.Addrs(os.Getenv(envVariableName)),
	)

	service := micro.NewService(
		micro.Name("tasks"),
		micro.Version("latest"),
		micro.Registry(registryAddress),
	)

	service.Init()
	proto.RegisterTasksHandler(service.Server(), new(handler.TaskHandler))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
