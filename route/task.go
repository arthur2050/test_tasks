package route

import (
	"github.com/gin-gonic/gin"
	"test_tasks/controller"
)

func TaskRoutes(r *gin.RouterGroup, taskController *controller.TaskController) {
	taskRoute := r.Group("/task")

	taskRoute.GET("/:id", taskController.ShowTask)
	taskRoute.POST("/", taskController.CreateTask)
	taskRoute.PUT("/:id", taskController.UpdateTask)
	taskRoute.DELETE("/:id", taskController.DeleteTask)
	r.GET("/tasks", taskController.GetAll)
}
