package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	infrastructure "alchemypdf.api/lib/alchemypdf.api.infrastructure"
	ioc "alchemypdf.api/lib/alchemypdf.api.ioc"
	model "alchemypdf.api/lib/alchemypdf.api.model"
	"alchemypdf.api/services/alchemypdf.api.app/app"
)

func main() {
	infra := infrastructure.NewInfrastructure()
	iocContainer := ioc.NewIoc(infra, model.NewModel(infra.DBConn()))
	app := app.NewApi(infra, iocContainer).Build()

	run(infra, iocContainer, app)
}

func run(infra infrastructure.IInfrastructure, _ ioc.IIoc, api app.IApi) {
	// Start server
	go func() {
		if err := api.Start(); err != nil && err != http.ErrServerClosed {
			infra.Logger().Fatal(fmt.Sprintf("shutting down the server: %v", err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := api.Stop(ctx); err != nil {
		infra.Logger().Fatal(err.Error())
	}
}
