package repository

import (
	"github.com/fmo/golang/crud/infrastructure/db"
	"github.com/fmo/golang/crud/model"
)

type Task struct {
	db *db.Database
}

func NewTask(db *db.Database) *Task {
	return &Task{
		db,
	}
}

func (t *Task) Save(task model.Task) (*model.Task, error) {
	var insertedTask model.Task
	query := `INSERT INTO tasks (name) VALUES ($1) RETURNING id`
	err := t.db.QueryRow(query, task.Name).Scan(&insertedTask.ID)
	if err != nil {
		return nil, err
	}
	return &insertedTask, nil
}

func (t *Task) Update(task model.Task) error {
	query := `UPDATE tasks SET name=$1 WHERE id=$2`
	_, err := t.db.Exec(query, task.Name, task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t *Task) Delete(taskID int) error {
	query := `DELETE from tasks WHERE id=$1`
	_, err := t.db.Exec(query, taskID)
	if err != nil {
		return err
	}
	return nil
}

func (t *Task) Get(taskID int) (*model.Task, error) {
	var task model.Task
	query := "SELECT id, name FROM tasks WHERE id=$1"
	err := t.db.QueryRow(query, taskID).Scan(&task.ID, &task.Name)
	if err != nil {
		return nil, err
	}
	return &task, nil
}
