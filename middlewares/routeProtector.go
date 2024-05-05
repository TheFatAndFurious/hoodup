package middlewares

import (
	"net/http"

	authJwt "goserver.com/auth"
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
            http.Redirect(w, r, "/login", http.StatusFound)
            return
        }
  
		err := authJwt.VerifyToken(token)
		if err != nil {
            http.Redirect(w, r, "/login", http.StatusFound)
            return
        }
		
		next.ServeHTTP(w, r)
		})	
	}
	
	
