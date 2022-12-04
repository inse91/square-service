package task

import (
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	handlers "square-service/internal/_handlers"
	"square-service/pkg/logging"
)

const (
	taskExampleURL string = "/task/example"
	taskCreateURL  string = "/task"
)

type handler struct {
	//log logging.Logger
	logger *logging.Logger
}

func NewHandler(log *logging.Logger) handlers.Handler {
	return &handler{
		logger: log,
	}
}

func (h *handler) Register(r *echo.Echo) {
	r.GET(taskExampleURL, h.GetTaskExample)
	r.GET("/tasks", h.GetTaskList)
	r.POST(taskCreateURL, h.CreateTask)
}

func (h *handler) CreateTask(ctx echo.Context) error {
	h.logger.Info("task creation")
	bytes, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return http.ErrBodyNotAllowed
	}

	return ctx.String(http.StatusOK, "task created "+string(bytes))
}

func (h *handler) GetTaskList(ctx echo.Context) error {
	ctx.Set("Content-Type", "123")
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
