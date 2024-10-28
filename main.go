package main

import (
	"crud-repo-2/database"
	"crud-repo-2/entity"
	"crud-repo-2/f"
	"crud-repo-2/handlers"
	"crud-repo-2/middleware"
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

	loginHandler := handlers.NewTokenJwtHandler()
	v1.HandleFunc("/login", loginHandler.LoginHandler).Methods("POST")
	v1.HandleFunc("/protected", ProtectedHandler).Methods("GET")

	//middleware
	var handler http.Handler = mux
	// handler = middleware.MiddlewareAuth(handler)
	handler = middleware.CheckJWTToken(handler)
	handler = middleware.RoleCheckMiddleware("admin", "manager")(handler)

	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = handler
	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {

	f.WriteToJson(w, r, "Welcome")

}
