package task

import (
	"context"
	"square-service/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateTaskDTO) (*Task, error) {
	return nil, nil
}
