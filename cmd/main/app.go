package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"square-service/internal/config"
	"square-service/internal/handlers"
	task "square-service/internal/task"
	"square-service/pkg/database"
	"square-service/pkg/logging"
)

func main() {

	router := echo.New()
	logger := logging.GetLogger()
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatal("error while loading config", zap.Error(err))
	}

	mongoDB, err := database.NewMongoClient(context.Background(), "")
	if err != nil {
		return
	}

	task.NewStorage(mongoDB, "tasks", logger)
	//storage.Create()

	router.Use(mw.RequestLoggerWithConfig(mw.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(ctx echo.Context, v mw.RequestLoggerValues) error {
			//lgr, _ := logging.GetLogger()
			logger.Info("",
				zap.String("URI", ctx.Request().URL.String()),
				zap.Int("status_code", ctx.Response().Status),
				zap.String("method", ctx.Request().Method))
			return nil
		},
	}))
	router.Use(mw.Recover())

	handler := task.NewHandler(logger)
	handler.Register(router)

	router.GET("/:name", handlers.IndexHandler)

	address := fmt.Sprintf(":%s", cfg.Port)
	log.Fatal(router.Start(address))

	//server := http.Server{
	//	Addr:         ":8080",
	//	Handler:      router,
	//	ReadTimeout:  5 * time.Second,
	//	WriteTimeout: 5 * time.Second,
	//}
	//
	//server.Serve()
	//
	//log.Fatal(server.ListenAndServe())

}
