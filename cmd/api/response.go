package api

import (
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}

func (a *application) dataResponse(w http.ResponseWriter, statusCode int, data any) error {
	return writeJson(w, statusCode, &Response{StatusCode: statusCode, Data: data})
}

func (a *application) errorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) error {
	a.logger.Errorw("error", "method", r.Method, "path", r.URL.Path, "error", message)
	return writeJson(w, statusCode, &Response{StatusCode: statusCode, Error: message})
}

func (a *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusInternalServerError, &Response{StatusCode: http.StatusInternalServerError, Error: "internal server error"})
}

func (a *application) forbidden(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("forbidden", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusForbidden, &Response{StatusCode: http.StatusForbidden, Error: err.Error()})
}

func (a *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("bad request", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusBadRequest, &Response{StatusCode: http.StatusBadRequest, Error: err.Error()})
}

func (a *application) notFound(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("not found", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusNotFound, &Response{StatusCode: http.StatusNotFound, Error: err.Error()})
}

func (a *application) unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Errorw("unauthorized", "method", r.Method, "path", r.URL.Path, "error", err)
	_ = writeJson(w, http.StatusUnauthorized, &Response{StatusCode: http.StatusUnauthorized, Error: err.Error()})
}
