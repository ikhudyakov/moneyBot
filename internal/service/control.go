package service

import (
	"database/sql"
	c "moneybot/internal/config"
	m "moneybot/internal/model"
	r "moneybot/internal/repository"
)

type ControlService struct {
	repo r.Control
	conf *c.Config
	db   *sql.DB
}

func NewControlService(repo r.Control, conf *c.Config, db *sql.DB) *ControlService {
	return &ControlService{
		repo: repo,
		conf: conf,
		db:   db,
	}
}

func (c *ControlService) GetUser(myMoneyId int64) (*m.User, error) {
	var user *m.User
	var err error

	user, err = c.repo.GetUser(myMoneyId)

	return user, err
}
