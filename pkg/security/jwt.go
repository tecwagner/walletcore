package security

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tecwagner/walletcore-service/internal/entity"
)

const secret = "wallet-core-service"

func NewJWTToken(user *entity.Client) (string, error) {
	// Crie um token JWT com informações do usuário
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // Define a expiração do token
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

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

// ParseToken parses a token and handles token expiration
// func ParseToken(tokenString string) (*jwt.Token, error) {

// 	token, err := parseHS256(tokenString)
// 	if err != nil && err.Error() == "Token is expired" {
// 		token, err = parseHS256(tokenString)
// 	}

// 	return token, err
// }

// func parseHS256(tokenString string) (*jwt.Token, error) {
// 	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(secret), nil
// 	})
// }

// GetClaims retrieves claims information
func GetClaims(token *jwt.Token) (jwt.MapClaims, error) {
	if !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	}
	err := token.Claims.(jwt.MapClaims).Valid()
	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
