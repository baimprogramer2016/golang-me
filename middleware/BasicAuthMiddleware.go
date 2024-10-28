package middleware

import (
	"crud-repo-2/f"
	"net/http"
)

const USERNAME = "root"
const PASSWORD = "pass"

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			f.WriteToJsonError(w, r, "Something Wrong")
			return
		}
		isVavlid := (username == USERNAME) && (password == PASSWORD)
		if !isVavlid {
			f.WriteToJsonError(w, r, "Unathorization")
			return
		}
		next.ServeHTTP(w, r)

	})
}
