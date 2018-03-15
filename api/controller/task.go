package controller

import (
	"api/handler"
	"api/model"
	"api/store"
	"net/http"
)

type Task struct {
	*Base
	store store.Manager
}

func NewTask(s store.Manager) *Task {
	prefix := "/todo/tasks"

	t := &Task{
		store: s,
		Base: &Base{
			Prefix: prefix,
		},
	}

	routes := model.Routes{
		model.Route{
			Name:    "Create",
			Method:  "POST",
			Pattern: t.Prefix,
			Func:    t.Create,
		},
	}

	t.Routes = append(t.Routes, routes...)

	return t
}

func (t *Task) Create(w http.ResponseWriter, r *http.Request) {
	task := &model.Task{}

	err := handler.ParseJSON(r, task)
	if err != nil {
		handler.SendJSONError(w, "Error while decoding task", http.StatusBadRequest)
		return
	}

	err = t.store.Create(task)
	if err != nil {
		handler.SendJSONError(w, "Error while creating task", http.StatusInternalServerError)
		return
	}

	handler.SendJSONWithStatus(w, task, http.StatusCreated)
}
