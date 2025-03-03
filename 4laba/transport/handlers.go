package transport

import (
    "net/http"
    "log"
    "golang.org/x/crypto/bcrypt"
    "github.com/gin-gonic/gin"
    "main/modules"
)

type BaseHandler struct {
    db *modules.DB
}

func NewBaseHandler(db *modules.DB) *BaseHandler {
    return &BaseHandler{db: db}
}

func (h *BaseHandler) Register(c *gin.Context) {
    var customer modules.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
        return
    }

    if err := h.db.CreateCustomer(customer); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось зарегистрировать пользователя"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Пользователь успешно зарегистрирован"})
}

func (h *BaseHandler) Login(c *gin.Context) {
    var loginRequest struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    // Парсим JSON
    if err := c.ShouldBindJSON(&loginRequest); err != nil {
        log.Printf("Ошибка при парсинге JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
        return
    }

    // Логируем попытку входа
    log.Printf("Попытка входа для email: %s", loginRequest.Email)

    // Получаем пользователя по email
    customer, err := h.db.GetCustomerByEmail(loginRequest.Email)
    if err != nil {
        log.Printf("Ошибка при поиске пользователя: %v", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
        return
    }

    // Логируем найденного пользователя
    log.Printf("Пользователь найден: %v", customer)

    // Сравниваем пароль
    if err := bcrypt.CompareHashAndPassword([]byte(customer.PasswordHash), []byte(loginRequest.Password)); err != nil {
        log.Printf("Ошибка при проверке пароля: %v", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный пароль"})
        return
    }

    // Генерируем токен
    token, err := modules.GenerateJWT(customer.ID)
    if err != nil {
        log.Printf("Ошибка при генерации токена: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен"})
        return
    }

    // Логируем успешный вход
    log.Printf("Токен успешно сгенерирован для пользователя: %s", loginRequest.Email)
    c.JSON(http.StatusOK, gin.H{"token": token})

    log.Printf("Хэш из БД: %s", customer.PasswordHash)
    log.Printf("Введённый пароль: %s", loginRequest.Password)
}

func (h *BaseHandler) GetCustomer(c *gin.Context) {
    customerID := c.GetInt("customerID")

    customer, err := h.db.GetCustomerByID(customerID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
        return
    }

    c.JSON(http.StatusOK, customer)
}