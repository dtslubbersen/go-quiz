package api

import (
	"go-quiz/internal/store"
	"net/http"
)

type userKey string

const userCtxKey userKey = "user"

func getUserFromCtx(r *http.Request) *store.User {
	user, exists := r.Context().Value(userCtxKey).(*store.User)

	if !exists {
		return &store.User{
			Id: 1,
		}
	}

	return user
}
