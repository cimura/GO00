package	piscine

import "ft"

func	PrintComb() {
	for a := '0'; a <= '9'; a++ {
		for b := a + 1; b <= '9'; b++ {
			for c := b + 1; c <= '9'; c++ {
				ft.PrintRune(a)
				ft.PrintRune(b)
				ft.PrintRune(c)
				if a != '7' {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
			}
		}
	}
	ft.PrintRune('\n')
}