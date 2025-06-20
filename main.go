package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test_tasks/controller"
	"test_tasks/repository"
	"test_tasks/route"
	"test_tasks/service/crud"
)

func main() {
	r := gin.Default()
	api := r.Group("/api/v1")
	taskRepository := repository.NewTaskRepository()
	taskCrud := crud.NewTaskCrud(taskRepository)
	taskController := controller.NewTaskController(taskCrud)

	route.TaskRoutes(api, taskController)
	err := r.Run(":8080")
	if err != nil {
		return
	}
	fmt.Println("Start task server...")
}
