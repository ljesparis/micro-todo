package client

import (
	proto "github.com/ljesparis/micro-todo/services/tasks/proto"
	"github.com/micro/go-micro/client"
)

func NewTasksClient(name string, c client.Client) proto.TasksService {
	return proto.NewTasksService(name, c)
}
