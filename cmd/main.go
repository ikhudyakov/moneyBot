package main

import (
	"database/sql"
	"log"
	c "moneybot/internal/config"
	h "moneybot/internal/handler"
	r "moneybot/internal/repository"
	s "moneybot/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var service *s.Service
	var db *sql.DB
	var err error
	var conf *c.Config
	var path string
	var defaultPath string = "./configs/config.yaml"

	path = defaultPath

	// flag.StringVar(&path, "config", "./configs/config.yaml", "example -config ./configs/config.yaml")
	// migrationup := flag.Bool("migrationup", false, "use migrationup to perform migrationup")
	// migrationdown := flag.Bool("migrationdown", false, "use migrationdown to perform migrationdown")

	// flag.Parse()

	conf, err = c.GetConfig(path)
	if err != nil {
		log.Printf("%s, use default config '%s'", err, defaultPath)
		conf, err = c.GetConfig(defaultPath)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// if err = repository.Migration(conf, *migrationup, *migrationdown); err != nil {
	// 	log.Println(err)
	// }

	if db, err = r.Connect(conf); err != nil {
		log.Println(err)
		return
	}

	repos := r.NewRepository(db)
	service = s.NewService(repos, conf, db)
	handler := h.NewHandler(service)
	app := new(App)

	go func() {
		if err = app.Run(*handler, *conf); err != nil {
			log.Fatalf("ошибка при запуске приложения: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("приложение останавливается")

	app.Shutdown()

	if err := db.Close(); err != nil {
		log.Fatalf("произошла ошибка при закрытии соединения с БД: %s", err.Error())
	}

}
