package main

import (
	"fmt"
	"log"
	"strconv"
)

var roman_number = map[string]string{
	"I":    "1",
	"II":   "2",
	"III":  "3",
	"IV":   "4",
	"V":    "5",
	"VI":   "6",
	"VII":  "7",
	"VIII": "8",
	"IX":   "9",
	"X":    "10",
}

var numbers = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

func check_number(sl []string, number string) bool {
	for _, value := range sl {
		if value == number {
			return true
		}
	}
	return false
}

func main() {
	var number_1, operator, number_2 string

	fmt.Print("Введите два числа и оператор между ними через пробелы(числа могут быть римские(от I до X)" +
		"или арабские(от 1 до 10), операторы допустимы для использования +, -, *, /. ): ")
	fmt.Scanln(&number_1, &operator, &number_2)

	_, KeyCheck_1 := roman_number[number_1]
	_, KeyCheck_2 := roman_number[number_2]

	if !KeyCheck_1 || !KeyCheck_2 {
		if !check_number(numbers[:], number_1) || !check_number(numbers[:], number_2) {
			panic("Нужно ввести либо две арабские цифры от 1 до 10 включительно или либо две римские цифры от I до X включительно")
		}
	}

	for key := range roman_number {
		if number_1 == key {
			number_1 = roman_number[key]
			for key := range roman_number {
				if number_2 == key {
					number_2 = roman_number[key]
				}
			}

		}
	}

	number_1_int, err := strconv.Atoi(number_1)
	if err != nil {
		log.Fatal(err)
	}

	number_2_int, err := strconv.Atoi(number_2)
	if err != nil {
		log.Fatal(err)
	}

	switch operator {
	case "+":
		fmt.Println(number_1_int + number_2_int)
	case "-":
		fmt.Println(number_1_int - number_2_int)
	case "*":
		fmt.Println(number_1_int * number_2_int)
	case "/":
		fmt.Println(number_1_int / number_2_int)
	default:
		fmt.Println("Введен неверный оператор, нужно выбрать из (+, -, *, /)")
	}

}
