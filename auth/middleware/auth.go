package middleware

import (
	"context"
	"net/http"

	"github.com/mstgnz/microservice/config"
)

func TokenValidate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		validateToken, err := config.ValidateToken(token)
		if err != nil {
			_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Invalid Token", Error: err.Error()})
			return
		}
		if !validateToken.Valid {
			_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Invalid Token", Error: err.Error()})
			return
		}
		ctx := context.WithValue(r.Context(), "token", validateToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
