package main

import "fmt"

func a(slice_string []string) {
	for index := 0; index < len(slice_string)-1; index++ {
		if slice_string[index] == slice_string[index+1] {
			slice_string[index] = ""
		}
	}
}

func b(slice_string []string) []string {
	j := 0
	for i := 1; i < len(slice_string); i++ {
		if slice_string[j] != slice_string[i] {
			j++
			slice_string[j] = slice_string[i]
		}
	}

	return slice_string[:j+1]
}

func main() {
	aa := []string{"sasha", "sasha", "afsasf", "car", "rac"}
	a(aa)
	fmt.Println(aa)

	strings := []string{"a", "a", "b", "c", "c", "c", "d", "d", "e"}
	result := b(strings)
	fmt.Println(result)
}
