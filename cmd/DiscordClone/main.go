package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/VladiTNT/DiscordClone/internal/app"
	"github.com/VladiTNT/DiscordClone/internal/config"
)

func main() {
	var a app.App
	a.Init(config.Default())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := a.Run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Server shutdown failed: %v\n", err)
		os.Exit(1)
	}
}
