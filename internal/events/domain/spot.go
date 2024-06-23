package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrSpotNameIsRequired   = errors.New("spot name is required")
	ErrSpotInvalidName      = errors.New("spot name must be at least 2 characters long")
	ErrSpotNotFound         = errors.New("spot not found")
	ErrSoptAlreadyReserved  = errors.New("spot already reserved")
	ErrSpotStatusInvalid    = errors.New("spot status must be either AVAILABLE or SOLD")
	ErrSpotTicketIsRequired = errors.New("ticket id is required")
	ErrSpotNameStartLetter  = errors.New("spot name must start with an uppercase letter")
	ErrSpotNameEndNumber    = errors.New("spot name must end with a number")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "AVAILABLE"
	SpotStatusSold      SpotStatus = "SOLD"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	// Verifica se o erro obtido do validador é diferente de nulo, se sim, retorna o erro
	if err := spot.Validate(); err != nil {
		return nil, err // retorna um spot vazio e o erro obtido do validador
	}

	// retorna o novo spot e nil (nenhum erro)
	return spot, nil
}

func (s *Spot) Validate() error {
	if len(s.Name) == 0 {
		return ErrSpotNameIsRequired
	}
	if len(s.Name) < 2 {
		return ErrSpotInvalidName
	}

	if s.Status != SpotStatusAvailable && s.Status != SpotStatusSold {
		return ErrSpotStatusInvalid
	}

	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameStartLetter
	}
	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotNameEndNumber
	}

	return nil // retorna vazio se não houver erros
}

func (s *Spot) Reserve(ticketID string) error {
	if s.Status == SpotStatusSold {
		return ErrSoptAlreadyReserved
	}

	if ticketID == "" {
		return ErrSpotTicketIsRequired
	}

	s.Status = SpotStatusSold
	s.TicketID = ticketID

	return nil
}
