package api

import (
	"go-quiz/internal/store"
	"net/http"
)

type userKey string

const userCtxKey userKey = "user"

func getUserFromCtx(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtxKey).(*store.User)
	return user
}
