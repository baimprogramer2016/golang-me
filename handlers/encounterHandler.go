package handlers

import (
	"crud-repo-2/f"
	"crud-repo-2/repositories"
	"crud-repo-2/requests"
	"crud-repo-2/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type encounterHandler struct {
	encounterService services.EncounterService
}

func NewEncounterHandler(db *gorm.DB) *encounterHandler {
	encounterRepository := repositories.NewEncounterRepository(db)
	encounterService := services.NewEncounterService(encounterRepository)
	return &encounterHandler{
		encounterService: encounterService,
	}
}

func (h *encounterHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {

	encounters, err := h.encounterService.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	f.WriteToJson(w, r, encounters)
}

func (h *encounterHandler) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idString, ok := vars["id"]
	if !ok {
		http.Error(w, "ID not found", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(idString)
	encounter, err := h.encounterService.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	f.WriteToJson(w, r, encounter)
}

func (h *encounterHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var validate *validator.Validate
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	//terima request start
	var encouterRequest requests.EncounterRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&encouterRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//terima request end

	//Validasi Start
	validate = validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(encouterRequest)

	if err != nil {
		f.WriteToJsonError(w, r, f.ErrorValidation(err))
		return
	}
	//Validasi End

	//service simpan
	encounter, err := h.encounterService.Create(encouterRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//service end

	//response
	f.WriteToJson(w, r, encounter)
}

func (h *encounterHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var validate *validator.Validate
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var encouterRequest requests.EncounterRequest

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&encouterRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	validate = validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(encouterRequest)
	if err != nil {
		f.WriteToJsonError(w, r, f.ErrorValidation(err))
		return
	}

	encounter, err := h.encounterService.Update(id, encouterRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	f.WriteToJson(w, r, encounter)

}

func (h *encounterHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var vars = mux.Vars(r)
	var idString string = vars["id"]

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id == 0 {
		http.Error(w, "ID Not Found", http.StatusBadRequest)
		return
	}

	encounter, err := h.encounterService.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	f.WriteToJson(w, r, encounter)

}
