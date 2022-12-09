package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler interface {
	Register(r *echo.Echo)
}

func IndexHandler(ctx echo.Context) error {

	name := ctx.Param("name")
	greetingTemplate := "Hello, %s!"

	err := ctx.String(
		http.StatusOK,
		fmt.Sprintf(greetingTemplate, name),
	)

	return err

}
