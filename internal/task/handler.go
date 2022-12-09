package task

import (
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	handlers "square-service/internal/handlers"
	"square-service/pkg/logging"
)

const (
	taskExampleURL string = "/Task/example"
	taskCreateURL  string = "/Task"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(r *echo.Echo) {
	r.GET(taskExampleURL, h.GetTaskExample)
	r.GET("/tasks", h.GetTaskList)
	r.POST(taskCreateURL, h.CreateTask)
}

func (h *handler) CreateTask(ctx echo.Context) error {
	h.logger.Info("Task creation")
	bytes, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return http.ErrBodyNotAllowed
	}

	return ctx.String(http.StatusOK, "Task created "+string(bytes))
}

func (h *handler) GetTaskList(ctx echo.Context) error {

	err := ctx.String(
		http.StatusOK,
		"list of tasks",
	)

	return err
}

func (h *handler) GetTaskExample(ctx echo.Context) error {

	err := ctx.JSON(
		http.StatusOK,
		GetExampleTask(),
	)

	return err
}
