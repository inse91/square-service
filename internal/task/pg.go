package task

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"square-service/pkg/logging"
)

type postgresDB struct {
	client *pgx.Conn
	logger *logging.Logger
}

func (p postgresDB) Create(ctx context.Context, dto *CreateTaskDTO) (id string, err error) {

	q := `INSERT
			INTO public.tasks (description, tags, priority) 
			VALUES ($1, $2, $3)
			RETURNING id`
	err = p.client.
		QueryRow(ctx, q, dto.Description, dto.Tags, dto.Priority).
		Scan(&id)

	if err != nil {
		return "", err
	}

	p.logger.Info("created task with id:", zap.String("id", id))
	return id, nil

}

func (p postgresDB) FindAll(ctx context.Context) (tasks []*Task, err error) {

	q := `
		SELECT id, description, tags, priority
			FROM public.tasks  
			WHERE isdeleted = false OR isdeleted is null 
		`

	rows, err := p.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	tasks = make([]*Task, 0, 2000)

	for rows.Next() {

		var task Task
		err = rows.Scan(&task.ID, &task.Description, &task.Tags, &task.Priority)
		if err != nil {
			p.logger.Error("error")
		}

		tasks = append(tasks, &task)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return

}

func (p postgresDB) FindTask(ctx context.Context, id string) (*Task, error) {
	return nil, errors.New("method not implemented")
}

func (p postgresDB) UpdateTask(ctx context.Context, task Task) error {
	return errors.New("method not implemented")
}

func (p postgresDB) Delete(ctx context.Context, id string) error {
	return errors.New("method not implemented")
}

func NewPostgresTaskStorage(pgConn *pgx.Conn, logger *logging.Logger) Storage {
	return &postgresDB{
		client: pgConn,
		logger: logger,
	}
}
