package service

import (
	"fmt"
	"to-do-list/internal/entity"
	"to-do-list/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title string, description string) error {
	if title == "" {
		return fmt.Errorf("title is required")
	}

	task := &entity.Task{Title: title, Description: description, Completed: false}
	return s.repo.Create(task)
}

func (s *TaskService) FindById(id int) (entity.Task, error) {
	return s.repo.FindById(id)
}

func (s *TaskService) FindAll() ([]entity.Task, error) {
	return s.repo.FindAll()
}
