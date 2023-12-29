package main

import (
	"flag"
	"fmt"
	"strconv"
)

/*Упражнение 2.2. Напишите программу общего назначения для преобразования
единиц, аналогичную c f , которая считывает числа из аргументов командной строки
(или из стандартного ввода, если аргументы командной строки отсутствуют) и преобразует каждое число в другие единицы,
как температуру — в градусы Цельсия и Фаренгейта, длину — в футы и метры, вес — в фунты и килограммы и т.д.*/

func main() {
	temperature := flag.Bool("t", false, "Преобразовать в температуру")
	length := flag.Bool("l", false, "Преобразовать в длину")
	weight := flag.Bool("w", false, "Преобразовать в вес")

	flag.Parse()

	args := flag.Args()
	//fmt.Print(args)

	if len(args) == 0 {
		//стандартный ввод
	} else {
		for _, el := range args {
			value, err := strconv.ParseFloat(el, 64)
			if err != nil {
				fmt.Printf("Ошибка при преобразовании аргумента %s: %v\n", value, err)
				return
			}

			if *temperature {
				convertTemperature(value)
			} else if *length {
				convertLength(value)
			} else if *weight {
				convertWeight(value)
			}

		}
	}
}

func convertTemperature(value float64) {
	c := value
	f := (c * 9 / 5) + 32

	fmt.Printf("%.2f°C = %.2f°F\n", c, f)
}

func convertLength(value float64) {
	f := value
	m := f * 0.3048
	fmt.Printf("%.2f feet = %.2f meters\n", f, m)
}

// Преобразование веса
func convertWeight(value float64) {
	p := value
	k := p * 0.45359237

	fmt.Printf("%.2f pounds = %.2f kilograms\n", p, k)
}
