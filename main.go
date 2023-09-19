package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romToArabMap = map[string]int8{

	"I": 1,
	"V": 5,
	"X": 10,
}
var arabToRomMap = map[int64]string{
	1:  "I",
	4:  "IV",
	5:  "V",
	9:  "IX",
	10: "X",
}

func romToArab(romNum string) int8 {
	var result int8 = 0
	for i := 0; i < len(romNum); i++ {
		var value int8 = romToArabMap[string(romNum[i])]
		if i+1 < len(romNum) {
			var nextValue int8 = romToArabMap[string(romNum[i+1])]
			if value < nextValue {
				result -= value

			} else {
				result += value
			}
		} else {
			result += value
		}
	}
	return result
}

func arabToRom(result int64) string {

	var romNum string
	for result > 0 {
		for _, key := range []int64{10, 9, 5, 4, 1} {
			if result >= key {
				result -= key
				romNum += arabToRomMap[key]
				break
			}
		}
	}
	return romNum
}
func mathOper(num1, num2 int64, oper string) (string, int64) {
	switch {
	case oper == "+":
		return "Ответ:", num1 + num2

	case oper == "-":
		return "Ответ:", num1 - num2

	case oper == "*":
		return "Ответ:", num1 * num2

	case oper == "/":
		return "Ответ:", num1 / num2

	default:
		return "Вывод ошибки: Не известрый оператор!", 0
	}
}
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	val := scanner.Text()

	expr := strings.Split(val, " ")

	if len(expr) < 3 {
		fmt.Println("Вывод ошибки: Строка не является математической операцией.")
	} else if len(expr) > 3 {
		fmt.Println("Вывод ошибки: Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	} else {
		num1, err1 := strconv.ParseInt(expr[0], 10, 8)
		num2, err2 := strconv.ParseInt(expr[2], 10, 8)

		switch { // Проверка системы исчисления
		case err1 != nil && err2 != nil:
			num1 = int64(romToArab(expr[0]))
			num2 = int64(romToArab(expr[2]))

			if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
				fmt.Println("Вывод ошибки: Не правильный диапазон чисел.")
			} else {
				_, result := mathOper(num1, num2, expr[1])
				if result < 1 {
					fmt.Println("Вывод ошибки: В римской системе нет отрицательных чисел.")
				}
				fmt.Println(arabToRom(result))
			}

		case err1 != nil && err2 == nil:
			fmt.Println("Вывод ошибки: Используются одновременно разные системы счисления.")
		case err1 == nil && err2 != nil:
			fmt.Println("Вывод ошибки: Используются одновременно разные системы счисления.")
		default:
			if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
				fmt.Println("Вывод ошибки: Не правильный диапазон чисел.")
			} else {
				fmt.Println(mathOper(num1, num2, expr[1]))
			}
		}

	}
}
