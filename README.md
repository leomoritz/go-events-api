# Sistema de Venda de Ingressos

Bem-vindo ao sistema de venda de ingressos, um projeto desenvolvido durante a semana Imersão Full Stack && Full Cycle. Este projeto é uma aplicação backend escrita em Go (Golang) que fornece APIs para a reserva de ingressos e se conecta com parceiros para realizar a reserva dos assentos.

## Tecnologias Utilizadas

- **Linguagem:** Go (Golang)
- **Banco de Dados:** MySQL
- **Conexão com o Banco:** SQL puro (sem ORM)

## Funcionalidades

- **Reserva de Ingressos:** API para reservar assentos em eventos.
- **Integração com Parceiros:** Conexão com APIs de parceiros para a reserva de assentos.
- **Gerenciamento de Sessões:** Controle de disponibilidade e reservas de assentos em diferentes sessões de eventos.

## Estrutura do Projeto

A estrutura do projeto está organizada da seguinte forma:
```
├── cmd 
│ └── events
│   └── main.go # Ponto de entrada da aplicação
├── internal
│ └── events
│   └── domain # Estrutura de domínio com definição das entidades e repositório da aplicação
|   └── infra # Estrutura de infra
|     └── http # para captura das requisições feitas a este serviço
|     └── repository # para realizar o acesso e persistência no BD desta aplicação
|     └── service # para realizar requisições a APIs externas e tratar as respostas
|   └── usecase # Implementação da regra de negócio dos endpoints da aplicação.
```

## Configuração e Execução

### Pré-requisitos (sem Docker)

1. Realizar download da linguagem Go
- [Go](https://go.dev/)

2. Instalar uma IDE ou um editor de texto como o VSCode caso não possua.
- [VSCode](https://code.visualstudio.com/download)
- [GoLand](https://www.jetbrains.com/go/) *Free-30-day trial*

3. Realizar a instalação da extensão Go (opcional && apenas VSCode)
- Go (Go Team at Google)

4. Na barra de pesquisa digitar **>go** e clicar em **Install/update tools**. Selecionar todos e clicar em *OK*.


### Configuração do Banco de Dados

**PENDING**

### Configuração da Aplicação

1. Clone o repositório
```
git clone https://github.com/leomoritz/go-events-api.git
cd go-events-api
```

2. Executando a aplicação
```
go run cmd/events/main.go
```

## Endpoints da API

### Reserva de Ingressos

- **`POST /checkout`**
  - **Descrição:** Reserva um ingresso.
  - **Body:**
    - `event_id` (string): ID do evento.
    - `spots` (string[]): Assentos.
    - `ticket_type` (string): Tipo do ingresso.
    - `card_hash` (string): Hash do cartão.
    - `email` (string): E-mail do comprador.

- **`GET /events`**
  - **Descrição:** Lista todos os eventos.
  - **Parâmetros:** Nenhum.

- **`GET /events/{eventID}`**
  - **Descrição:** Busca um evento pelo ID.
  - **Parâmetros:**
    - `id` (string): ID do evento.

- **`GET /events/{eventID}/spots`**
  - **Descrição:** Busca os assentos/lugares de um evento pelo ID.
  - **Parâmetros:**
    - `id` (string): ID do evento.


