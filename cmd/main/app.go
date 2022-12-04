package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"square-service/internal/_handlers"
	task "square-service/internal/_task"
)

func main() {

	echoRouter := echo.New()

	echoRouter.Use(middleware.Logger())
	echoRouter.Use(middleware.Recover())

	handler := task.NewHandler()
	handler.Register(echoRouter)

	echoRouter.GET("/:name", handlers.IndexHandler)

	log.Fatal(echoRouter.Start(":8080"))

	//server := http.Server{
	//	Addr:         ":8080",
	//	Handler:      echoRouter,
	//	ReadTimeout:  5 * time.Second,
	//	WriteTimeout: 5 * time.Second,
	//}
	//
	//server.Serve()
	//
	//log.Fatal(server.ListenAndServe())

}
