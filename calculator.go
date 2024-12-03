package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите выражение:")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic("Ошибка ввода данных")
	}

	// Удаляем лишние пробелы
	expression := strings.TrimSpace(input)

	// Вычисляем
	result, err := calculate(expression)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func calculate(expression string) (string, error) {
	// Определяем оператор
	var operator string
	if strings.Contains(expression, "+") {
		operator = "+"
	} else if strings.Contains(expression, "-") {
		operator = "-"
	} else if strings.Contains(expression, "*") {
		operator = "*"
	} else if strings.Contains(expression, "/") {
		operator = "/"
	} else {
		return "", fmt.Errorf("Неподдерживаемая операция")
	}

	// Разбиваем выражение на части, если их не 2, то выводим ошибку
	parts := strings.Split(expression, operator)
	if len(parts) != 2 {
		return "", fmt.Errorf("Некорректный формат выражения")
	}

	// Убираем пробелы
	left := strings.TrimSpace(parts[0])
	right := strings.TrimSpace(parts[1])

	// Проверяем длину аргументов, если больше 10, то выводим ошибку
	if len(left) > 10 {
		return "", fmt.Errorf("Левая строка превышает 10 символов")
	}
	if len(right) > 10 {
		return "", fmt.Errorf("Правая строка превышает 10 символов")
	}

	// Расчитываем выражение
	var result string
	switch operator {
	case "+":
		result = left + right

	case "-":
		result = strings.ReplaceAll(left, right, "")

	case "*":
		num, err := strconv.Atoi(right)
		if err != nil || num < 1 || num > 10 {
			return "", fmt.Errorf("Второй аргумент должен быть числом от 1 до 10")
		}
		result = strings.Repeat(left, num)

	case "/":
		num, err := strconv.Atoi(right)
		if err != nil || num < 1 || num > 10 {
			return "", fmt.Errorf("Второй аргумент должен быть числом от 1 до 10")
		}
		newLen := len(left) / num
		result = left[:newLen]

	default:
		return "", fmt.Errorf("Неподдерживаемая операция")
	}

	// Обрезаем строку, если она превышает 40 символов
	if len(result) > 40 {
		result = result[:40] + "..."
	}

	return result, nil
}
