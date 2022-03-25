package authentication

import (
	"net/http"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the request url is /books and doens't have an authorization header then return a 401 Unauthorized
		token := r.Header.Get("Authorization")
		if strings.HasPrefix(r.URL.Path, "/books") {
			if token != "" {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Token not found", http.StatusUnauthorized)
			}
		} else {
			next.ServeHTTP(w, r)
		}

	})
}
