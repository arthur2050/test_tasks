package crud

import (
	"test_tasks/entity"
	"test_tasks/repository"
	"time"
)

type TaskCrud struct {
	repo *repository.TaskRepository
}

func NewTaskCrud(repo *repository.TaskRepository) *TaskCrud {
	return &TaskCrud{repo: repo}
}

func (c *TaskCrud) Create(userID int, title string, description string) *entity.Task {
	id := c.repo.GetNextID()
	task := &entity.Task{
		ID:          id,
		UserId:      userID,
		Title:       title,
		Description: description,
		Status:      entity.StatusPending,
		CreatedAt:   time.Now(),
	}
	c.repo.Save(task)

	go c.run(task)
	return task
}

func (c *TaskCrud) run(task *entity.Task) {
	task.Status = entity.StatusRunning
	task.StartedAt = time.Now()
	c.repo.Save(task)

	// 1 минута для быстрого теста
	time.Sleep(1 * time.Minute)

	task.Status = entity.StatusDone
	task.CompletedAt = time.Now()
	c.repo.Save(task)
}

func (c *TaskCrud) Get(id int) (*entity.Task, bool) {
	return c.repo.Get(id)
}

func (c *TaskCrud) Delete(id int) bool {
	return c.repo.Delete(id)
}

func (c *TaskCrud) Update(id int, updated *entity.Task) (*entity.Task, bool) {
	existing, found := c.repo.Get(id)
	if !found {
		return nil, false
	}

	if updated.Title != "" {
		existing.Title = updated.Title
	}
	if updated.Description != "" {
		existing.Description = updated.Description
	}
	if updated.Status != "" {
		existing.Status = updated.Status
	}
	c.repo.Save(existing)
	return existing, true
}

func (c *TaskCrud) GetAll() []entity.Task {
	return c.repo.GetAll()
}
