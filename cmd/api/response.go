package api

import "net/http"

func (a *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusInternalServerError, "the server encountered an error")
}

func (a *application) forbidden(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("forbidden", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusForbidden, "forbidden")
}

func (a *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("bad request", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusBadRequest, err)
}

func (a *application) notFound(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("not found", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusNotFound, "not found")
}

func (a *application) unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("unauthorized", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusUnauthorized, "unauthorized")
}
