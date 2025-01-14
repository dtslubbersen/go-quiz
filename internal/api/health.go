package api

import (
	"net/http"
)

func (a *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"version": version,
	}

	if err := a.dataResponse(w, http.StatusOK, data); err != nil {
		a.internalServerError(w, r, err)
	}
}
