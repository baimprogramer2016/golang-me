package main

import (
	"crud-repo-2/database"
	"crud-repo-2/entity"
	"crud-repo-2/f"
	"crud-repo-2/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
- API
- Middleware
- JWT
- Crud
- Interface
*/

const USERNAME = "root"
const PASSWORD = "pass"

func main() {

	db, err := database.ConnectMysql()
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&entity.Encounter{})

	mux := mux.NewRouter()
	v1 := mux.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).Methods("GET")

	//encounter
	encounterHandler := handlers.NewEncounterHandler(db)
	v1.HandleFunc("/encounters", encounterHandler.GetAllHandler).Methods("GET")
	v1.HandleFunc("/encounter/{id}", encounterHandler.GetByIdHandler).Methods("GET")
	v1.HandleFunc("/encounter", encounterHandler.CreateHandler).Methods("POST")
	v1.HandleFunc("/encounter/{id}", encounterHandler.UpdateHandler).Methods("PUT")
	v1.HandleFunc("/encounter/{id}", encounterHandler.DeleteHandler).Methods("DELETE")

	//token dan Middleware
	tokenStatisHandler := handlers.NewTokenStatisHandler()
	v1.HandleFunc("/token-statis", tokenStatisHandler.GetTokenValue).Methods("GET")
	v1.HandleFunc("/check-role", tokenStatisHandler.GetTokenValue).Methods("GET")

	//middleware
	var handler http.Handler = mux
	handler = MiddlewareAuth(handler)
	handler = RoleCheckMiddleware("admin", "manager")(handler)

	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = handler
	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}

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

func RoleCheckMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Misalnya, role pengguna disimpan di header Authorization (hanya contoh)
			userRole := r.Header.Get("Role")

			// Memeriksa apakah role pengguna ada di daftar allowedRoles
			for _, role := range allowedRoles {
				if userRole == role {
					next.ServeHTTP(w, r) // Lanjutkan ke handler berikutnya
					return
				}
			}

			// Jika role tidak cocok, return 403 Forbidden
			f.WriteToJsonError(w, r, "Forbidden")
		})
	}
}
