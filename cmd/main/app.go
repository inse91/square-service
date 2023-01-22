package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"square-service/internal/config"
	"square-service/internal/task"
	"square-service/pkg/database"
	"square-service/pkg/logging"
	"syscall"
)

func main() {

	router := echo.New()
	logger := logging.GetLogger()
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatal("failed loading config", zap.Error(err))
	}

	ctx := context.Background()

	pgClient, err := database.NewPostgresClient(ctx, cfg.PostgresDB)
	if err != nil {
		logger.Fatal("", zap.Error(err))
		return
	}

	taskStorage := task.NewPostgresTaskStorage(pgClient, logger)

	router.Use(
		mw.Recover(),
		mw.RequestLoggerWithConfig(mw.RequestLoggerConfig{
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
		}),
	)

	service := task.NewService(taskStorage, logger)

	handler := task.NewHandler(service, logger)
	handler.Register(router)

	address := fmt.Sprintf(":%s", cfg.Port)

	// start server
	go func() {
		err = router.Start(address)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed starting server %v", err))
		}
	}()

	// shut server down
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	sig := <-interrupt
	logger.Info(fmt.Sprintf("shutting down server: %#v", sig))

	err = router.Shutdown(ctx)
	if err != nil {
		logger.Fatal("failed shutting server down")
	}

}
