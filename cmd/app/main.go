package main

import (
	"context"
	"github.com/elijahelrod/vespene/config"
	"github.com/elijahelrod/vespene/internal/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Main runs the application
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

	log.Print("Starting Vespene")
	app.Run(ctx, cfg)
}
