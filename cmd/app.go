package main

import (
	"log"
	c "moneybot/internal/config"
	h "moneybot/internal/handler"
	"os"
)

type App struct {
}

func (s *App) Run(handler h.Handler, conf c.Config) error {
	if err := handler.Init(conf.TelegramBotToken); err != nil {
		return err
	}

	log.Println("приложение запущено")

	handler.Update()

	return nil
}

func (s *App) Shutdown() {
	os.Exit(1)
}
