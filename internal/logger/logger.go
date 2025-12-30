package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func New(fileName string) (*log.Logger, error) {
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, err
	}

	fileFlags := os.O_CREATE | os.O_WRONLY | os.O_APPEND
	file, err := os.OpenFile("logs/"+fileName, fileFlags, 0644)
	if err != nil {
		return nil, err
	}

	logger := log.New()

	logger.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetLevel(log.InfoLevel)
	logger.SetOutput(io.MultiWriter(os.Stdout, file))

	return logger, nil
}
