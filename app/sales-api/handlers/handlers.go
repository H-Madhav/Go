package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/dimfeld/httptreemux/v5"
)

func API(build string, shutdown chan os.Signal, log *log.Logger) *httptreemux.ContextMux {
	tm := httptreemux.NewContextMux()

	tm.Handle(http.MethodGet, "/test", readiness)
	return tm
}
