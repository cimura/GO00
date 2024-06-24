package piscine

import "ft"

func Printalphabet() {
	for r := 'a'; r <= 'z'; r++ {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}