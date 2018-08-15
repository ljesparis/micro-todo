package tasks

import (
	"encoding/json"
	"net/http"

	"github.com/ljesparis/micro-todo/gateway/utils"
	proto "github.com/ljesparis/micro-todo/services/tasks/proto"
)

func CreateTask(writer http.ResponseWriter, request *http.Request) {
	var code int
	var message string

	defer func() {
		utils.JsonResponse(writer, code, utils.Response{
			"message": message,
		})
	}()

	var createTaskBody struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	err := json.NewDecoder(request.Body).Decode(&createTaskBody)
	if err != nil {
		code = http.StatusBadRequest
		message = "bad json format"

		return
	}

	if len(createTaskBody.Name) == 0 && len(createTaskBody.Description) == 0 {
		code = http.StatusOK
		message = "name and description field should not be empty"

		return
	}

	req := new(proto.CreateTaskRequest)
	req.Name = createTaskBody.Name
	req.Description = createTaskBody.Description

	_, err = taskServiceClient.CreateTask(defaultContext, req)
	if err != nil {
		code = http.StatusNotFound
		message = "task could not be created"
	} else {
		code = http.StatusNoContent
		message = ""
	}
}
