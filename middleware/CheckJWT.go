package middleware

import (
	"crud-repo-2/f"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var SecretKey = []byte("rahasia-banget")

func CheckJWTToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		excludedPaths := []string{"login", "token-statis"}
		fmt.Println(r.URL.Path)
		// Cek apakah URL saat ini ada dalam daftar pengecualian
		for _, path := range excludedPaths {
			path_split := strings.Split(r.URL.Path, "/")
			if path_split[2] == path {
				next.ServeHTTP(w, r) // Lanjutkan tanpa pengecekan token
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			f.WriteToJsonError(w, r, "Missing authorization header")
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := VerifyToken(tokenString)
		if err != nil {
			f.WriteToJsonError(w, r, err.Error())
			return
		}
		next.ServeHTTP(w, r)

	})
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
