package requests

type EncounterRequest struct {
	// ID       int    `json:"id" binding:"required"`
	Nama     string `json:"nama" validate:"required"`
	Poli     string `json:"poli" validate:"required"`
	Diagnosa string `json:"diagnosa" validate:"required"`
	Umur     int    `json:"umur" validate:"required"`
}
