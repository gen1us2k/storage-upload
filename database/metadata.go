package database

type Metadata struct {
	ID       string `db:"id"`
	FileName string `db:"filename"`
	FilePath string `db:"filepath"`
}
