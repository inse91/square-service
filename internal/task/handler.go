package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	handlers "square-service/internal/handlers"
	"square-service/pkg/logging"
)

const (
	taskExampleURL   string = "/task/example"
	taskCreateURL    string = "/task"
	taskGetByIdURL   string = "/task/:id"
	tasksGetAllURL   string = "/tasks"
	taskDeleteOneURL string = "/task/:id"
	taskUpdateOneURL string = "/task"
)

type handler struct {
	logger  *logging.Logger
	service *service
}

func NewHandler(service *service, logger *logging.Logger) handlers.Handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) Register(r *echo.Echo) {
	r.GET(taskExampleURL, h.GetTaskExample)
	r.GET(tasksGetAllURL, h.GetTaskList)
	r.POST(taskCreateURL, h.CreateTask)
	r.GET(taskGetByIdURL, h.GetTaskById)
	r.DELETE(taskDeleteOneURL, h.DeleteTaskById)
	r.PUT(taskUpdateOneURL, h.UpdateTask)
}

func (h *handler) UpdateTask(ctx echo.Context) error {

	h.logger.Info("updating task")

	var task *Task = nil
	d := json.NewDecoder(ctx.Request().Body)
	err := d.Decode(&task)
	if err != nil {
		if err := ctx.String(http.StatusBadRequest, "wrong json"); err != nil {
			return err
		}
		return err
	}

	err = h.service.UpdateOne(ctx.Request().Context(), task)
	if err != nil {
		if err := ctx.String(http.StatusInternalServerError, "failed updating task"); err != nil {
			return err
		}
		return err
	}

	err = ctx.String(http.StatusAccepted, "task updated")
	if err != nil {
		return err
	}

	return nil

}

func (h *handler) DeleteTaskById(ctx echo.Context) error {

	h.logger.Info("deleting task by id")
	id := ctx.Param("id")
	if id == "" {
		err := ctx.String(http.StatusBadRequest, "no id provided")
		if err != nil {
			err = ctx.String(http.StatusInternalServerError, "internal error")
			return err
		}
		return errors.New("no id provided")
	}

	err := h.service.DeleteOne(ctx.Request().Context(), id)
	if err != nil {
		err = ctx.String(http.StatusNotFound, fmt.Sprintf("task with id: %s not found", id))
		if err != nil {
			err = ctx.String(http.StatusInternalServerError, "internal error")
			return err
		}
		return err
	}

	err = ctx.String(http.StatusOK, "deleted successfully")

	return nil
}

func (h *handler) GetTaskById(ctx echo.Context) error {

	h.logger.Info("Getting task by id")
	id := ctx.Param("id")
	if id == "" {
		err := ctx.String(
			http.StatusBadRequest,
			fmt.Sprintf("no id in query params"),
		)
		if err != nil {
			return err
		}
		return errors.New("no id in query params")
	}

	task, err := h.service.GetOne(ctx.Request().Context(), id) // GetOne
	if err != nil {
		err := ctx.String(
			http.StatusNotFound,
			fmt.Sprintf("task not found"),
		)
		return err
	}

	err = ctx.JSON(http.StatusOK, task)
	if err != nil {
		err = ctx.String(
			http.StatusInternalServerError,
			fmt.Sprintf(""), // todo - message??
		)
		return err
	}

	return nil

}

func (h *handler) CreateTask(ctx echo.Context) error {

	h.logger.Info("creating new task")

	var dto *CreateTaskDTO = nil
	d := json.NewDecoder(ctx.Request().Body)
	err := d.Decode(&dto)
	if err != nil {
		if err := ctx.String(http.StatusBadRequest, "wrong json"); err != nil {
			return err
		}
		return err
	}

	id, err := h.service.Create(ctx.Request().Context(), dto)
	if err != nil {
		if err := ctx.String(http.StatusInternalServerError, "failed creating task"); err != nil {
			return err
		}
		return err
	}

	return ctx.String(http.StatusCreated, fmt.Sprintf("Task created with id: %s", id))

}

func (h *handler) GetTaskList(ctx echo.Context) error {

	tasks, err := h.service.GetAll(ctx.Request().Context())
	if err != nil {
		return err
	}

	err = ctx.JSON(http.StatusOK, tasks)
	if err != nil {
		return err
	}

	return nil

}

func (h *handler) GetTaskExample(ctx echo.Context) error {

	err := ctx.JSON(http.StatusOK, GetExampleTask())
	if err != nil {
		return err
	}

	return nil

}
