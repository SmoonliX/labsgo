package main

import (
	"errors"
	"fmt"
)

// Функция hello принимает строку name и возвращает строку с приветствием.
func hello(name string) string {
	return fmt.Sprintf("Привет, %s!", name)
}

// Функция printEven принимает два целых числа a и b (границы диапазона) и выводит все чётные числа в этом диапазоне (включительно).
// Если левая граница a больше правой b, функция возвращает ошибку.
func printEven(a, b int64) error {
	if a > b {
		return errors.New("левая граница диапазона больше правой")
	}

	// Перебираем все числа от a до b и выводим чётные.
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	return nil
}

// Функция apply принимает два числа a и b типа float64 и строку operator, которая определяет арифметическую операцию.
// Функция возвращает результат операции и ошибку, если операция не поддерживается или происходит деление на ноль.
func apply(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil // Сложение
	case "-":
		return a - b, nil // Вычитание
	case "*":
		return a * b, nil // Умножение
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль") // Проверка деления на ноль
		}
		return a / b, nil // Деление
	default:
		return 0, errors.New("действие не поддерживается") // Неподдерживаемая операция
	}
}

func main() {
	// Тестирование функции hello
	fmt.Println(hello("Андрей")) // Вывод: Привет, Андрей!

	// Тестирование функции printEven
	// Вывод чётных чисел в диапазоне от 1 до 10
	err := printEven(1, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	// Тестирование функции printEven с ошибкой (левая граница больше правой)
	err = printEven(10, 1)
	if err != nil {
		fmt.Println("Ошибка:", err) // Вывод: Ошибка: левая граница диапазона больше правой
	}

	// Тестирование функции apply с операцией сложения
	result, err := apply(6, 5, "+")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}

	// Тестирование функции apply с делением на ноль
	result, err = apply(10, 0, "/")
	if err != nil {
		fmt.Println("Ошибка:", err) // Вывод: Ошибка: деление на ноль
	} else {
		fmt.Println("Результат:", result)
	}
}
