package	piscine

import "ft"

func	Printreversealalphabet() {
	for r := 'z'; r >= 'a'; r-- {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

