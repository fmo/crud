package task

import "github.com/fmo/golang/crud/model"

type Repository interface {
	Save(task model.Task) (*model.Task, error)
	Update(task model.Task) error
	Delete(taskID int) error
	Get(taskID int) (*model.Task, error)
}
