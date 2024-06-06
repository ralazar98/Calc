package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	Solution(GetVarOper())
	test := "Тестовая строка"
	fmt.Printf(test)
}

func Solution(operator, firstValue, secondValue string) {
	firstNumber, _ := strconv.Atoi(firstValue)
	secondNumber, _ := strconv.Atoi(secondValue)

	switch {
	case firstNumber == 0 && secondNumber == 0:
		firstNumber, secondNumber = RomanToArabic(firstValue, secondValue)
		resultArab := Calculation(operator, firstNumber, secondNumber)
		if resultArab < 0 {
			fmt.Println("Выдача паники,так как в римской системе нет отрицательных чисел. ")
			os.Exit(0)
		}
		resultRoman := ArabicToRoman(resultArab)
		fmt.Println(resultRoman)
	case firstNumber == 0 && secondNumber != 0 || firstNumber != 0 && secondNumber == 0:
		fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
	case firstNumber < 1 || firstNumber > 10 || secondNumber < 1 || secondNumber > 10:
		fmt.Println("Выдача паники, как минимум 1 число вне диапазона ")
	default:
		fmt.Println(Calculation(operator, firstNumber, secondNumber))
	}

}

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

func GetVarOper() (string, string, string) {

	var operator, firstValue, secondValue string
	var counterOfOperation int

	//Считывание введенного выражения
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите операцию:")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	stringLength := len(text)

	//Определение чисел и проверка на количество операндов и оператора
	for i := 0; i < stringLength; i++ {
		checkedOperator := string(text[i])
		if checkedOperator == "+" || checkedOperator == "-" || checkedOperator == "/" || checkedOperator == "*" {
			if counterOfOperation == 1 {
				fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — " +
					"два операнда и один оператор (+, -, /, *)")
				os.Exit(0)
			}
			firstValue = string(text[0:i])
			secondValue = string(text[i+1 : stringLength])
			operator = checkedOperator
			counterOfOperation += 1

		}
	}
	if counterOfOperation == 0 {
		fmt.Println("Выдача паники, так как строка не является математической операцией.")
		os.Exit(0)
	}
	return operator, firstValue, secondValue
}
