package modules

import (
	"context"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Customer представляет структуру данных для покупателя
type Customer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	BirthDate    string `json:"birthDate"`
	Email        string `json:"email"`
	Password     string `json:"password"` // Исходный пароль (используется только при регистрации)
	PasswordHash string `json:"-"`        // Хэшированный пароль (не экспортируется в JSON)
}

// CreateCustomer создаёт нового пользователя в базе данных
func (db *DB) CreateCustomer(customer Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Хэшируем исходный пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Ошибка при хэшировании пароля: %v", err)
		return err
	}

	log.Printf("Регистрируем пользователя: %s", customer.Email)

	// Сохраняем хэшированный пароль в базу данных
	_, err = db.pool.Exec(ctx,
		`INSERT INTO customers (name, surname, birth_date, email, password_hash) 
         VALUES ($1, $2, $3, LOWER($4), $5)`, // LOWER(email) гарантирует, что email сохраняется в нижнем регистре
		customer.Name, customer.Surname, customer.BirthDate, customer.Email, string(hashedPassword),
	)
	if err != nil {
		log.Printf("Ошибка при вставке пользователя в базу данных: %v", err)
		return err
	}

	log.Printf("Пользователь успешно зарегистрирован: %s", customer.Email)
	return nil
}

// GetCustomerByEmail возвращает пользователя по email
func (db *DB) GetCustomerByEmail(email string) (Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var customer Customer

	log.Printf("Поиск пользователя с email: %s", email)

	// Выполняем запрос к базе данных
	err := db.pool.QueryRow(ctx,
		`SELECT id, name, surname, birth_date, email, password_hash 
         FROM customers WHERE LOWER(email) = LOWER($1)`, email, // Ищем email без учета регистра
	).Scan(&customer.ID, &customer.Name, &customer.Surname, &customer.BirthDate, &customer.Email, &customer.PasswordHash)

	log.Print(customer)
	if err != nil {
		log.Printf("Ошибка при поиске пользователя: %v", err)
		return customer, err
	}

	log.Printf("Пользователь найден: %v", customer)
	return customer, nil
}

// GetCustomerByID возвращает пользователя по ID
func (db *DB) GetCustomerByID(id int) (Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var customer Customer

	log.Printf("Поиск пользователя по ID: %d", id)

	err := db.pool.QueryRow(ctx,
		`SELECT id, name, surname, birth_date, email 
         FROM customers WHERE id = $1`, id,
	).Scan(&customer.ID, &customer.Name, &customer.Surname, &customer.BirthDate, &customer.Email)

	if err != nil {
		log.Printf("Ошибка при получении пользователя по ID: %v", err)
		return customer, err
	}

	log.Printf("Пользователь найден: %v", customer)
	return customer, nil
}
