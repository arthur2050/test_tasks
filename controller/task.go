package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test_tasks/entity"
	"test_tasks/service/crud"
)

type TaskController struct {
	crud *crud.TaskCrud
}

func NewTaskController(c *crud.TaskCrud) *TaskController {
	return &TaskController{crud: c}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var req entity.CreateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	task := c.crud.Create(req.UserID, req.Title, req.Description)
	ctx.JSON(http.StatusCreated, task)
}

func (c *TaskController) ShowTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	task, found := c.crud.Get(id)
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	if c.crud.Delete(id) {
		ctx.JSON(http.StatusOK, gin.H{"message": "task deleted"})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
	}
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	var updated entity.Task
	if err := ctx.ShouldBindJSON(&updated); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	task, ok := c.crud.Update(id, &updated)
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) GetAll(ctx *gin.Context) {
	tasks := c.crud.GetAll()
	ctx.JSON(http.StatusOK, tasks)
}
