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
	"square-service/internal/task"
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

	ctx := context.Background()

	mongoDB, err := database.NewMongoClient(ctx, cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.DataBase)
	if err != nil {
		logger.Fatal("", zap.Error(err))
		return
	}

	storage := task.NewStorage(mongoDB, cfg.MongoDB.Collection, logger)

	var tasks []*task.Task
	tasks, err = storage.FindAll(ctx)
	if err != nil {
		logger.Fatal("error updating task", zap.Error(err))
		return
	}

	fmt.Printf("%#v", &tasks)

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

}
