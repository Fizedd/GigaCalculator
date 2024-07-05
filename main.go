package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Карта для преобразования римских цифр в целые числа
var romanToIntMap = map[string]int{
	"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10,
	"XL": 40, "L": 50, "XC": 90, "C": 100,
	
}

// Карта для преобразования целых чисел в римские цифры
var intToRomanMap = []struct {
	value  int
	symbol string
}{
	
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func main() {
	// Создание нового считывателя для ввода с консоли
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите ваше вычисление (например, 3+4 или IV+VI):")
	// Чтение строки, введенной пользователем
	input, _ := reader.ReadString('\n')
	// Удаление лишних пробелов в начале и конце строки
	input = strings.TrimSpace(input)

	// Разделение входной строки на токены (числа и оператор)
	tokens := splitInput(input)
	if len(tokens) != 3 {
		panic("Неправильный формат ввода")
	}

	// Присваивание значений операндов и оператора
	a, b := tokens[0], tokens[2]
	op := tokens[1]

	isRoman := false
	// Преобразование операндов в целые числа
	num1, err1 := strconv.Atoi(a)
	num2, err2 := strconv.Atoi(b)

	if err1 != nil && err2 != nil {
		// Проверка, являются ли оба операнда римскими цифрами
		num1 = romanToInt(a)
		num2 = romanToInt(b)
		isRoman = true
	} else if err1 != nil || err2 != nil {
		panic("Смешивание римских и арабских цифр не допускается")
	}

	// Проверка, находятся ли числа в диапазоне от 1 до 10 включительно
	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		panic("Числа должны быть от 1 до 10")
	}

	var result int

	// Выполнение арифметической операции
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			panic("Деление на ноль не допускается")
		}
		result = num1 / num2
	default:
		panic("Неправильная операция")
	}

	// Вывод результата в римских или арабских цифрах
	if isRoman {
		if result <= 0 {
			panic("Римские цифры должны быть положительными")
		}
		fmt.Println(intToRoman(result))
	} else {
		fmt.Println(result)
	}
}

// Функция для разделения входной строки на токены
func splitInput(input string) []string {
	tokens := []string{}
	current := strings.Builder{}
	for _, ch := range input {
		if unicode.IsSpace(ch) {
			continue
		}
		if unicode.IsDigit(ch) || unicode.IsLetter(ch) {
			current.WriteRune(ch)
		} else {
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			tokens = append(tokens, string(ch))
		}
	}
	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}
	return tokens
}

// Функция для преобразования римских цифр в целые числа
func romanToInt(roman string) int {
	total := 0
	i := 0
	for i < len(roman) {
		if i+1 < len(roman) {
			if val, exists := romanToIntMap[roman[i:i+2]]; exists {
				total += val
				i += 2
				continue
			}
		}
		if val, exists := romanToIntMap[roman[i:i+1]]; exists {
			total += val
			i++
		} else {
			panic("Неправильная римская цифра")
		}
	}
	return total
}

// Функция для преобразования целых чисел в римские цифры
func intToRoman(num int) string {
	var result strings.Builder
	for _, entry := range intToRomanMap {
		for num >= entry.value {
			result.WriteString(entry.symbol)
			num -= entry.value
		}
	}
	return result.String()
}
