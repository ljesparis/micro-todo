package handler

import (
	"context"

	"github.com/ljesparis/micro-todo/services/tasks/db"
	"github.com/ljesparis/micro-todo/services/tasks/models"
	proto "github.com/ljesparis/micro-todo/services/tasks/proto"
)

func (t *TaskHandler) UpdateTask(ctx context.Context, request *proto.UpdateTaskRequest, response *proto.CommonResponse) error {
	id := request.Id
	name := request.Name
	des := request.Description
	done := request.Done

	db.UpdateTask(&models.Task{
		ID:          int(id),
		Name:        name,
		Description: des,
		Done:        done,
	})

	response.Code = "201"
	response.Message = "task updated"

	return nil
}
