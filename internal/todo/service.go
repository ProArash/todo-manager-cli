package todo

import (
	"context"

	"gorm.io/gorm"
)

type TaskModel struct {
	gorm.Model
	Title     string
	Completed bool
}

type TaskService struct {
	db *gorm.DB
}

func ServiceInstance(db *gorm.DB) *TaskService {
	return &TaskService{db: db}
}

func (service *TaskService) CreateTask(ctx context.Context, title string) (*TaskModel, error) {
	task := &TaskModel{Title: title, Completed: true}
	if err := service.db.WithContext(ctx).Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}
