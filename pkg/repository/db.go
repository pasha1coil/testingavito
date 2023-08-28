package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	Users     = "users"
	slugs     = "slugs"
	UsersSlug = "UsersSlug"
	History   = "History"
)

type Config struct {
	Host   string
	Port   string
	Uname  string
	Pass   string
	NameDB string
	SSL    string
}

func NewDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Uname, cfg.NameDB, cfg.Pass, cfg.SSL))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
