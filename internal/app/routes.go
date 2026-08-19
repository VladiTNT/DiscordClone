package app

import (
	"net/http"

	"github.com/VladiTNT/DiscordClone/internal/handlers"
	"github.com/VladiTNT/DiscordClone/web"
)

func (a *App) RegisterRoutes() {
	// Static assets handler
	a.Router.Handle("GET /assets/", http.FileServerFS(web.AssetsFS))

	// Pages
	pageHandler := handlers.NewPageHandler(a.Templates, a.Logger)
	a.Router.HandleFunc("GET /", pageHandler.HomePage)
	a.Router.HandleFunc("GET /pages/signup", pageHandler.SignupPage)
	a.Router.HandleFunc("GET /pages/login", pageHandler.LoginPage)

	// Auth
	// POST /api/auth/signup
	// POST /api/auth/login
}
