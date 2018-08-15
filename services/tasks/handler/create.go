package handler

import (
	"context"

	"github.com/ljesparis/micro-todo/services/tasks/db"
	"github.com/ljesparis/micro-todo/services/tasks/models"
	proto "github.com/ljesparis/micro-todo/services/tasks/proto"
)

func (t *TaskHandler) CreateTask(ctx context.Context, request *proto.CreateTaskRequest, response *proto.CommonResponse) error {
	name := request.Name
	des := request.Description

	db.CreateTask(&models.Task{
		Name:        name,
		Description: des,
		Done:        false,
		ID:          len(db.Tasks) + 1,
	})

	response.Message = "task created"
	response.Code = "201"
	return nil
}
