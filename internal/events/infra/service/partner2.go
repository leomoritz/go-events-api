package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Partner2 struct {
	BaseURL string
}

type Partner2ReservationRequest struct {
	Lugares      []string `json:"lugares"`
	TipoIngresso string   `json:"tipo_ingresso"`
	Email        string   `json:"email"`
}

type Partner2ReservationResponse struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	Lugar        string `json:"spot"`
	TipoIngresso string `json:"tipo_ingresso"`
	Status       string `json:"status"`
	EventID      string `json:"evento_id"`
}

func (p *Partner2) MakeReservation(req *ReservationRequest) ([]ReservationResponse, error) {
	// Convertendo o request para o padrão do request do Partner 2
	partnerRequest := Partner2ReservationRequest{
		Lugares:      req.Spots,
		TipoIngresso: req.TicketType,
		Email:        req.Email,
	}

	// Converte o request para JSON
	body, err := json.Marshal(partnerRequest)

	if err != nil {
		return nil, err
	}

	// Monta URL do endpoint do parceiro 2
	url := fmt.Sprintf("%s/events/%s/reserve", p.BaseURL, req.EventID)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	// Realiza a requisição HTTP com o parceiro 2
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
	var partnerResp []Partner2ReservationResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&partnerResp); err != nil {
		return nil, err
	}

	// Tenta converter o response para o formato genérico do projeto
	responses := make([]ReservationResponse, len(partnerResp))
	for i, r := range partnerResp {
		responses[i] = ReservationResponse{
			ID:         r.ID,
			Email:      r.Email,
			Spot:       r.Lugar,
			TicketType: r.TipoIngresso,
			Status:     r.Status,
			EventID:    r.EventID,
		}
	}

	return responses, nil
}
