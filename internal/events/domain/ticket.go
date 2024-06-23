package domain

import (
	"errors"

	"github.com/google/uuid"
)

type TicketType string

const (
	TicketTypeHalf TicketType = "HALF"
	TicketTypeFull TicketType = "FULL"
)

var (
	ErrTicketPriceMustBeGreaterThanZero = errors.New("ticket price must be greater than 0")
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}

func NewTicket(event *Event, spot *Spot, ticketType TicketType) (*Ticket, error) {
	if !isValidTicketType(ticketType) {
		return nil, errors.New("invalid ticket type")
	}

	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketType: ticketType,
		Price:      event.Price,
	}

	ticket.calculatePrice()

	if errors := ticket.Validate(); errors != nil {
		return nil, errors
	}

	return ticket, nil
}

func isValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (ticket *Ticket) calculatePrice() {
	if ticket.TicketType == TicketTypeHalf {
		ticket.Price = ticket.Price / 2 // ou ticket.Price =/ 2
	}
}

func (ticket *Ticket) Validate() error {
	if ticket.Price <= 0 {
		return ErrTicketPriceMustBeGreaterThanZero
	}

	return nil
}
