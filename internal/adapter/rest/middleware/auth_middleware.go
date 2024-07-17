package middleware

import (
	"context"
	"log"
	"net/http"
	"shiharaikun/internal/usecase"
)

type AuthMiddleware struct {
	useCase usecase.UserUseCase
}

func NewAuthMiddleware(useCase usecase.UserUseCase) *AuthMiddleware {
	return &AuthMiddleware{useCase: useCase}
}

func (a *AuthMiddleware) HandleSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionCoolie, err := r.Cookie("session-id")
		if err != nil {
			log.Print(err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		user, err := a.useCase.GetUserBySessionID(r.Context(), sessionCoolie.Value)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
