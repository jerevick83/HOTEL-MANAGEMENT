package models

import (
	"encoding/json"
	"io"
)

// Reservation holds reservation data
type Reservation struct {
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	ArrivalDate   string `json:"arrival-date"`
	DepartureDate string `json:"departure-date"`
}

func (p *Reservation) ToJson(w io.Writer) error {
	encodedData := json.NewEncoder(w)
	return encodedData.Encode(p)
}

func (p *Reservation) FromJson(w io.Reader) error {
	decodedData := json.NewDecoder(w)
	return decodedData.Decode(p)
}
