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
	schema := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    CREATE TABLE IF NOT EXISTS files (
		id uuid DEFAULT uuid_generate_v4(),
		filename VARCHAR(255) NOT NULL,
		filepath VARCHAR(512) NOT NULL,
        size INTEGER NOT NULL,
		PRIMARY KEY(id)
    );`

	if _, err := p.conn.Exec(schema); err != nil {
		return err
	}
	return nil
}

// SaveMetadata saves metadata into a database.
func (p *Postgres) SaveMetadata(ctx context.Context, m *Metadata) (*Metadata, error) {
	created := Metadata{}
	err := p.conn.Get(&created, "INSERT INTO files (filename, filepath, size) VALUES($1, $2, $3) RETURNING id, filename, filepath, size", m.Filename, m.Path, m.Size)

	return &created, err
}

// GetFiles returns a list of files.
func (p *Postgres) GetFiles(ctx context.Context) ([]Metadata, error) {
	files := []Metadata{}
	err := p.conn.Select(&files, "SELECT * FROM files")
	return files, err
}

// Close closes connection.
func (p *Postgres) Close() error {
	return p.conn.Close()
}
