package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://root:root@db/db?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	err := d.db.Close()
	if err != nil {
		log.Fatal("closing db connection error")
		return
	}
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
