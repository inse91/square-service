package main

import (
	"square-service/cmd/app"
)

func main() {
	
	app.Start()
	//router := echo.New()
	//
	//logger, err := logging.GetLogger()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//cfg, err := config.GetConfig()
	//if err != nil {
	//	logger.Fatal(err.Error())
	//}
	//
	//ctx := context.Background()
	//
	//pgClient, err := database.NewPostgresClient(ctx, cfg.PostgresDB)
	//if err != nil {
	//	logger.Fatal(err.Error())
	//	return
	//}
	//
	//taskStorage := task.NewPostgresTaskStorage(pgClient, logger)
	//
	//router.Use(
	//	mw.Recover(),
	//	mw.LoggerWithConfig(mw.LoggerConfig{
	//		Format: "${uri} :: ${status} :: ${method} :: ${latency_human} :: ${time_rfc3339} :: ${error}\n",
	//		Output: os.Stdout,
	//	}),
	//)
	//
	//service := task.NewService(taskStorage, logger)
	//handler := task.NewHandler(service, logger)
	//
	//handler.Register(router)
	//
	//address := fmt.Sprintf(":%s", cfg.Port)
	//
	//// start server
	//go func() {
	//	logger.Info("starting server")
	//	err = router.Start(address)
	//	if err != nil {
	//		logger.Fatal(err.Error())
	//	}
	//}()
	//
	//// shut server down
	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	//
	//sig := <-interrupt
	//logger.Info(fmt.Sprintf("shutting down server: %#v", sig))
	//
	//err = router.Shutdown(ctx)
	//if err != nil {
	//	logger.Fatal("failed shutting server down")
	//}

}
