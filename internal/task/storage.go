package task

import "context"

type Storage interface {
	Create(ctx context.Context, task Task) (string, error)
	FindAll(ctx context.Context) ([]*Task, error)
	FindTask(ctx context.Context, id string) (*Task, error)
	UpdateTask(ctx context.Context, task Task) error
	Delete(ctx context.Context, id string) error
}
