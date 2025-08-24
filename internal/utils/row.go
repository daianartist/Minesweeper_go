package utils

import "github.com/alem-platform/ap"

func Row(r rune) {
	ap.PutRune('|')
	for i := 0; i < 7; i++ {
		ap.PutRune(r)
	}
}

func FreeSpace(n int) {
	for i := 0; i < n; i++ {
		ap.PutRune(' ')
	}
}

func CharacterRow(r rune) {
	ap.PutRune('|')
	for i := 0; i < 6; i++ {
		if i == 3 {
			ap.PutRune(r)
		}
		ap.PutRune(' ')
	}
}
