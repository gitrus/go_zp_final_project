package taskservice

import (
	"go_zp_final_project/internal/domain/entities"
	"go_zp_final_project/internal/domain/ports"
)

type TaskService struct {
	storage ports.Storage
}

func NewTaskService(storage ports.Storage) *TaskService {
	return &TaskService{storage: storage}
}

func (s *TaskService) CreateTask(task entities.Task) error {
	return nil
}

func (s *TaskService) GetTask() entities.Task {
	return entities.Task{}
}

func (s *TaskService) DoneTask(task entities.Task) error {
	return nil
}

func (s *TaskService) DeleteTask(task entities.Task) error {
	s.DoneTask(task)
	return nil
}
