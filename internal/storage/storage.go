package storage

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

const (
	TaskListLimit     = 30
	defaultDBFilePath = "./../database/scheduler.db"
	envVarDBFile      = "TODO_DBFILE"
)

func DBFilePath(log *slog.Logger) (string, error) {

	filePath := os.Getenv(envVarDBFile)
	if filePath != "" {
		log.Debug(fmt.Sprintf("Using DB file path from environment variable %s", envVarDBFile))
		return filePath, nil
	}

	currentDir, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable file directory: %w", err)
	}

	log.Debug("Using default DB file path")

	return filepath.Join(filepath.Dir(currentDir), defaultDBFilePath), nil

}
