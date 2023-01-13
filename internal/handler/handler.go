package handler

import (
	"log"
	s "moneybot/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	service *s.Service
	bot     *tgbotapi.BotAPI
}

func NewHandler(service *s.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Init(telegramBotToken string) error {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		return err
	}
	h.bot = bot

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return nil
}
