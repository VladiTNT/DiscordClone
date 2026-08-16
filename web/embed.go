package web

import "embed"

//go:embed assets/*
var AssetsFS embed.FS

//go:embed templates/*.tmpl pages/*.html
var HtmlFS embed.FS
