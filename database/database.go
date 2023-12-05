package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type FileStorage interface {
	SaveMetadata(context.Context, *Metadata) (*Metadata, error)
	GetFiles(context.Context) ([]Metadata, error)
	Close() error
}

var _ FileStorage = &Postgres{}

type Postgres struct {
	conn *sqlx.DB
}

// NewPostgres creates a new Postgres interface.
func NewPostgres(dsn string) (*Postgres, error) {
	db, err := sqlx.Connect("postgres", dsn)
	return &Postgres{conn: db}, err
}

// Migrate runs migration against the created connection to a Postgres database.
func (p *Postgres) Migrate() error {
	schema := `CREATE TABLE IF NOT EXISTS files (
		id uuid DEFAULT uuid_generate_v4(),
		filename VARCHAR(255) NOT NULL,
		filepath VARCHAR(512) NOT NULL,
		PRIMARY KEY(id)
		);`

	if _, err := p.conn.Exec(schema); err != nil {
		return err
	}
	return nil
}

// SaveMetadata saves metadata into a database.
func (p *Postgres) SaveMetadata(ctx context.Context, m *Metadata) (*Metadata, error) {
	return nil, nil
}

// GetFiles returns a list of files.
func (p *Postgres) GetFiles(ctx context.Context) ([]Metadata, error) {
	return []Metadata{}, nil
}

// Close closes connection.
func (p *Postgres) Close() error {
	return p.conn.Close()
}
