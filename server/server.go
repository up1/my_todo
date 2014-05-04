package server

import (
	"encoding/json"
	"net/http"
	"log"
	"strconv"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/up1/my_todo/task"
)

var tasks = task.NewTaskManager()

const PathPrefix = "/task/"

func RegisterHandlers() {
	route := mux.NewRouter()
	route.HandleFunc(PathPrefix, errorHandler(ListTasks)).Methods("GET")
	route.HandleFunc(PathPrefix, errorHandler(NewTask)).Methods("POST")
	route.HandleFunc(PathPrefix + "{id}", errorHandler(UpdateTask)).Methods("PUT")
	http.Handle(PathPrefix, route)
}

type badRequest struct{ error }
type notFound struct{ error }

func errorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case notFound:
			http.Error(w, "task not found", http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}

func ListTasks(w http.ResponseWriter, r *http.Request) error {
	res := struct{ Tasks []*task.Task }{tasks.GetAll()}
	return json.NewEncoder(w).Encode(res)
}

func NewTask(w http.ResponseWriter, r *http.Request) error {
	req := struct{ Title string }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return badRequest{err}
	}
	t, err := task.NewTask(req.Title)
	if err != nil {
		return badRequest{err}
	}
	return tasks.Save(t)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(r)
	if err != nil {
		return badRequest{err}
	}
	var t task.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		return badRequest{err}
	}
	if t.ID != id {
		return badRequest{fmt.Errorf("inconsistent task IDs")}
	}
	if _, ok := tasks.Find(id); !ok {
		return notFound{}
	}
	return tasks.Save(&t)
}

func parseID(r *http.Request) (int64, error) {
	txt, ok := mux.Vars(r)["id"]
	if !ok {
		return 0, fmt.Errorf("task id not found")
	}
	return strconv.ParseInt(txt, 10, 0)
}

func main() {
	RegisterHandlers()
	http.ListenAndServe(":8080", nil)
}
