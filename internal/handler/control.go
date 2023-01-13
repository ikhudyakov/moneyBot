package handler

import (
	"log"
	m "moneybot/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) Update() {
	var err error

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := h.bot.GetUpdatesChan(u)

	for update := range updates {
		var user *m.User
		user, err = h.service.GetUser(update.Message.From.ID)
		if err != nil {
			log.Println(err)
		}

		if user != nil {
			log.Println(user)
		}

		// Ответ по умолчанию
		reply := "Добро пожаловать в MyMoney Bot, " + update.Message.From.FirstName + "\nДля регистрации в MyMoney Bot нужно ввести MyMoney id в личном кабинете\nВсе команды - /help"
		if update.Message == nil {
			continue
		}

		// Логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// 	// Обработка комманд
		// 	switch update.Message.Command() {
		// 	case "help":
		// 		reply = "Добавить расход:\nРасход-id категории-сумма-описание" +
		// 			"\n\nДобавить доход:\nДоход-id категории-сумма-описание" +
		// 			"\n\nПосмотреть категории расходов и доходов - /category" +
		// 			"\n\nУзнать баланс - /balance" +
		// 			"\n\nПолучить MyMoney id - /getID"
		// 	case "category":
		// 		reply = category()
		// 	case "getID":
		// 		reply = strconv.Itoa(update.Message.From.ID)
		// 	case "balance":
		// 		reply = "В разработке"
		// 	}

		// 	// Добавление расхода
		// 	if strings.HasPrefix(update.Message.Text, "Расход") && userId != 0 {
		// 		in := strings.Split(update.Message.Text, "-")
		// 		if add(in, "COSTS", userId) > 0 {
		// 			reply = "Расход добавлен"
		// 		} else {
		// 			reply = "Ошибка добавления записи"
		// 		}
		// 	}

		// 	// Добавление дохода
		// 	if strings.HasPrefix(update.Message.Text, "Доход") && userId != 0 {
		// 		in := strings.Split(update.Message.Text, "-")
		// 		if add(in, "RECEIPTS", userId) > 0 {
		// 			reply = "Доход добавлен"
		// 		} else {
		// 			reply = "Ошибка добавления записи"
		// 		}
		// 	}

		// Создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// Отправляем
		h.bot.Send(msg)
	}

}
