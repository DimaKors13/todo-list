package httpserver

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"todo-list/internal/lib/logger"
)

const (
	defaultServerPort = "7540"
	envVarPort        = "TODO_PORT"
)

func HTTPServerPort(log *slog.Logger) string {

	envPort := os.Getenv(envVarPort)
	if envPort != "" {
		_, err := strconv.ParseInt(envPort, 10, 32)
		if err != nil {
			log.Error(fmt.Sprintf("failed to parse %s as port", envVarPort), logger.Err(err))
		} else {
			log.Debug(fmt.Sprintf("Using http-server port from environment variable %s", envVarPort))
			return envPort
		}
	}

	log.Debug("Using http-server port from default value")
	return defaultServerPort

}
