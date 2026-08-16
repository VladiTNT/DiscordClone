package web

import (
	"html/template"
	"time"
)

func Funcs() template.FuncMap {
	return template.FuncMap{
		"fTime": func(t time.Time) string {
			return t.Format(time.RFC822)
		},
	}
}
