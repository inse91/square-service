package task

import (
	"context"
	"square-service/pkg/logging"
)

type service struct {
	storage Storage
	logger  *logging.Logger
}

func newService(storage Storage, logger *logging.Logger) *service {
	return &service{
		storage: storage,
		logger:  logger,
	}
}

func (s *service) Create(ctx context.Context, dto CreateTaskDTO) (*Task, error) {

	return nil, nil
}

func (s *service) GetAll(ctx context.Context) (tasks []*Task, err error) {

	tasks, err = s.storage.FindAll(ctx)

	return
}
