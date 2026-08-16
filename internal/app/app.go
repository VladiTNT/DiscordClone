package app

import (
	"context"
	"fmt"
	"html/template"
	"log/slog"
	"net"
	"net/http"
	"strconv"

	"github.com/VladiTNT/DiscordClone/internal/config"
	"github.com/VladiTNT/DiscordClone/internal/middleware"
	"github.com/VladiTNT/DiscordClone/web"
)

type App struct {
	Config    *config.Config
	Logger    *slog.Logger
	Templates *template.Template

	Router      *http.ServeMux
	Middlewares []middleware.MiddlewareFunc
	Server      *http.Server
}

func (a *App) Init(cfg *config.Config) {
	// Config
	a.Config = cfg

	// Logger
	a.Logger = slog.New(slog.NewTextHandler(cfg.LogWrtier, cfg.LogOpts))

	// Templates
	a.Templates = template.Must(
		template.New("").
			Funcs(web.Funcs()).
			ParseFS(web.HtmlFS, "templates/*.tmpl", "pages/*.html"),
	)

	// Setup router
	a.Router = http.NewServeMux()
	a.RegisterRoutes()

	// Middleware
	a.Middlewares = []middleware.MiddlewareFunc{
		middleware.Logger,
	}

	// HTTP Server
	a.Server = new(http.Server)
	a.Server.Addr = net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
	a.Server.ReadTimeout = cfg.ReadTimeout
	a.Server.ReadHeaderTimeout = cfg.ReadHeaderTimeout
	a.Server.WriteTimeout = cfg.WriteTimeout
	a.Server.IdleTimeout = cfg.IdleTimeout
	a.Server.MaxHeaderBytes = cfg.MaxHeaderBytes
	a.Server.Handler = middleware.Stack(a.Middlewares...)(a.Router)
}

func (a *App) Run(ctx context.Context) error {
	go func() {
		fmt.Printf("HTTP server running on %s\n", a.Server.Addr)
		if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()
	fmt.Println("Shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), a.Config.ServerShutdownTime)
	defer cancel()

	if err := a.Server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	fmt.Println("Shutdown complete!")
	return nil
}
