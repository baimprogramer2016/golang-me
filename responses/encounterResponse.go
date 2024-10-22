package responses

type EncounterResponse struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Poli     string `json:"poli"`
	Diagnosa string `json:"diagnosa"`
	Umur     int    `json:"umur"`
}
