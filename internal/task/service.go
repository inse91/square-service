package task

import (
	"context"
	"square-service/pkg/logging"
)

type service struct {
	storage Storage
	logger  *logging.Logger
}

func NewService(storage Storage, logger *logging.Logger) *service {
	return &service{
		storage: storage,
		logger:  logger,
	}
}

func (s *service) UpdateOne(ctx context.Context, task *Task) error {

	err := s.storage.UpdateTask(ctx, *task)
	if err != nil {
		return err
	}

	return nil

}

func (s *service) DeleteOne(ctx context.Context, id string) error {

	err := s.storage.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil

}

func (s *service) Create(ctx context.Context, dto *CreateTaskDTO) (string, error) {

	id, err := s.storage.Create(ctx, dto)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *service) GetAll(ctx context.Context) (tasks []*Task, err error) {
	tasks, err = s.storage.FindAll(ctx)
	return
}

func (s *service) GetOne(ctx context.Context, id string) (task *Task, err error) {

	task, err = s.storage.FindTask(ctx, id)
	if err != nil {
		return nil, err
	}

	return

}
