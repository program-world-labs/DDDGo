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
	var l pwlogger.Interface
	// Logger
	if cfg.Env.EnvName != "dev" {
		l = pwlogger.NewProductionLogger(cfg.Log.Project)
	} else {
		l = pwlogger.NewDevelopmentLogger(cfg.Log.Project)
	}

	// Tracer
	err := operations.InitNewTracer(cfg.Telemetry.Host, cfg.Telemetry.Port, cfg.Telemetry.Batcher, cfg.Telemetry.SampleRate, cfg.Telemetry.Enabled)
	if err != nil {
		l.Panic().Err(err).Str("Tracer", "Run").Msg("InitNewTracer error")
	}

	// Http Server
	httpServer, err := NewHTTPServer(cfg, l)
	if err != nil {
		l.Panic().Err(err).Str("app", "Run").Msg("InitializeHTTPServer error")
	}

	// Message Router
	router, err := NewMessageRouter(cfg, l)
	if err != nil {
		l.Panic().Err(err).Str("app", "Run").Msg("InitializeMessageRouter error")
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

	// Close message server
	router.Close()
}
