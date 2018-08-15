package handler

import (
	"context"

	"github.com/ljesparis/micro-todo/services/tasks/db"
	proto "github.com/ljesparis/micro-todo/services/tasks/proto"
)

func (t *TaskHandler) DeleteTask(ctx context.Context, request *proto.CommonRequest, response *proto.CommonResponse) error {
	err := db.DeleteTask(int(request.Id))

	if err != nil {
		response.Message = err.Error()
		response.Code = "200"

		return err
	}

	response.Message = "task deleted"
	response.Code = "200"

	return nil
}
