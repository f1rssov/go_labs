package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// 2.3 Перепишем функцию PopCount с использованием цикла:
func PopCount3(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

// 2.4 Напишем версию PopCount с использованием сдвига аргумента по всем 64 позициям:
func PopCount4(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if (x & 1) == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

// 2.5 х&(х-1):
func PopCount5(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
