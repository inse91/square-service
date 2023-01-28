package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	errs "github.com/pkg/errors"
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
	logger  logging.Logger
	service *service
}

func NewHandler(service *service, logger logging.Logger) handlers.Handler {
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

func (h *handler) CreateTask(ctx echo.Context) error {

	h.logger.Infof("creating new task")

	var dto *CreateTaskDTO = nil
	d := json.NewDecoder(ctx.Request().Body)
	err := d.Decode(&dto)
	if err != nil {
		h.logger.Errorf("failed decoding json body: %s", err.Error())
		_ = ctx.String(http.StatusBadRequest, "failed decoding json body")
		return errs.Wrap(err, "failed decoding json body")
	}

	id, err := h.service.Create(ctx.Request().Context(), dto)
	if err != nil {
		h.logger.Errorf("failed decoding json body: %s", err.Error())
		_ = ctx.String(http.StatusInternalServerError, "failed creating task")
		return errs.Wrap(err, "failed decoding json body")
	}

	err = ctx.String(http.StatusCreated, fmt.Sprintf("task created with id: %s", id))
	if err != nil {
		h.logger.Errorf("failed writing to response: %s", err.Error())
		return errs.Wrap(err, "failed writing to response")
	}

	h.logger.Infof("task created with id: %s", id)
	return nil

}

func (h *handler) GetTaskList(ctx echo.Context) error {

	tasks, err := h.service.GetAll(ctx.Request().Context())
	if err != nil {
		h.logger.Errorf("failed getting tasks: %s", err.Error())
		_ = ctx.String(http.StatusInternalServerError, "failed getting tasks")
		return errs.Wrap(err, "failed getting tasks")
	}

	err = ctx.JSON(http.StatusOK, tasks)
	if err != nil {
		h.logger.Errorf("failed writing to response: %s", err.Error())
		return errs.Wrap(err, "failed writing to response")
	}

	return nil

}

// TODO - proper logging and errors handling below ⬇⬇⬇

func (h *handler) UpdateTask(ctx echo.Context) error {

	h.logger.Infof("updating task")

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
		return ctx.String(http.StatusInternalServerError, "failed updating task"+err.Error())
	}

	err = ctx.String(http.StatusAccepted, "task updated")
	if err != nil {
		return err
	}

	return nil

}

func (h *handler) DeleteTaskById(ctx echo.Context) error {

	h.logger.Infof("deleting task by id")
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
	return err
}

func (h *handler) GetTaskById(ctx echo.Context) error {

	id := ctx.Param("id")
	h.logger.Infof("Getting task by id: %s", id)
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
		_ = ctx.String(
			http.StatusInternalServerError,
			fmt.Sprintf("failed"), // todo - message??
		)
		h.logger.Infof("failed serializing task: %v", task)
		return errs.Wrap(err, "failed serializing task")
	}

	return err

}

func (h *handler) GetTaskExample(ctx echo.Context) error {

	err := ctx.JSON(http.StatusOK, GetExampleTask())
	if err != nil {
		_ = ctx.String(http.StatusInternalServerError, "failed sending example")
		return errs.Wrap(err, "failed sending example")
	}

	return nil

}
