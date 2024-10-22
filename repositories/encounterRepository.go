package repositories

import (
	"crud-repo-2/entity"

	"gorm.io/gorm"
)

type EncounterRepository interface {
	GetAll() ([]entity.Encounter, error)
	GetById(id int) (entity.Encounter, error)
	Create(encounter entity.Encounter) (entity.Encounter, error)
	Update(encounter entity.Encounter) (entity.Encounter, error)
	Delete(encounter entity.Encounter) (entity.Encounter, error)
}

type connection struct {
	db *gorm.DB
}

// parameter DB
func NewEncounterRepository(db *gorm.DB) *connection {
	return &connection{
		db: db,
	}
}

func (connection *connection) GetAll() ([]entity.Encounter, error) {
	var encounters []entity.Encounter
	result := connection.db.Find(&encounters)
	return encounters, result.Error
}

func (connection *connection) GetById(id int) (entity.Encounter, error) {
	var encounter entity.Encounter
	result := connection.db.First(&encounter, id)
	return encounter, result.Error
}

func (connection *connection) Create(encounter entity.Encounter) (entity.Encounter, error) {
	result := connection.db.Create(&encounter)
	return encounter, result.Error
}
func (connection *connection) Update(encounter entity.Encounter) (entity.Encounter, error) {
	result := connection.db.Save(&encounter)
	return encounter, result.Error
}

func (connection *connection) Delete(encounter entity.Encounter) (entity.Encounter, error) {
	result := connection.db.Delete(&encounter)
	return encounter, result.Error
}
