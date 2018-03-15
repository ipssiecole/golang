package store

import "api/model"

const NoPaging = -1

type Manager interface {
	All(offset, limit int) ([]model.Task, error)

	Create(*model.Task) error

	Find(string) (model.Task, error)

	FindByStatus(bool) (model.Task, error)

	Update(*model.Task) error

	Delete(string) error
}
