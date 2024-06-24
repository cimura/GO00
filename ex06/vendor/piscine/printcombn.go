package	piscine

import "ft"

func combination(n int, indx int, min int) {
	if indx == n {
		return
	}
	for i := min; i <= min + n; i++ {
		ft.PrintRune(rune(i + '0'))
		ft.PrintRune(rune(i + '0'))
		ft.PrintRune(rune(i + '0'))
	}
	ft.PrintRune(',')
	ft.PrintRune(' ')
	combination(n, indx + 1, min + 1)
}

func PrintCombN(n int) {
	combination(n, 0, 0)
}