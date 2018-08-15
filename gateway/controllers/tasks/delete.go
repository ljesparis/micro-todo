package tasks

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ljesparis/micro-todo/gateway/utils"
	proto "github.com/ljesparis/micro-todo/services/tasks/proto"
)

func DeleteTask(writer http.ResponseWriter, request *http.Request) {
	var code int
	var message string

	defer func() {
		utils.JsonResponse(writer, code, utils.Response{
			"message": message,
		})
	}()

	id, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)

	req := new(proto.CommonRequest)
	req.Id = int32(id)

	res, err := taskServiceClient.DeleteTask(defaultContext, req)
	if err != nil {
		message = err.Error()
		code = http.StatusNotFound
	} else {
		message = res.Message
		code = http.StatusOK
	}
}
