package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"sync"

	"github.com/VladiTNT/DiscordClone/web/templates"
)

var errBadBuffer = errors.New("failed to retrive buffer from buffer pool")

type PageHandler struct {
	Templates *template.Template
	Logger    *slog.Logger
	bufPool   sync.Pool
}

func NewPageHandler(t *template.Template, l *slog.Logger) *PageHandler {
	return &PageHandler{
		Templates: t, Logger: l,
		bufPool: sync.Pool{
			New: func() any { return new(bytes.Buffer) },
		},
	}
}

func (ph *PageHandler) bufferHtml(w io.Writer, name string, data any) error {
	buf, ok := ph.bufPool.Get().(*bytes.Buffer)
	defer ph.bufPool.Put(buf)
	if !ok {
		return errBadBuffer
	}

	buf.Reset()

	if err := ph.Templates.ExecuteTemplate(buf, name, data); err != nil {
		return err
	}

	_, err := buf.WriteTo(w)
	return err
}

func (ph *PageHandler) pageErr(w http.ResponseWriter, err_ error, status int) {
	ph.Logger.Error(err_.Error(), "Status", status)
	w.WriteHeader(status)
	if err := ph.bufferHtml(w, "ErrorPage", templates.Page{
		Top: templates.PageTop{
			Title: fmt.Sprintf("Error %d", status),
		},
		Content: map[string]any{
			"Status": status,
			"Error":  err_.Error(),
		},
		Bottom: templates.PageBottom{},
	}); err != nil {
		ph.Logger.Error(err.Error())
	}
}

func (ph *PageHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	if err := ph.bufferHtml(w, "HomePage", templates.Page{
		Top: templates.PageTop{
			Title: "Home",
		},
		Content: map[string]any{
			"Name": "VladiTNT",
		},
		Bottom: templates.PageBottom{},
	}); err != nil {
		ph.pageErr(w, err, http.StatusInternalServerError)
	}
}
