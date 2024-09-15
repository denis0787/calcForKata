package main

import (
   "fmt"
   "strconv"
   "strings"
)

// Маппинг римских чисел на арабские
var romanToArabic = map[string]int{
   "I":    1,
   "II":   2,
   "III":  3,
   "IV":   4,
   "V":    5,
   "VI":   6,
   "VII":  7,
   "VIII": 8,
   "IX":   9,
   "X":    10,
}

// Массив арабских чисел для преобразования в римские
var arabicToRoman = []string{
   "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
      "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
      "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
      "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
      "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L",
      "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
      "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
      "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
      "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
      "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C",
}

// Функция для проверки, является ли строка римским числом
func isRoman(input string) bool {
   _, exists := romanToArabic[input]
   return exists
}

// Функция для преобразования римского числа в арабское
func toArabic(input string) (int, bool) {
   value, exists := romanToArabic[input]
   return value, exists
}

// Функция для преобразования арабского числа в римское
func toRoman(num int) (string, bool) {
   if num <= 0 || num >= len(arabicToRoman) {
      return "", false
   }
   return arabicToRoman[num], true
}

// Функция для выполнения арифметической операции
func calculate(a, b int, operator string) int {
   switch operator {
   case "+":
      return a + b
   case "-":
      return a - b
   case "*":
      return a * b
   case "/":
      return a / b
   default:
      panic("неизвестный оператор")
   }
}

// Основная функция
func main() {
   var expression string
   fmt.Print("Введите выражение: ")
   fmt.Scanln(&expression)

   // Удаление пробелов
   expression = strings.ReplaceAll(expression, " ", "")

   // Определение оператора
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
      panic("операция не распознана")
   }

   // Разделение строки на числа
   parts := strings.Split(expression, operator)
   if len(parts) != 2 {
      panic("неверный формат выражения")
   }

   // Определение типа чисел (римские или арабские)
   isRomanNumeral := isRoman(parts[0])
   if isRomanNumeral != isRoman(parts[1]) {
      panic("разные форматы чисел")
   }

   var a, b int
   var err error
   if isRomanNumeral {
      a, _ = toArabic(parts[0])
      b, _ = toArabic(parts[1])
   } else {
      a, err = strconv.Atoi(parts[0])
      if err != nil || a < 1 || a > 10 {
         panic("введено некорректное число")
      }
      b, err = strconv.Atoi(parts[1])
      if err != nil || b < 1 || b > 10 {
         panic("введено некорректное число")
      }
   }

   // Выполнение операции
   result := calculate(a, b, operator)

   // Обработка вывода результата
   if isRomanNumeral {
      if result < 1 {
         panic("результат меньше единицы не может быть представлен римскими числами")
      }
      romanResult, ok := toRoman(result)
      if !ok {
         panic("ошибка преобразования в римские числа")
      }
      fmt.Println("Результат:", romanResult)
   } else {
      fmt.Println("Результат:", result)
   }
}
