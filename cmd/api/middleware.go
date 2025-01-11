package api

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-quiz/internal/store"
	"net/http"
	"strconv"
	"strings"
)

func (a *application) JwtTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			a.unauthorized(w, r, fmt.Errorf("missing authorization header"))
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			a.unauthorized(w, r, fmt.Errorf("corrupted authorization header"))
			return
		}

		token := parts[1]
		jwtToken, err := a.authenticator.ValidateToken(token)

		if err != nil {
			a.unauthorized(w, r, err)
		}

		claims, _ := jwtToken.Claims.(jwt.MapClaims)
		userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 64)

		if err != nil {
			a.unauthorized(w, r, err)
			return
		}

		ctx := r.Context()

		user, err := a.store.Users.GetById(store.UserId(userId))
		if err != nil {
			a.unauthorized(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, userCtxKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}