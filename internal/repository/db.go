package repository

import (
	"database/sql"
	"fmt"
	c "moneybot/internal/config"

	_ "github.com/lib/pq"
)

// Подключение к базе данных
func Connect(conf *c.Config) (*sql.DB, error) {
	var err error
	var conn string
	var db *sql.DB

	conn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.DBname)
	if db, err = sql.Open("postgres", conn); err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
