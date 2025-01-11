package api

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-quiz/internal/store"
	"net/http"
	"time"
)

type CreateTokenPayload struct {
	Email    string `json:"email" validate:"required,email,max=256" example:"demo@quiz.com"`
	Password string `json:"password" validate:"required,min=8,max=64" example:"password"`
}

// CreateToken godoc
//
//	@Summary		Generates an authentication token
//	@Description	Creates a JWT token for a user after validating their credentials
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		CreateTokenPayload	true	"User credentials"
//	@Success		201		{string}	string				"JWT token"
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/auth/token [post]
func (a *application) createTokenHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateTokenPayload

	if err := readJson(w, r, payload); err != nil {
		a.badRequest(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		a.badRequest(w, r, err)
		return
	}

	user, err := a.store.Users.GetByEmail(payload.Email)

	if err != nil {
		switch {
		case errors.Is(err, store.NotFoundError):
			a.unauthorized(w, r, err)
		default:
			a.internalServerError(w, r, err)
		}
		return
	}

	if err := user.Password.Compare(payload.Password); err != nil {
		a.unauthorized(w, r, err)
		return
	}

	claims := jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(a.configuration.authentication.expireAfter).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": a.configuration.authentication.iss,
		"aud": a.configuration.authentication.iss,
	}

	token, err := a.authenticator.GenerateToken(claims)
	if err != nil {
		a.internalServerError(w, r, err)
		return
	}

	if err := a.writeDataResponse(w, http.StatusCreated, token); err != nil {
		a.internalServerError(w, r, err)
	}
}
