package repository

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leomoritz/go-events-api/internal/events/domain"
)

/*OBS:
* Este projeto utilizou SQL puro para realizar o acesso ao banco de dados.
* Contudo, existem ORMs que podem ser utilizados para este fim. Segue abaixo exemplo de ORM:
* - https://gorm.io/docs/
 */

type MysqlEventRepository struct {
	db *sql.DB
}

func NewMysqlEventRepository(db *sql.DB) (domain.EventRepository, error) {
	return &MysqlEventRepository{db: db}, nil
}

// CreateEvent implements domain.EventRepository.
func (r *MysqlEventRepository) CreateEvent(event *domain.Event) error {
	query := `
		INSERT INTO events (id, name, location, organization, rating, date, image_url, capacity, price, partner_id)
		VALUES (?,?,?,?,?,?,?,?,?,?)
	`

	// Executa a query no banco de dados e retorna o número de linhas afetadas, caso contrário, retorna um erro
	_, err := r.db.Exec(query, event.ID, event.Name, event.Location, event.Organization, event.Rating,
		event.Date, event.ImageURL, event.Capacity, event.Price, event.PartnerID)

	return err
}

func (r *MysqlEventRepository) CreateSpot(spot *domain.Spot) error {
	query := `
		INSERT INTO spots (id, event_id, name, status, ticket_id)
		VALUES (?, ?, ?, ?, ?)
	`

	// Executa a query no banco de dados e retorna o número de linhas afetadas, caso contrário, retorna um erro
	_, err := r.db.Exec(query, spot.ID, spot.EventID, spot.Name, spot.Status, spot.TicketID)

	return err
}

func (r *MysqlEventRepository) ReserveSpot(spotID, ticketID string) error {
	query := `
        UPDATE spots
        SET status =?, ticket_id =?
        WHERE id =?
    `

	// Executa a query no banco de dados e retorna o número de linhas afetadas, caso contrário, retorna um erro
	_, err := r.db.Exec(query, domain.SpotStatusSold, ticketID, spotID)

	return err
}

func (r *MysqlEventRepository) CreateTicket(ticket *domain.Ticket) error {
	query := `
        INSERT INTO tickets (id, event_id, spot_id, ticket_type, price)
        VALUES (?, ?, ?, ?, ?)
    `

	// Executa a query no banco de dados e retorna o número de linhas afetadas, caso contrário, retorna um erro
	_, err := r.db.Exec(query, ticket.ID, ticket.EventID, ticket.Spot.ID, ticket.TicketType, ticket.Price)

	return err
}

func (r *MysqlEventRepository) FindEventById(eventID string) (*domain.Event, error) {
	query := `
		SELECT id, name, location, organization, rating, date, image_url, capacity, price, partner_id
		FROM events
		WHERE id =?
	`

	rows, err := r.db.Query(query, eventID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var event domain.Event

	// o & significa que a cada loop, o valor do atributo de event será populado com o valor que veio do banco.
	err = rows.Scan(
		&event.ID,
		&event.Name,
		&event.Location,
		&event.Organization,
		&event.Rating,
		&event.Date,
		&event.ImageURL,
		&event.Capacity,
		&event.Price,
		&event.PartnerID,
	)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *MysqlEventRepository) FindSpotsByEventID(eventID string) ([]*domain.Spot, error) {
	query := `
		SELECT id, event_id, name, status, ticket_id
		FROM spots
		WHERE event_id =?    
	`

	rows, err := r.db.Query(query, eventID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var spots []*domain.Spot
	for rows.Next() {
		var spot domain.Spot

		err = rows.Scan(
			&spot.ID,
			&spot.EventID,
			&spot.Name,
			&spot.Status,
			&spot.TicketID,
		)

		if err != nil {
			return nil, err
		}

		spots = append(spots, &spot)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return spots, nil
}

func (r *MysqlEventRepository) FindSpotByName(eventID, spotName string) (*domain.Spot, error) {
	query := `
	    SELECT s.id, s.event_id, s.name, s.status, s.ticket_id,
			   t.id, t.event_id, t.spot_id, t.ticket_type, t.price
		FROM spots s
		LEFT JOIN tickets t ON s.id = t.spot_id
		WHERE s.event_id = ? AND s.name =?
	`

	row, err := r.db.Query(query, spotName)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	var spot domain.Spot
	var ticket domain.Ticket
	var ticketID, ticketEventID, ticketSpotID, ticketType sql.NullString
	var ticketPrice sql.NullFloat64

	err = row.Scan(
		&spot.ID,
		&spot.EventID,
		&spot.Name,
		&spot.Status,
		&spot.TicketID,
		&ticketID,
		&ticketEventID,
		&ticketSpotID,
		&ticketType,
		&ticketPrice,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrSpotNotFound
		}
	}

	// Se for válido, cria um ticket e associa o ID ao spot
	if ticketID.Valid {
		ticket.ID = ticketID.String
		ticket.EventID = ticketEventID.String
		ticket.Spot = &spot
		ticket.TicketType = domain.TicketType(ticketType.String)
		ticket.Price = ticketPrice.Float64
		spot.TicketID = ticket.ID
	}

	return &spot, nil
}

func (r *MysqlEventRepository) ListEvents() ([]domain.Event, error) {
	query := `SELECT id,
	name,
	location,
	organization,
	rating,
	date,
	image_url,
	capacity,
	price,
	partner_id
	FROM events`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []domain.Event

	for rows.Next() {
		var event domain.Event

		err = rows.Scan(
			&event.ID,
			&event.Name,
			&event.Location,
			&event.Organization,
			&event.Rating,
			&event.Date,
			&event.ImageURL,
			&event.Capacity,
			&event.Price,
			&event.PartnerID,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
