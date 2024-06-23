package service

type ReservationRequest struct {
	EventID    string   `json:"eventID"`
	Spots      []string `json:"spots"`
	TicketType string   `json:"ticketKind"`
	CardHash   string   `json:"cardHash"`
	Email      string   `json:"email"`
}

type ReservationResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Spot       string `json:"spot"`
	TicketType string `json:"ticketKind"`
	Status     string `json:"status"`
	EventID    string `json:"eventID"`
}

type Partner interface {
	MakeReservation(req *ReservationRequest) ([]ReservationResponse, error)
}
