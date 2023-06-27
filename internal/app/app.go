// Package app configures and runs application.
package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/program-world-labs/pwlogger"

	"github.com/program-world-labs/DDDGo/config"
	"github.com/program-world-labs/DDDGo/pkg/operations"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	// Tracer
	operations.GoogleCloudOperationInit(cfg.GCP.Project, cfg.GCP.Monitor)

	var l pwlogger.Interface
	// Logger
	if cfg.Env.EnvName != "dev" {
		l = pwlogger.NewProductionLogger(cfg.GCP.Project)
	} else {
		l = pwlogger.NewDevelopmentLogger(cfg.GCP.Project)
	}

	httpServer, err := NewHTTPServer(cfg, l)
	if err != nil {
		l.Err(err).Str("app", "Run").Msg("InitializeHTTPServer error")
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info().Str("app", "Run").Msgf("Got signal %s, exiting now", s.String())
	case err = <-httpServer.Notify():
		l.Err(err).Str("app", "Run").Msg("httpServer.Notify error")
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Err(err).Str("app", "Run").Msg("httpServer.Shutdown error")
	}
}
