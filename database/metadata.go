package database

type Metadata struct {
	ID       string `db:"id"`
	Filename string `db:"filename"`
	Path     string `db:"filepath"`
	Size     int64  `db:"size"`
}
