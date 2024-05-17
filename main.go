package main

import (
	"context"
	"log"
	"os"

	"cdr.dev/slog"
	"cdr.dev/slog/sloggers/slogjson"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Nikolay200669/go-auto-invoice/interfaces/handlers"
	"github.com/jinzhu/gorm"

	"github.com/Nikolay200669/go-auto-invoice/infrastructure/persistence"
)

var db *gorm.DB

func initDB() *gorm.DB {
	if db == nil {
		var err error
		db, err = gorm.Open("sqlite3", "database.db")
		if err != nil {
			log.Fatal(err)
		}
	}
	return db
}

func main() {

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Создание логгера
	logger := slog.Make(slogjson.Sink(logFile))

	// Использование логгера
	logger.Info(context.Background(), "Application started")

	// Инициализация базы данных
	db := initDB()
	repo, err := persistence.NewGormRepository(db)
	if err != nil {
		logger.Error(context.Background(),
			"Failed to create Gorm repository", "error", err)
	}

	// Создание HTTP-сервера и обработчиков
	server := handlers.NewHTTPServer(repo, &logger)
	server.Start()
}
