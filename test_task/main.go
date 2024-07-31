package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter an expression: ")
	expression, _ := reader.ReadString('\n')

	result := calculate(expression)
	fmt.Printf("Result: %f\n", result)
}

func calculate(expression string) float64 {
	expression = strings.TrimSuffix(expression, "\n")

	// разбиение  на операнды и оператора
	tokens := strings.Split(expression, " ")
	if len(tokens) != 3 {
		fmt.Println("Выдача паники, так как формат математической  операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		os.Exit(1)
	}

	var result float64

	// Определение системы счисления
	ArabicNumbers := "1234567890"
	containsA := strings.ContainsAny(tokens[0], ArabicNumbers)
	containsB := strings.ContainsAny(tokens[2], ArabicNumbers)

	if containsA == true && containsB == true {

		operand1, err := strconv.ParseFloat(tokens[0], 64)
		if err != nil {
			fmt.Println("Invalid operand 1:", err)
			os.Exit(1)
		}

		operand2, err := strconv.ParseFloat(tokens[2], 64)
		if err != nil {
			fmt.Println("Invalid operand 2:", err)
			os.Exit(1)
		}

		operator := tokens[1]

		switch operator {
		case "+":
			result = operand1 + operand2
		case "-":
			result = operand1 - operand2
		case "*":
			result = operand1 * operand2
		case "/":
			if operand2 == 0 {
				fmt.Println("Division by zero")
				os.Exit(1)
			}
			result = operand1 / operand2
		default:
			fmt.Println("Invalid operator:", operator)
			os.Exit(1)
		}

	} else if containsA == false && containsB == false {
		operand1 := romanToArabic(tokens[0])
		operand2 := romanToArabic(tokens[2])
		operator := tokens[1]

		switch operator {
		case "+":
			result = operand1 + operand2
		case "-":
			result = operand1 - operand2
		case "*":
			result = operand1 * operand2
		case "/":
			if operand2 == 0 {
				fmt.Println("Division by zero")
				os.Exit(1)
			}
			result = operand1 / operand2
		default:
			fmt.Println("Invalid operator:", operator)
			os.Exit(1)
		}

	} else {
		fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
	}
	return result
}

func romanToArabic(roman string) int {
	romanMap := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	arabic := 0
	i := 0
	for i < len(roman) {
		if i+1 < len(roman) && romanMap[roman[i:i+2]] > 0 {
			arabic += romanMap[roman[i:i+2]]
			i += 2
		} else {
			arabic += romanMap[string(roman[i])]
			i++
		}
	}
	return arabic
}
