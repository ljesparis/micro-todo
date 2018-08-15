package tasks

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ljesparis/micro-todo/gateway/utils"
	proto "github.com/ljesparis/micro-todo/services/tasks/proto"
)

func UpdateTask(writer http.ResponseWriter, request *http.Request) {
	var code int
	var message string

	defer func() {
		utils.JsonResponse(writer, code, utils.Response{
			"message": message,
		})
	}()

	id, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)

	var updateBody struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Done        bool   `json:"done"`
	}

	err := json.NewDecoder(request.Body).Decode(&updateBody)
	if err != nil {
		code = http.StatusBadRequest
		message = "bad json format"

		return
	}

	req := new(proto.UpdateTaskRequest)
	req.Id = int32(id)

	req.Name = updateBody.Name
	req.Done = updateBody.Done
	req.Description = updateBody.Description

	res, err := taskServiceClient.UpdateTask(defaultContext, req)
	if err != nil {
		code = http.StatusNotFound
		message = err.Error()
	} else {
		code = http.StatusOK
		message = res.Message
	}
}
