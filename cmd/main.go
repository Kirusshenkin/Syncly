package main

import (
	config "github.com/Kirusshenkin/syncly.git/configs"
	"github.com/Kirusshenkin/syncly.git/internal/infra/db"
	http2 "github.com/Kirusshenkin/syncly.git/internal/interfaces/http"
	"log"
	"net/http"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Настраиваем подключение к базе данных
	postgresPool := db.NewPostgresPool(cfg.Database.DSN)
	defer postgresPool.Close()

	// Настраиваем HTTP маршруты (пустые для примера)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Chat Backend!"))
	})

	// Запускаем HTTP сервер
	server := http2.NewServer(mux, cfg.App.Port)
	server.Start()

}
