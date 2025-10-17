package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
	"github.com/piyushk8/StudentAPI/internal/config"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	// Create students table if not exists
	query := `
	CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		age INTEGER NOT NULL
	);
	`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &Sqlite{Db: db}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO STUDENTS (name, email, age) VALUES (?,?,?)")

	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}
