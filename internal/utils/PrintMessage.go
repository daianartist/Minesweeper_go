package utils

import "github.com/alem-platform/ap"

func PrintMessage(s string) {
	for _, j := range s {
		ap.PutRune(j)
	}
}

func PrintMessageln(s ...string) {
	if len(s) == 0 {
		ap.PutRune('\n')
		return
	}
	for i := 0; i < len(s); i++ {
		PrintMessage(s[i])
	}
	ap.PutRune('\n')
}
