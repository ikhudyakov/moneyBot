package repository

import (
	"database/sql"
	m "moneybot/internal/model"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Repository struct {
	Control
}

type Control interface {
	GetUser(userId int64) (*m.User, error)
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Control: NewControlPostgres(db),
	}
}
