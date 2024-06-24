package	piscine

import "ft"

func PrintComb2() {
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			ft.PrintRune(rune(a / 10 + '0'))
			ft.PrintRune(rune(a % 10 + '0'))
			ft.PrintRune(' ')
			ft.PrintRune(rune(b / 10 + '0'))
			ft.PrintRune(rune(b % 10 + '0'))
			if a != 98 {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
	ft.PrintRune('\n')
}