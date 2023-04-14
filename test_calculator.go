package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

var roman_number_int = map[string]int{
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
	"XL":   40,
	"L":    50,
	"XC":   90,
	"C":    100,
}

var numbers_int = []int{100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

var numbers = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

func check_number(sl []string, number string) bool {
	for _, value := range sl {
		if value == number {
			return true
		}
	}
	return false
}

func change_str_to_int(number_1 string, number_2 string) (int, int) {
	number_1_int, err := strconv.Atoi(number_1)
	if err != nil {
		log.Fatal(err)
	}

	number_2_int, err := strconv.Atoi(number_2)
	if err != nil {
		log.Fatal(err)
	}
	return number_1_int, number_2_int
}

func get_result(num_1 int, operator string, num_2 int) int {
	switch operator {
	case "+":
		result := num_1 + num_2
		return result
	case "-":
		result := num_1 - num_2
		return result
	case "*":
		result := num_1 * num_2
		return result
	case "/":
		result := num_1 / num_2
		return result
	default:
		panic("Введен неверный оператор, нужно выбрать из (+, -, *, /)")
	}
}

func main() {
	fmt.Print("Введите два числа и оператор между ними через пробелы(числа могут быть римские(от I до X)" +
		"или арабские(от 1 до 10), операторы допустимы для использования +, -, *, /. ): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	arr := strings.Split(text, " ")
	if len(arr) != 3 {
		panic("Нужно использовать только 2 операнда и 1 оператор. Запустите заного.")
	}

	var number_1, operator, number_2 = arr[0], arr[1], arr[2]

	_, KeyCheck_1 := roman_number[number_1]
	_, KeyCheck_2 := roman_number[number_2]

	if !KeyCheck_1 || !KeyCheck_2 {
		if !check_number(numbers[:], number_1) || !check_number(numbers[:], number_2) {
			panic("Нужно ввести либо две арабские цифры от 1 до 10 включительно или либо две римские цифры от I до X включительно")
		} else {

			number_1_int, number_2_int := change_str_to_int(number_1, number_2)

			fmt.Println(get_result(number_1_int, operator, number_2_int))
		}
	} else {

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

		number_1_int, number_2_int := change_str_to_int(number_1, number_2)

		res := get_result(number_1_int, operator, number_2_int)

		if res <= 0 {
			panic("При вводе римских чисел не может быть ответ меньше или равный 0." +
				"Таких чисел нет в римской системе. Введите другие значения. ")
		}

		var res_num_roman string

		for res > 0 {
			for _, value := range numbers_int {
				for i := value; i <= res; {
					for i, v := range roman_number_int {
						if v == value {
							res_num_roman += i
							res -= value
						}
					}
				}
			}
		}
		fmt.Println(res_num_roman)
	}

}
