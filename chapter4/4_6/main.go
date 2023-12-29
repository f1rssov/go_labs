package main

import "fmt"

func f(slice_byte []byte) []byte {
	var result_slice []byte

	for i := 0; i < len(slice_byte); {
		if slice_byte[i] == ' ' {
			for slice_byte[i] == ' ' {
				i++
			}
			result_slice = append(result_slice, ' ')
		} else {
			result_slice = append(result_slice, slice_byte[i])
			i++
		}
	}

	return result_slice
}

func main() {
	input := []byte("   Hello    world !!!!")
	fmt.Println(string(f(input)))
}
