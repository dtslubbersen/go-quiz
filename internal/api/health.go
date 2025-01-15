package api

import (
	"net/http"
)

func (a *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"version": version,
	}

	a.dataResponse(w, r, http.StatusOK, data)
}
