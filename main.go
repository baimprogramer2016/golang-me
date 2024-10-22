package main

import (
	"crud-repo-2/database"
	"crud-repo-2/entity"
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

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
