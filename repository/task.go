package repository

import (
	"sync"
	"test_tasks/entity"
)

type TaskRepository struct {
	mu     sync.RWMutex
	tasks  map[int]*entity.Task
	nextID int
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks:  make(map[int]*entity.Task),
		nextID: 1,
	}
}

func (r *TaskRepository) GetNextID() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := r.nextID
	r.nextID++
	return id
}

func (r *TaskRepository) Save(task *entity.Task) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[task.ID] = task
}

func (r *TaskRepository) Get(id int) (*entity.Task, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	task, ok := r.tasks[id]
	return task, ok
}

func (r *TaskRepository) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.tasks[id]; ok {
		delete(r.tasks, id)
		return true
	}
	return false
}

func (r *TaskRepository) GetAll() []entity.Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]entity.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		tasks = append(tasks, *t)
	}
	return tasks
}
