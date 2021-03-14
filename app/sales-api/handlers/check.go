package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

func readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "Ok",
	}

	return json.NewEncoder(w).Encode(status)
}
