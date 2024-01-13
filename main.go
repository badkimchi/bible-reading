package main

import (
	"app/internal/logger"
	"app/internal/server"
	"app/internal/tools"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	port := tools.EnvPortOr("3000")

	logger.Stdout.Info("starting server on port " + port[1:])

	// start listening on port
	if err := server.StartServer(port); err != nil {
		logger.Stderr.Error("server exited with error", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
