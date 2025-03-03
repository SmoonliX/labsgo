package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Структура для парсинга JSON из запроса (Задание 3)
type RequestBody struct {
	Text string `json:"text"`
}

// Обработчик для GET-запроса с query-параметрами (Задание 1) localhost:8080/hello?name=Andrew&age=22
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем query-параметры из запроса
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	// Формируем ответ
	response := fmt.Sprintf("Меня зовут %s, мне %s лет", name, age)

	// Отправляем ответ клиенту
	w.Write([]byte(response))
}

// Обработчик для сложения (Задание 2)http://localhost:8080/add?a=10&b=5
//http://localhost:8080/sub?a=10&b=5
//http://localhost:8080/mul?a=10&b=5
//http://localhost:8080/div?a=10&b=5
func addHandler(w http.ResponseWriter, r *http.Request) {
	a, errA := strconv.Atoi(r.URL.Query().Get("a"))
	b, errB := strconv.Atoi(r.URL.Query().Get("b"))

	if errA != nil || errB != nil {
		http.Error(w, "Неверные параметры", http.StatusBadRequest)
		return
	}

	result := a + b
	w.Write([]byte(fmt.Sprintf("Результат: %d", result)))
}

// Обработчик для вычитания (Задание 2)
func subHandler(w http.ResponseWriter, r *http.Request) {
	a, errA := strconv.Atoi(r.URL.Query().Get("a"))
	b, errB := strconv.Atoi(r.URL.Query().Get("b"))

	if errA != nil || errB != nil {
		http.Error(w, "Неверные параметры", http.StatusBadRequest)
		return
	}

	result := a - b
	w.Write([]byte(fmt.Sprintf("Результат: %d", result)))
}

// Обработчик для умножения (Задание 2)
func mulHandler(w http.ResponseWriter, r *http.Request) {
	a, errA := strconv.Atoi(r.URL.Query().Get("a"))
	b, errB := strconv.Atoi(r.URL.Query().Get("b"))

	if errA != nil || errB != nil {
		http.Error(w, "Неверные параметры", http.StatusBadRequest)
		return
	}

	result := a * b
	w.Write([]byte(fmt.Sprintf("Результат: %d", result)))
}

// Обработчик для деления (Задание 2)
func divHandler(w http.ResponseWriter, r *http.Request) {
	a, errA := strconv.Atoi(r.URL.Query().Get("a"))
	b, errB := strconv.Atoi(r.URL.Query().Get("b"))

	if errA != nil || errB != nil {
		http.Error(w, "Неверные параметры", http.StatusBadRequest)
		return
	}

	if b == 0 {
		http.Error(w, "Деление на ноль невозможно", http.StatusBadRequest)
		return
	}

	result := a / b
	w.Write([]byte(fmt.Sprintf("Результат: %d", result)))
}

// Обработчик для POST-запроса (Задание 3)http://localhost:8080/count
func countCharsHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что метод запроса — POST
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Декодируем JSON из тела запроса
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	// Подсчитываем количество вхождений каждого символа
	charCount := make(map[rune]int)
	for _, char := range requestBody.Text {
		charCount[char]++
	}

	// Отправляем ответ в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(charCount)
}

func main() {
	// Регистрируем обработчики для всех заданий
	http.HandleFunc("/hello", helloHandler) // Задание 1
	http.HandleFunc("/add", addHandler)    // Задание 2
	http.HandleFunc("/sub", subHandler)    // Задание 2
	http.HandleFunc("/mul", mulHandler)    // Задание 2
	http.HandleFunc("/div", divHandler)    // Задание 2
	http.HandleFunc("/count", countCharsHandler) // Задание 3

	// Запускаем сервер на порту 8080
	fmt.Println("Сервер -> http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}