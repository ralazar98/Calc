package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	Solution(GetValOper())

}

// Solution выбирает решение задачи взависимости от введенных данных
func Solution(operator, firstValue, secondValue string) {
	firstNumberInt, _ := strconv.Atoi(firstValue)
	secondNumberInt, _ := strconv.Atoi(secondValue)

	switch {
	case firstNumberInt == 0 && secondNumberInt == 0:
		firstNumberInt, secondNumberInt = RomanToArabic(firstValue, secondValue)
		resultArab := Calculation(operator, firstNumberInt, secondNumberInt)
		if resultArab < 1 {
			fmt.Println("Выдача паники,так как в римской системе нет отрицательных чисел и ответ " +
				"не может быть равен 0. ")
			os.Exit(0)
		}
		fmt.Println(ArabicToRoman(resultArab))
	case firstNumberInt == 0 && secondNumberInt != 0 || firstNumberInt != 0 && secondNumberInt == 0:
		fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
	case firstNumberInt < 1 || firstNumberInt > 10 || secondNumberInt < 1 || secondNumberInt > 10:
		fmt.Println("Выдача паники, как минимум 1 число вне диапазона ")
	default:
		fmt.Println(Calculation(operator, firstNumberInt, secondNumberInt))
	}

}

// RomanToArabic переводит римские числа в арабские для дальнейшего вычисления результата
// Работает только с числами от 1-10.
func RomanToArabic(firstRomanValue, secondRomanValue string) (int, int) {
	var firstValue, secondValue int
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

	firstValue = romanToArabic[firstRomanValue]
	secondValue = romanToArabic[secondRomanValue]
	if firstValue == 0 || secondValue == 0 {
		fmt.Printf("Выдача паники, операнды должны быть целыми числами от 1 до 10 в римской или арабской " +
			"системе исчисления")
		os.Exit(0)
	}
	return firstValue, secondValue
}

// ArabicToRoman возвращает полученный результат в римскую систему исчисления
func ArabicToRoman(result int) string {
	var arabicToRoman = map[int]string{
		1:   "I",
		4:   "IV",
		5:   "V",
		9:   "IX",
		10:  "X",
		40:  "XL",
		50:  "L",
		90:  "XC",
		100: "C",
	}
	var romanResult string
	var i int = 8
	numbers := [9]int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	for result > 0 {
		for numbers[i] > result {
			i--
		}
		romanResult += arabicToRoman[numbers[i]]
		result -= numbers[i]
	}
	return romanResult
}

// Calculation выполняет операции сложения,вычитания,деления,уножения взависимости от полученных данных,
// возвращает результат вычислений
func Calculation(operator string, firstValue, secondValue int) int {
	var result int
	switch operator {
	case "+":
		result = firstValue + secondValue
	case "-":
		result = firstValue - secondValue
	case "/":
		result = firstValue / secondValue
	case "*":
		result = firstValue * secondValue
	}
	return result
}

// GetValOper считывает вводимую в терминал строку и возвращает операнды и оператор, если это возможно
func GetValOper() (string, string, string) {

	var operator, firstValue, secondValue string
	var counterOfOperators int

	//Считывание введенного выражения и удаление случайных пробелов
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите математическую операцию:")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, " ", "")
	stringLength := len(text)

	//Определение операндов,оператора и проверка на их количество
	for i := 0; i < stringLength; i++ {
		checkedOperator := string(text[i])
		if checkedOperator == "+" || checkedOperator == "-" || checkedOperator == "/" || checkedOperator == "*" {
			if counterOfOperators == 1 {
				fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — " +
					"два операнда и один оператор (+, -, /, *)")
				os.Exit(0)
			}
			firstValue = string(text[0:i])
			secondValue = string(text[i+1 : stringLength])
			operator = checkedOperator
			counterOfOperators += 1
		}
	}
	if counterOfOperators == 0 {
		fmt.Println("Выдача паники, так как строка не является математической операцией.")
		os.Exit(0)
	}
	return operator, firstValue, secondValue
}
