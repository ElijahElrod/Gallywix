package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/elijahelrod/vespene/config"
	"github.com/elijahelrod/vespene/internal/app"
)

// main loads the configuration, and runs the trading application
func main() {
	ctx, cancel := signal.NotifyContext(context.TODO(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-ctx.Done()
		cancel()
	}()

	// Initialize Config
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		log.Fatalf("Failed to init with config, %v", err)
	}

	app.Run(ctx, cfg)
}
