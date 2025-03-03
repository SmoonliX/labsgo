package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "main/config"
    "main/modules"
    "main/transport"
)

func main() {
    // Получаем конфигурацию подключения к базе данных
    cfg := config.GetConfig()

    // Создаём подключение к базе данных
    db, err := modules.NewDB(cfg.DBConnString)
    if err != nil {
        log.Fatalf("Не удалось подключиться к базе данных: %v", err)
    }
    defer db.Close()

    // Создаём обработчики
    handler := transport.NewBaseHandler(db)

    // Настраиваем роутер
    r := gin.Default()

    // Регистрация и авторизация
    r.POST("/register", handler.Register)
    r.POST("/login", handler.Login)

    // Защищённые маршруты
    authGroup := r.Group("/")
    authGroup.Use(transport.AuthRequired)
    {
        authGroup.GET("/customer/:id", handler.GetCustomer)
    }

    // Запускаем сервер
    r.Run(":8080")
}