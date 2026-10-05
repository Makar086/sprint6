package main

import (
	"log"
	"os"

	"https://github.com/Makar086/sprint6.git/internal/server"
	
)

func main() {
	
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags|log.Lshortfile)

	logger.Println("Инициализация сервера...")

	
	srv := server.New() 

	logger.Println("Запуск сервера на порту :8080...")

	
	if err := srv.Start(":8080"); err != nil {
		
		logger.Fatalf("Критическая ошибка при запуске сервера: %v", err)
	}
}
}
