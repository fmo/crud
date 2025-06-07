package task

import "github.com/fmo/golang/crud/model"

type Service interface {
	Create(task model.Task) (*model.Task, error)
	Read(taskID int) (*model.Task, error)
	Update(task model.Task) error
	Delete(taskID int) error
}
