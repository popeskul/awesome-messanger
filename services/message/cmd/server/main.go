package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/popeskul/awesome-messanger/services/message/internal/di"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	app, err := di.InitializeApp()
	if err != nil {
		app.Logger.Error("Failed to initialize application", "error", err)
		os.Exit(1)
	}

	go func() {
		if err = app.Run(); err != nil {
			app.Logger.Error("Failed to run application", "error", err)
			stop()
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err = app.Stop(shutdownCtx); err != nil {
		app.Logger.Error("Error during application shutdown", "error", err)
		os.Exit(1)
	}

	app.Logger.Info("Server exiting")
}
