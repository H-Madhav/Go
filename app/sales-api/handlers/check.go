package handlers

import (
	"context"
	"net/http"

	"github.com/madhav1993/Go/foundation/web"
)

func readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "Ok",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
