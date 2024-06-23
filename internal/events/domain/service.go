package domain

import (
	"errors"
	"fmt"
)

type spotService struct{}

func NewSpotService() *spotService {
	return &spotService{}
}

var (
	ErrInvalidQuantity = errors.New("quantity must be greater than 0")
)

func (s *spotService) generateEvents(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	/**
	*'A' + i / 10: Esta expressão calcula um caractere começando de 'A'.
	* O valor de 'A' é 65 na tabela ASCII. Ao adicionar i / 10 a 'A', obtemos
	* um caractere deslocado de 'A'. Por exemplo, se i é 0 a 9, i / 10 será 0,
	* então a expressão resultará em 'A'. Se i é 10 a 19, i / 10 será 1, resultando
	* em 'B', e assim por diante.

	* i % 10 + 1: Esta expressão calcula um número de 1 a 10. O operador % (módulo)
	* obtém o restante da divisão de i por 10. Adicionando 1, o valor resultante será
	* no intervalo de 1 a 10.
	 */
	for i := range quantity {
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1)
		spot, err := NewSpot(event, spotName)

		if err != nil {
			return err
		}
		event.Spots = append(event.Spots, *spot)
	}

	return nil
}
