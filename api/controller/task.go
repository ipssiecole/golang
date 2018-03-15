package controller

import (
	"api/handler"
	"api/model"
	"api/store"
	"net/http"
	"strconv"
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
		model.Route{
			Name:    "All",
			Method:  "GET",
			Pattern: t.Prefix,
			Func:    t.All,
		},
	}

	t.Routes = append(t.Routes, routes...)

	return t
}

func (t *Task) All(w http.ResponseWriter, r *http.Request) {
	var err error

	offset := store.NoPaging
	limit := store.NoPaging

	offset, err = strconv.Atoi(handler.GetParams("offset", r))
	if err != nil {
		offset = store.NoPaging
	}

	limit, err = strconv.Atoi(handler.GetParams("limit", r))
	if err != nil {
		limit = store.NoPaging
	}

	tasks, err := t.store.All(offset, limit)
	if err != nil {
		handler.SendJSONError(w, "Error All", http.StatusInternalServerError)
		return
	}

	handler.SendJSONOk(w, &tasks)
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
