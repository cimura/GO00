package	piscine

import "ft"

func combination(n int, min int, indx int, array [10]byte) bool {
	for i := min; i <= 10-(n-indx); i++ {
		array[indx] = byte(i + '0')
		if indx == n - 1 {
			for j := 0; j < n; j++ {
				ft.PrintRune(rune(array[j]))
			}
			if array[0] != byte(10 - n + '0') {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		} else {
			combination(n, i + 1, indx + 1, array)
		}
	}
	return true
}

func PrintCombN(n int) {
	var array [10]byte
	combination(n, 0, 0, array)
	ft.PrintRune('\n')
}