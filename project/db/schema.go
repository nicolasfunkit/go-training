package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS tickets (
	ticket_id UUID PRIMARY KEY,
	price_amount NUMERIC(10, 2) NOT NULL,
	price_currency CHAR(3) NOT NULL,
	customer_email VARCHAR(255) NOT NULL
);
DELETE FROM tickets;
`

func InitializeDatabaseSchema(db *sqlx.DB) error {
	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("could not initialize database schema: %w", err)
	}

	return nil

}
