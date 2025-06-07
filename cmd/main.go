package main

import (
	"fmt"
	"github.com/fmo/golang/crud/infrastructure/db"
	"github.com/fmo/golang/crud/infrastructure/repository"
	"github.com/fmo/golang/crud/model"
	"github.com/fmo/golang/crud/service/task"
)
import _ "github.com/lib/pq"

func main() {
	database := db.NewDatabase("user=user dbname=tasks sslmode=disable password=password port=5433")

	taskRepo := repository.NewTask(database)
	taskService := task.NewService(taskRepo)

	// CREATE
	t, err := taskService.Create(model.Task{
		Name: "fifth task",
	})
	if err != nil {
		panic(err)
	}

	// READ
	gotTask, err := taskService.Read(t.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(gotTask.Name)

	gotTask.Name = "Updated second task"

	// UPDATE
	err = taskService.Update(*gotTask)
	if err != nil {
		panic(err)
	}

	gotUpdatedTask, err := taskService.Read(gotTask.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(gotUpdatedTask.Name)

	// DELETE
	err = taskService.Delete(gotUpdatedTask.ID)
	if err != nil {
		panic(err)
	}

	deletedTask, err := taskService.Read(gotUpdatedTask.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(deletedTask.Name)
}
