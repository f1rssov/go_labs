package main

import (
	"bytes"
	"fmt"
	"unicode"
)

/*
Упражнение 3.10. Напишите нерекурсивную версию функции
comma, использующую b y t e s . B u ffe r вместо конкатенации строк.
*/
func comma10(s string) string {
	var buf bytes.Buffer

	for index, value := range s {
		if index > 0 && index%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(value)
	}

	return buf.String()
}

/*
Упражнение 3.11. Усовершенствуйте функцию comma так, чтобы она корректно
работала с числами с плавающей точкой и необязательным знаком
*/
func comma11(s string) string {
	var resultString string

	counterDigit := 0
	for _, value := range s {
		if unicode.IsDigit(value) {
			counterDigit++
		}

		resultString += string(value)

		if counterDigit > 0 && counterDigit%3 == 0 {
			resultString += ","
		}
	}

	return resultString
}

func main() {
	fmt.Println(comma10("123454342234"))
	fmt.Print(comma11("12.341212412"))
}
