package entity

import "time"

type Status string

const (
	StatusPending Status = "pending"
	StatusRunning Status = "running"
	StatusDone    Status = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	StartedAt   time.Time `json:"started_at,omitempty"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
	UserId      int       `json:"user_id"`
}

type CreateTaskRequest struct {
	UserID      int    `json:"user_id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}
