package config

import (
	"io"
	"log/slog"
	"os"
	"time"
)

type Config struct {
	Host string
	Port int

	MaxHeaderBytes int

	ReadTimeout        time.Duration
	ReadHeaderTimeout  time.Duration
	WriteTimeout       time.Duration
	IdleTimeout        time.Duration
	ServerShutdownTime time.Duration

	LogWrtier io.Writer
	LogOpts   *slog.HandlerOptions
}

func Default() *Config {
	return &Config{
		Host: "localhost",
		Port: 8080,

		MaxHeaderBytes: 2 << 20,

		ReadTimeout:        10 * time.Second,
		ReadHeaderTimeout:  5 * time.Second,
		WriteTimeout:       15 * time.Second,
		IdleTimeout:        1 * time.Minute,
		ServerShutdownTime: 30 * time.Second,

		LogWrtier: os.Stdout,
		LogOpts: &slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelDebug,
		},
	}
}
