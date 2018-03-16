package controller

import (
	"api/handler"
	"api/model"
	"api/store"
	"log"
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
		model.Route{
			Name:    "Find",
			Method:  "GET",
			Pattern: t.Prefix + "/{id}",
			Func:    t.FindById,
		},
		model.Route{
			Name:    "FindByStatus",
			Method:  "GET",
			Pattern: t.Prefix + "/status/{status:(?:true|false)}",
			Func:    t.FindByStatus,
		},
		model.Route{
			Name:    "Update",
			Method:  "PUT",
			Pattern: t.Prefix,
			Func:    t.Update,
		},
		model.Route{
			Name:    "Delete",
			Method:  "DELETE",
			Pattern: t.Prefix + "/{id}",
			Func:    t.Delete,
		},
	}

	t.Routes = append(t.Routes, routes...)

	return t
}

func (t *Task) All(w http.ResponseWriter, r *http.Request) {
	var err error

	offset := store.NoPaging
	limit := store.NoPaging

	offset, err = strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = store.NoPaging
	}

	limit, err = strconv.Atoi(r.URL.Query().Get("limit"))
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

func (t *Task) FindById(w http.ResponseWriter, r *http.Request) {
	id := handler.GetParams("id", r)

	task, err := t.store.Find(id)
	if err != nil {
		log.Println(err)
		handler.SendJSONError(w, "Error while retrieving task", http.StatusInternalServerError)
		return
	}

	handler.SendJSONOk(w, task)
}

func (t *Task) Delete(w http.ResponseWriter, r *http.Request) {
	id := handler.GetParams("id", r)

	err := t.store.Delete(id)
	if err != nil {
		log.Println(err)
		handler.SendJSONError(w, "Error while deleting task", http.StatusInternalServerError)
		return
	}

	handler.SendJSONWithStatus(w, "", http.StatusNoContent)
}

func (t *Task) Update(w http.ResponseWriter, r *http.Request) {
	task := &model.Task{}

	err := handler.ParseJSON(r, task)
	if err != nil {
		log.Println(err)
		handler.SendJSONError(w, "", http.StatusBadRequest)
		return
	}

	err = t.store.Update(task)
	if err != nil {
		log.Println(err)
		handler.SendJSONError(w, "Error while updating task", http.StatusInternalServerError)
		return
	}

	handler.SendJSONWithStatus(w, "", http.StatusNoContent)
}

func (t *Task) FindByStatus(w http.ResponseWriter, r *http.Request) {
	status, err := strconv.ParseBool(handler.GetParams("status", r))
	if err != nil {
		log.Println(err)
		handler.SendJSONError(w, "", http.StatusBadRequest)
		return
	}

	tasks, err := t.store.FindByStatus(status)
	if err != nil {
		log.Println(err)
		handler.SendJSONError(w, "Error while retrieving task", http.StatusInternalServerError)
		return
	}

	handler.SendJSONOk(w, tasks)
}
