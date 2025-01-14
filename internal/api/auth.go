package api

import (
	"errors"
	"github.com/dtslubbersen/go-quiz/internal/store"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type CreateTokenPayload struct {
	Email    string `json:"email" validate:"required,email,max=256" example:"demo@quiz.com"`
	Password string `json:"password" validate:"required,min=8,max=64" example:"password"`
}

type TokenCreatedResponse struct {
	Token     string       `json:"token"`
	ExpiresIn int64        `json:"expires_in"`
	User      UserResponse `json:"user"`
}

type UserResponse struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

// CreateToken godoc
//
//	@Summary		Generates an authentication token
//	@Description	Creates a JWT token for a user after validating their credentials
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		CreateTokenPayload	true "User credentials"
//	@Success		201		{object}	Response{data=TokenCreatedResponse}
//	@Failure		400		{object}	Response{error=string}
//	@Failure		401		{object}	Response{error=string}
//	@Failure		500		{object}	Response{error=string}
//	@Router			/auth/token [post]
func (a *Application) createTokenHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateTokenPayload

	if err := readJson(w, r, &payload); err != nil {
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
		"exp": time.Now().Add(a.cfg.authentication.expireAfter).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": a.cfg.authentication.iss,
		"aud": a.cfg.authentication.iss,
	}

	token, err := a.authenticator.GenerateToken(claims)
	if err != nil {
		a.internalServerError(w, r, err)
		return
	}

	response := TokenCreatedResponse{
		Token:     token,
		ExpiresIn: claims["exp"].(int64),
		User: UserResponse{
			Id:    int64(user.Id),
			Email: user.Email,
		},
	}

	if err := a.dataResponse(w, http.StatusCreated, response); err != nil {
		a.internalServerError(w, r, err)
	}
}
