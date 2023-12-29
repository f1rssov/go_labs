package main

import "fmt"

func reverse(slice_byte []byte) {
	length := len(slice_byte) / 2

	for i := 0; i < length; i++ {
		reverseIndex := len(slice_byte) - i - 1
		slice_byte[i], slice_byte[reverseIndex] = slice_byte[reverseIndex], slice_byte[i]
	}
}

func main() {
	input := []byte("hello world!")
	reverse(input)

	fmt.Print(string(input))
}
