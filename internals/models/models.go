package models

// Reservation holds reservation data
type Reservation struct {
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	ArrivalDate   string `json:"arrival-date"`
	DepartureDate string `json:"departure-date"`
}
