// Package middleware provides authentication middleware using JWT tokens.
// It contains functions to verify tokens and to secure HTTP routes with JWT-based authentication.
package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

// VerifyToken parses and verifies the provided JWT token string.
// It expects the signing key to be set in the environment variable "KEY".
//
// Parameters:
//   - tokenString: The JWT token string to be verified.
//
// Returns:
//   - *jwt.Token: The parsed JWT token if it's valid.
//   - error: An error if the token is invalid or the verification fails.
func VerifyToken(tokenString string) (*jwt.Token, error) {
	signingKey := []byte(os.Getenv("KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// JWTmiddleware is an HTTP middleware that checks the Authorization header
// for a valid JWT token. If the token is valid, the request is allowed to proceed,
// otherwise, it returns a 401 Unauthorized response.
//
// The token should be provided in the Authorization header as a Bearer token.
//
// Example:
//
//	Authorization: Bearer <token>
//
// Parameters:
//   - next: The next HTTP handler to call if the token is valid.
//
// Returns:
//   - http.Handler: A wrapped handler that ensures JWT authentication.
func JWTmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := VerifyToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
