package api

import (
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}

func (a *Application) dataResponse(w http.ResponseWriter, r *http.Request, statusCode int, data any) {
	if err := writeJson(w, statusCode, &Response{StatusCode: statusCode, Data: data}); err != nil {
		a.internalServerError(w, r, err)
	}
}

func (a *Application) errorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	a.logger.Errorw("error", "method", r.Method, "path", r.URL.Path, "error", message)

	if err := writeJson(w, statusCode, &Response{StatusCode: statusCode, Error: message}); err != nil {
		a.internalServerError(w, r, err)
	}
}

func (a *Application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusInternalServerError, &Response{StatusCode: http.StatusInternalServerError, Error: "internal server error"})
}

func (a *Application) forbidden(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("forbidden", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusForbidden, &Response{StatusCode: http.StatusForbidden, Error: err.Error()})
}

func (a *Application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("bad request", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusBadRequest, &Response{StatusCode: http.StatusBadRequest, Error: err.Error()})
}

func (a *Application) notFound(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("not found", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusNotFound, &Response{StatusCode: http.StatusNotFound, Error: err.Error()})
}

func (a *Application) unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("unauthorized", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusUnauthorized, &Response{StatusCode: http.StatusUnauthorized, Error: err.Error()})
}
