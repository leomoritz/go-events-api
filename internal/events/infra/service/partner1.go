package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Partner1 struct {
	BaseURL string
}

type Partner1ReservationRequest struct {
	Spots      []string `json:"spots"`
	TicketKind string   `json:"ticketKind"`
	Email      string   `json:"email"`
}

type Partner1ReservationResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Spot       string `json:"spot"`
	TicketKind string `json:"ticketKind"`
	Status     string `json:"status"`
	EventID    string `json:"eventID"`
}

func (p *Partner1) MakeReservation(req *ReservationRequest) ([]ReservationResponse, error) {
	// Convertendo o request para o padrão do request do Partner 1
	partnerRequest := Partner1ReservationRequest{
		Spots:      req.Spots,
		TicketKind: req.TicketType,
		Email:      req.Email,
	}

	// Converte o request para JSON
	body, err := json.Marshal(partnerRequest)

	if err != nil {
		return nil, err
	}

	// Monta URL do endpoint do parceiro 1
	url := fmt.Sprintf("%s/events/%s/reserve", p.BaseURL, req.EventID)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	// Realiza a requisição HTTP com o parceiro 1
	client := &http.Client{}
	httpResp, err := client.Do(httpReq)

	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close() // Fecha a requisição apenas no final do método

	if httpResp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", httpResp.StatusCode)
	}

	// Tenta converter o response de Json para Struct, caso não seja possível, retorna um erro
	var partnerResp []Partner1ReservationResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&partnerResp); err != nil {
		return nil, err
	}

	// Tenta converter o response para o formato genérico do projeto
	responses := make([]ReservationResponse, len(partnerResp))
	for i, r := range partnerResp {
		responses[i] = ReservationResponse{
			ID:         r.ID,
			Email:      r.Email,
			Spot:       r.Spot,
			TicketType: r.TicketKind,
			Status:     r.Status,
			EventID:    r.EventID,
		}
	}

	return responses, nil
}
