package middlewares

import (
	"context"
	"fmt"
	"net/http"

	authJwt "goserver.com/internal/auth"
)

func ProtectRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string
		found := false
		for _, cookie := range r.Cookies() {
			if cookie.Name == "jwt" {
				if cookie.Value != "" {
					token = cookie.Value
					found = true
				}
				break
			}
		}
		if !found {
			fmt.Println("ProtectRoute: No token found")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		claims, err := authJwt.VerifyToken(token)
		if err != nil {
			fmt.Printf("ProtectRoute: Token verification failed: %v\n", err)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		fmt.Printf("ProtectRoute: Token verified, claims: %v\n", claims)
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}