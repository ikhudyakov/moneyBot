package service

import (
	"database/sql"
	c "moneybot/internal/config"
	m "moneybot/internal/model"
	r "moneybot/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Control interface {
	GetUser(userId int64) (*m.User, error)
}

type Service struct {
	Control
}

func NewService(repos *r.Repository, conf *c.Config, db *sql.DB) *Service {
	return &Service{
		Control: NewControlService(repos.Control, conf, db),
	}
}
