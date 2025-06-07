package task

import "github.com/fmo/golang/crud/model"

type ServiceImpl struct {
	Repository Repository
}

func NewService(r Repository) *ServiceImpl {
	return &ServiceImpl{
		r,
	}
}

func (s ServiceImpl) Create(task model.Task) (*model.Task, error) {
	return s.Repository.Save(task)
}

func (s ServiceImpl) Read(taskID int) (*model.Task, error) {
	return s.Repository.Get(taskID)
}

func (s ServiceImpl) Update(task model.Task) error {
	return s.Repository.Update(task)
}

func (s ServiceImpl) Delete(taskID int) error {
	return s.Repository.Delete(taskID)
}
