package authentication

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const secret = "wallet-core-service"

// JWTAuthenticateMiddleware é um middleware do Echo para autenticar tokens JWT.
func JWTAuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := extractToken(r) 
		if err != nil { 
			respondWithError(w, http.StatusUnauthorized, "Authentication failed", err)
		}

		claims, err := ParseToken(token)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
			return
		}
		
		_, found := claims["sub"].(string)
		if !found {
			respondWithError(w, http.StatusUnauthorized, "Claim 'sub' not found in token", nil)
			return
		}
						
		next.ServeHTTP(w, r)
	})
}

func ParseToken(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return claims, err
}

func extractToken(r *http.Request) (string, error) {
    authorizationHeader := r.Header.Get("Authorization")
    if authorizationHeader == "" {
        return "", fmt.Errorf("Missing authorization token")
    }

    parts := strings.Split(authorizationHeader, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
        return "", fmt.Errorf("Invalid token format")
    }

    return parts[1], nil
}


func respondWithError(w http.ResponseWriter, status int, message string, err error) {
    w.WriteHeader(status)
    if err != nil {
        fmt.Fprintf(w, "%s: %v", message, err)
    } else {
        fmt.Fprint(w, message)
    }
}
