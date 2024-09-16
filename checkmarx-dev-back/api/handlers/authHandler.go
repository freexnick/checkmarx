package handlers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"checkmarx/api/helpers"
	"checkmarx/internal/domain/entity"
	"checkmarx/internal/domain/service"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: s,
	}

}

func (ah *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user entity.User

	if err := helpers.ReadJSON(w, r, &user); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, "malformed json", nil)
		return
	}

	token, err := ah.Service.Create(&user)
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	cookie := http.Cookie{
		Name:     "session",
		Value:    token.PlainText,
		Path:     "/",
		Expires:  time.Now().Add(72 * time.Hour),
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	helpers.WriteJSON(w, http.StatusCreated, token, nil)
}

func (ah *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var user entity.User

	if err := helpers.ReadJSON(w, r, &user); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, "malformed json", nil)
		return
	}

	credentials, err := ah.Service.Login(&user)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	cookie := http.Cookie{
		Name:     "session",
		Value:    credentials.Token.PlainText,
		Path:     "/",
		Expires:  time.Now().Add(72 * time.Hour),
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	helpers.WriteJSON(w, http.StatusOK, credentials, nil)
}

type contextKey string

const userContextKey = contextKey("user")

func (ah *AuthHandler) contextSetUser(r *http.Request, user *entity.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func (ah *AuthHandler) contextGetUser(r *http.Request) *entity.User {
	user, ok := r.Context().Value(userContextKey).(*entity.User)
	if !ok {
		panic("no user value in context")
	}
	return user
}

func (ah *AuthHandler) Get(w http.ResponseWriter, r *http.Request) {
	user := ah.contextGetUser(r)
	if user == nil {
		helpers.WriteJSON(w, http.StatusUnauthorized, "user not authenticated", nil)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, user, nil)
}

func (ah *AuthHandler) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Cookie")
		cookie, err := r.Cookie("session")

		if err != nil {
			helpers.WriteJSON(w, http.StatusUnauthorized, "missing auth token", nil)
			return
		}

		session := cookie.String()
		hash := strings.Split(session, "=")[1]
		user, err := ah.Service.GetByToken(hash)
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, nil, nil)
			return
		}
		r = ah.contextSetUser(r, user)

		next.ServeHTTP(w, r)
	})
}

//.With(conf.AuthHandler.Authenticate)
