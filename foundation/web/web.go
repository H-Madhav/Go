package web

import (
	"context"
	"net/http"
	"os"
	"syscall"

	"github.com/dimfeld/httptreemux/v5"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*httptreemux.ContextMux
	shutdown chan os.Signal
}

func NewApp(shutdown chan os.Signal) *App {
	app := App{
		ContextMux: httptreemux.NewContextMux(),
		shutdown:   shutdown,
	}

	return &app
}

func (a *App) Handle(method string, path string, handler Handler) {

	h := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(r.Context(), w, r); err != nil {
			a.SignalShutdown()
			return
		}
	}

	a.ContextMux.Handle(method, path, h)
}

func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
