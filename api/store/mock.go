package store

import "api/model"

type mock struct{}

func NewMockManager() Manager {
	return &mock{}
}

func (m *mock) All(offset, limit int) ([]model.Task, error) {
	return []model.Task{}, nil
}

func (m *mock) Create(t *model.Task) error {
	return nil
}

func (m *mock) Update(t *model.Task) error {
	return nil
}

func (m *mock) Delete(ID string) error {
	return nil
}

func (m *mock) Find(ID string) (*model.Task, error) {
	return nil, nil
}

func (m *mock) FindByStatus(flag bool) ([]model.Task, error) {
	return []model.Task{}, nil
}
