package handler

import (
	"context"
	"errors"

	"github.com/ljesparis/micro-todo/services/tasks/db"
	"github.com/ljesparis/micro-todo/services/tasks/models"
	proto "github.com/ljesparis/micro-todo/services/tasks/proto"
)

func (t *TaskHandler) GetClosedTask(ctx context.Context, request *proto.CommonRequest, response *proto.GetTaskResponse) error {
	taskID := request.GetId()
	task, err := db.GetTaskById(int(taskID))

	if err != nil {
		return err
	}

	if !task.Done {
		return errors.New("task does not exists")
	}

	response.Task = new(proto.Task)
	response.Task.Name = task.Name
	response.Task.Done = task.Done
	response.Task.Description = task.Description
	response.Task.Id = int32(task.ID)

	return nil
}

func (t *TaskHandler) GetClosedTasks(ctx context.Context, request *proto.GetTasksRequest, response *proto.GetTasksResponse) error {
	var tmp []models.Task
	for _, t := range db.Tasks {
		if t.Done {
			tmp = append(tmp, t)
		}
	}

	var tmp2 []models.Task
	if len(tmp) <= int(request.Offset) {
		tmp2 = make([]models.Task, 0)
	} else {
		tmp2 = tmp[request.Offset:]

		var tmp3 []models.Task
		for _, tt := range tmp2 {
			if len(tmp3) < int(request.Limit) {
				tmp3 = append(tmp3, tt)
			}
		}
		tmp2 = tmp3
	}

	if request.Offset-request.Limit >= 0 {
		response.Prev = new(proto.GetTasksResponse_ListNode)
		response.Prev.Offset = request.Offset - request.Limit
		response.Prev.Limit = request.Limit
	} else {
		response.Prev = nil
	}

	if len(tmp) > int(request.Offset+request.Limit) {
		response.Next = new(proto.GetTasksResponse_ListNode)
		response.Next.Limit = request.Limit
		response.Next.Offset = request.Offset + request.Limit
	} else {
		response.Next = nil
	}

	var tasks []*proto.Task
	for _, task := range tmp2 {
		tasks = append(tasks, &proto.Task{
			Name:        task.Name,
			Id:          int32(task.ID),
			Description: task.Description,
			Done:        task.Done,
		})
	}

	response.Tasks = tasks
	return nil
}
