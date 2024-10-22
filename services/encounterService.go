package services

import (
	"crud-repo-2/entity"
	"crud-repo-2/repositories"
	"crud-repo-2/requests"
	"crud-repo-2/responses"
)

type EncounterService interface {
	GetAll() ([]responses.EncounterResponse, error)
	GetById(id int) (entity.Encounter, error)
	Create(encounter requests.EncounterRequest) (entity.Encounter, error)
	Update(id int, encounter requests.EncounterRequest) (entity.Encounter, error)
	Delete(id int) (entity.Encounter, error)
}

type encounterServiceRepository struct {
	repository repositories.EncounterRepository
}

func NewEncounterService(repository repositories.EncounterRepository) *encounterServiceRepository {
	return &encounterServiceRepository{repository}
}

func (s *encounterServiceRepository) GetAll() ([]responses.EncounterResponse, error) {
	encounters, err := s.repository.GetAll()
	var encounterResponse []responses.EncounterResponse

	for _, encounter := range encounters {
		encounterResponse = append(encounterResponse, responses.EncounterResponse{
			ID:       encounter.ID,
			Nama:     encounter.Nama,
			Poli:     encounter.Poli,
			Diagnosa: encounter.Diagnosa,
			Umur:     encounter.Umur,
		})
	}
	return encounterResponse, err
}

func (s *encounterServiceRepository) GetById(id int) (entity.Encounter, error) {
	encounter, err := s.repository.GetById(id)
	return encounter, err
}
func (s *encounterServiceRepository) Create(encounterRequest requests.EncounterRequest) (entity.Encounter, error) {

	encounter := entity.Encounter{
		Nama:     encounterRequest.Nama,
		Poli:     encounterRequest.Poli,
		Diagnosa: encounterRequest.Diagnosa,
		Umur:     encounterRequest.Umur,
	}
	encounter, err := s.repository.Create(encounter)
	return encounter, err
}
func (s *encounterServiceRepository) Update(id int, encounterRequest requests.EncounterRequest) (entity.Encounter, error) {

	encounterResult, _ := s.repository.GetById(id)

	encounterResult.Nama = encounterRequest.Nama
	encounterResult.Poli = encounterRequest.Poli
	encounterResult.Diagnosa = encounterRequest.Diagnosa
	encounterResult.Umur = encounterRequest.Umur

	newEncounter, err := s.repository.Update(encounterResult)
	return newEncounter, err
}

func (s *encounterServiceRepository) Delete(id int) (entity.Encounter, error) {
	encounterResult, _ := s.repository.GetById(id)
	deleteEncounter, err := s.repository.Delete(encounterResult)
	return deleteEncounter, err
}
