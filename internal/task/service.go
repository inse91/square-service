package task

import (
	"context"
	errs "github.com/pkg/errors"
	"square-service/pkg/logging"
)

type service struct {
	storage Storage
	logger  logging.Logger
}

func NewService(storage Storage, logger logging.Logger) *service {
	return &service{
		storage: storage,
		logger:  logger,
	}
}

func (s *service) Create(ctx context.Context, dto *CreateTaskDTO) (id string, err error) {

	id, err = s.storage.Create(ctx, dto)
	if err != nil {
		s.logger.Errorf("failed creating task: %s", err.Error())
		err = errs.Wrap(err, "failed creating task")
		return
	}

	s.logger.Infof("successfully created tasks with id: %s", id)
	return
}

func (s *service) GetAll(ctx context.Context) (tasks []*Task, err error) {

	tasks, err = s.storage.FindAll(ctx)
	if err != nil {
		s.logger.Errorf("failed getting tasks from storage: %s", err.Error())
		err = errs.Wrap(err, "failed getting tasks from storage")
	}

	s.logger.Infof("successfully got %d tasks from storage", len(tasks))
	return
}

// TODO - proper logging and errors handling below ⬇⬇⬇

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

func (s *service) GetOne(ctx context.Context, id string) (task *Task, err error) {

	task, err = s.storage.FindTask(ctx, id)
	if err != nil {
		return nil, err
	}

	return

}
