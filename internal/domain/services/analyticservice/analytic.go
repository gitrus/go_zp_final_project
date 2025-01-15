package analyticservice

import (
	"go_zp_final_project/internal/domain/entities"
	"go_zp_final_project/internal/domain/ports"
)

type TaskService interface {
	GetTask() entities.Task
}

type AnalyticService struct {
	storage ports.Storage
	ts      TaskService
}

func NewAnalyticService(storage ports.Storage, ts TaskService) *AnalyticService {
	return &AnalyticService{storage: storage, ts: ts}
}
