package main

import (
	"database/sql"
	"net/http"

	httpHandler "github.com/leomoritz/go-events-api/internal/events/infra/http"
	"github.com/leomoritz/go-events-api/internal/events/infra/repository"
	"github.com/leomoritz/go-events-api/internal/events/infra/service"
	"github.com/leomoritz/go-events-api/internal/events/usecase"
)

func main() {
	// Abrindo conexão com o banco de dados
	db, err := sql.Open("mysql", "root:root@tcp(localhost:33006)/test_db")

	if err != nil {
		panic(err)
	}

	defer db.Close() // Encerra a conexão com o banco de dados no final do programa

	// Instanciando os repositorios
	eventRepo, err := repository.NewMysqlEventRepository(db)

	if err != nil {
		panic(err)
	}

	// Setando as URLs base das APIs dos parceiros
	partnerBaseURLs := map[int]string{
		1: "http://localhost:9080/api1",
		2: "http://localhost:9080/api2",
	}

	// Criando a fábrica de parceiros com base nas URLs das APIs
	partnerFactory := service.NewPartnerFactory(partnerBaseURLs)

	// Instanciando os casos de uso para os repositorios
	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	buyTicketsUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)

	// Instanciando o handler que irá tratar as requisições HTTP
	eventsHandler := httpHandler.NewEventsHandler(
		listEventsUseCase,
		listSpotsUseCase,
		getEventUseCase,
		buyTicketsUseCase,
	)

	// Montando o servidor web com um roteador para tratar as requisições HTTP
	/**
	* NewServerMux é uma função em linguagens de programação, especialmente em Go (Golang),
	* usada para criar um novo roteador de servidor HTTP (também chamado de "mux" ou "multiplexer").
	* Um roteador HTTP é responsável por direcionar requisições HTTP para diferentes manipuladores
	* (handlers) com base na URL da requisição. Em resumo, ele monta um hub para inserir as rotas HTTP.
	 */
	r := http.NewServeMux()
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /checkout", eventsHandler.BuyTickets)

	// Iniciando o servidor
	http.ListenAndServe(":8080", r)
}
