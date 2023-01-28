package app

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	errs "github.com/pkg/errors"
	"os"
	"os/signal"
	"square-service/internal/config"
	"square-service/internal/task"
	"square-service/pkg/database"
	"square-service/pkg/logging"
	"syscall"
)

func Start() error {

	logger := logging.GetStdLogger()

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Errorf("failed getting config: %v", err.Error())
		return errs.Wrap(err, "failed getting config")
	}
	logger.Infof("got config %#v", cfg)

	ctx := context.Background()
	pgClient, err := database.NewPostgresClient(ctx, cfg.PostgresDB)
	if err != nil {
		logger.Errorf("failed getting database client: %s", err.Error())
		return errs.Wrap(err, "failed getting database client")
	}
	logger.Infof("got pg client")

	// layers
	taskStorage := task.NewPostgresTaskStorage(pgClient, logger)
	service := task.NewService(taskStorage, logger)
	handler := task.NewHandler(service, logger)

	// new router
	router := echo.New()
	router.Use(
		mw.Recover(),
		mw.LoggerWithConfig(mw.LoggerConfig{
			Format: "${uri} :: ${status} :: ${method} :: ${latency_human} :: ${time_rfc3339} :: ${error}\n",
			Output: os.Stdout,
		}),
	)
	handler.Register(router)

	address := fmt.Sprintf(":%s", cfg.Port)

	// start server
	go func() {
		logger.Infof("starting server")
		err = router.Start(address)
		if err != nil {
			logger.Errorf("failed starting server: %v", err)
		}
	}()

	// shut server down
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	sig := <-interrupt

	logger.Infof("shutting down server: %#v", sig)
	err = router.Shutdown(ctx)
	if err != nil {
		logger.Errorf("failed shutting server down: %v", err)
		return errs.Wrap(err, "failed shutting server down")
	}

	return nil

}
