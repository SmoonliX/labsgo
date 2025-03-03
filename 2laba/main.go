package main

import (
	"errors"
	"fmt"
	"math"
)

// Задание 1.1: Функция formatIP для форматирования IP-адреса
func formatIP(ip [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// Задание 1.2: Функция listEven для получения чётных чисел в диапазоне
func listEven(a, b int) ([]int, error) {
	if a > b {
		return nil, errors.New("левая граница диапазона больше правой")
	}

	var evenNumbers []int // Срез для хранения чётных чисел

	for i := a; i <= b; i++ {
		if i%2 == 0 {
			evenNumbers = append(evenNumbers, i) // Добавляем чётное число в срез
		}
	}

	return evenNumbers, nil
}

// Задание 2: Функция countChars для подсчёта символов в строке
func countChars(s string) map[rune]int {
	charCount := make(map[rune]int) // Создаём пустую карту

	for _, char := range s { // Перебираем символы строки
		charCount[char]++ // Увеличиваем счётчик для текущего символа
	}

	return charCount
}

// Задание 3.1: Структура "точка"
type Point struct {
	X float64
	Y float64
}

// Задание 3.2: Структура "отрезок"
type Segment struct {
	Start Point
	End   Point
}

// Задание 3.3: Метод Length для вычисления длины отрезка
func (s Segment) Length() float64 {
	dx := s.End.X - s.Start.X
	dy := s.End.Y - s.Start.Y
	return math.Sqrt(dx*dx + dy*dy) // Формула расстояния между двумя точками
}

// Задание 3.4: Структура "треугольник"
type Triangle struct {
	A, B, C Point
}

// Задание 3.5: Структура "круг"
type Circle struct {
	Center Point
	Radius float64
}

// Задание 3.6: Метод Area для треугольника (формула Герона)
func (t Triangle) Area() float64 {
	a := math.Sqrt(math.Pow(t.B.X-t.A.X, 2) + math.Pow(t.B.Y-t.A.Y, 2))
	b := math.Sqrt(math.Pow(t.C.X-t.B.X, 2) + math.Pow(t.C.Y-t.B.Y, 2))
	c := math.Sqrt(math.Pow(t.A.X-t.C.X, 2) + math.Pow(t.A.Y-t.C.Y, 2))
	p := (a + b + c) / 2
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

// Задание 3.6: Метод Area для круга
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Задание 3.7: Интерфейс "фигура"
type Shape interface {
	Area() float64
}

// Задание 3.8: Функция для вывода площади фигуры
func printArea(s Shape) {
	fmt.Printf("Площадь фигуры: %.2f\n", s.Area())
}

func main() {
	// Задание 1.1: Пример использования функции formatIP
	ip := [4]byte{127, 0, 0, 1}
	fmt.Println("IP-адрес:", formatIP(ip)) // Вывод: 127.0.0.1

	// Задание 1.2: Пример использования функции listEven
	numbers, err := listEven(1, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Чётные числа:", numbers) // Вывод: [2 4 6 8 10]
	}

	// Задание 2: Пример использования функции countChars
	str := "hello"
	charCount := countChars(str)
	fmt.Println("Количество символов:", charCount) // Вывод: map[h:1 e:1 l:2 o:1]

	// Задание 3: Пример использования структур и интерфейсов
	// Создаём отрезок
	segment := Segment{Start: Point{0, 0}, End: Point{7, 5}}
	fmt.Printf("Длина отрезка: %.2f\n", segment.Length()) // Вывод: Длина отрезка: 8.60

	// Создаём треугольник
	triangle := Triangle{A: Point{0, 0}, B: Point{6, 0}, C: Point{0, 8}}
	printArea(triangle) // Вывод: Площадь фигуры: 24.00

	// Создаём круг
	circle := Circle{Center: Point{0, 0}, Radius: 10}
	printArea(circle) // Вывод: Площадь фигуры: 314.16
}
