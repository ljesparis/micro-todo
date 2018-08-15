package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/ljesparis/micro-todo/gateway/controllers/tasks"
)

var template string

func init() {
	f, err := os.Open(path.Join(os.Getenv("TEMPLATE_DIR"), "index.html"))
	if err != nil {
		panic(err)
	}

	d, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	template = string(d)
}

func Index(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/html;")
	fmt.Fprint(writer, template)
}

func main() {
	PORT := os.Getenv("PORT")
	addr := fmt.Sprintf("0.0.0.0:%s", PORT)
	fmt.Println(fmt.Sprintf("Starting server at address: %s", addr))
	r := mux.NewRouter()

	staticDir := os.Getenv("STATIC_DIR")
	r.PathPrefix("/ui").Handler(http.StripPrefix("/ui/", http.FileServer(http.Dir(staticDir))))
	r.HandleFunc("/", Index).Methods(http.MethodGet)
	r.HandleFunc("/create", tasks.CreateTask).Methods(http.MethodPost)
	r.HandleFunc("/tasks", tasks.AllTasks).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id:[0-9]+}", tasks.GetTask).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id:[0-9]+}/delete", tasks.DeleteTask).Methods(http.MethodDelete)
	r.HandleFunc("/tasks/{id:[0-9]+}/update", tasks.UpdateTask).Methods(http.MethodPut)

	fmt.Println(
		http.ListenAndServe(addr,
			handlers.RecoveryHandler()(
				handlers.LoggingHandler(os.Stdout, r),
			),
		),
	)
}
