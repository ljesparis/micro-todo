package tasks

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ljesparis/micro-todo/gateway/utils"
	"github.com/ljesparis/micro-todo/services/tasks/models"
	proto "github.com/ljesparis/micro-todo/services/tasks/proto"
)

func AllTasks(writer http.ResponseWriter, request *http.Request) {
	var code int
	var message utils.Response

	defer func() {
		utils.JsonResponse(writer, code, message)
	}()

	tmpOffset := request.URL.Query().Get("offset")
	if len(tmpOffset) == 0 {
		tmpOffset = "0"
	}

	offset, err := strconv.ParseInt(tmpOffset, 10, 64)
	if err != nil {
		code = http.StatusBadRequest
		message = utils.Response{"message": "bad request"}

		return
	}

	tmpLimit := request.URL.Query().Get("limit")
	if len(tmpLimit) == 0 {
		tmpLimit = "2"
	}

	limit, err := strconv.ParseInt(tmpLimit, 10, 64)
	if err != nil {
		code = http.StatusBadRequest
		message = utils.Response{"message": "bad request"}
		return
	}

	tmpOpen := request.URL.Query().Get("isopen")
	if len(tmpOpen) == 0 {
		tmpOpen = "true"
	}

	isopen, err := strconv.ParseBool(tmpOpen)
	if err != nil {
		code = http.StatusBadRequest
		message = utils.Response{"message": "bad request"}
		return
	}

	req := new(proto.GetTasksRequest)
	req.Offset = int32(offset)
	req.Limit = int32(limit)

	var res *proto.GetTasksResponse

	if isopen {
		res, err = taskServiceClient.GetOpenedTasks(defaultContext, req)
		if err != nil {
			code = http.StatusNotFound
			message = utils.Response{"message": err.Error()}
			return
		}
	} else {
		res, err = taskServiceClient.GetClosedTasks(defaultContext, req)
		if err != nil {
			code = http.StatusNotFound
			message = utils.Response{"message": err.Error()}
			return
		}
	}

	tasks := make([]models.Task, 0)
	for _, t := range res.Tasks {
		tasks = append(tasks, models.Task{
			Done:        t.Done,
			Description: t.Description,
			Name:        t.Name,
			ID:          int(t.Id),
		})
	}

	var i string
	if isopen {
		i = "&isopen=true"
	} else {
		i = "&isopen=false"
	}

	links := utils.Response{}
	if res.Next != nil {
		links["next"] = fmt.Sprintf("/tasks?limit=%d&offset=%d%s", res.Next.Limit, res.Next.Offset, i)
	}

	if res.Prev != nil {
		links["prev"] = fmt.Sprintf("/tasks?limit=%d&offset=%d%s", res.Prev.Limit, res.Prev.Offset, i)
	}

	code = http.StatusOK
	message = utils.Response{
		"tasks": tasks,
		"links": links,
	}
}

func GetTask(writer http.ResponseWriter, request *http.Request) {
	var code int
	var message utils.Response

	defer func() {
		utils.JsonResponse(writer, code, message)
	}()

	tmpOpen := request.URL.Query().Get("isopen")
	if len(tmpOpen) == 0 {
		tmpOpen = "true"
	}

	isopen, err := strconv.ParseBool(tmpOpen)
	if err != nil {
		code = http.StatusBadRequest
		message = utils.Response{"message": "bad request"}

		return
	}

	id, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)

	req := new(proto.CommonRequest)
	req.Id = int32(id)

	var res *proto.GetTaskResponse
	if isopen {
		res, err = taskServiceClient.GetOpenedTask(defaultContext, req)
		if err != nil {
			code = http.StatusNotFound
			message = utils.Response{"message": err.Error()}
			return
		}
	} else {
		res, err = taskServiceClient.GetClosedTask(defaultContext, req)
		if err != nil {
			code = http.StatusNotFound
			message = utils.Response{"message": err.Error()}
			return
		}
	}

	code = http.StatusOK
	message = utils.Response{
		"task": utils.Response{
			"name":        res.Task.Name,
			"description": res.Task.Description,
			"done":        res.Task.Done,
			"id":          res.Task.Id,
		},
	}
}
