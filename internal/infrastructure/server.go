package infrastructure

import (
	"context"
	"fmt"
	"go-fiber-template/lib/config"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

var (
	server *fiber.App
)

func Run() {
	server = fiber.New(config.FiberCfg(cfg))

	// Middleware
	server.Use(fiberzerolog.New(config.FiberZerologCfg(cfg)))
	server.Use(recover.New())
	server.Use(cors.New(config.CorsCfg))

	// Routes
	registerRoutes(server)

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// You can setup kafka in here if needed

	go func() {
		log.Info().Msgf("Server is running on port %s", cfg.Port)
		if err := server.Listen(fmt.Sprintf(":%s", cfg.Port)); err != nil {
			log.Error().Err(err).Msg("Failed to start server")
		}
	}()

	waitForShutdownSignal(cancel)
}

func waitForShutdownSignal(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Info().Msg("Received shutdown signal")
	shutdown(cancel)
}

func shutdown(cancel context.CancelFunc) {
	cancel()
	time.Sleep(1 * time.Second)
	log.Info().Msg("Shutting down server...")
	if err := server.ShutdownWithTimeout(3 * time.Second); err != nil {
		log.Error().Err(err).Msg("Failed to shutdown server")
	}
	cleanupResources()
	log.Info().Msg("Server shutdown complete")
}

func cleanupResources() {
	if err := dbInstance.Close(); err != nil {
		log.Error().Err(err).Msg("Failed to close database connection")
	}
}
