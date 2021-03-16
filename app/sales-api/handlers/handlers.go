package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/madhav1993/Go/business/mid"
	"github.com/madhav1993/Go/foundation/web"
)

func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log))

	app.Handle(http.MethodGet, "/readiness", readiness)
	return app
}
