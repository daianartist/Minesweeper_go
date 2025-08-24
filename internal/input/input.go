package input

import (
	"fmt"
	. "main/internal/utils"
)

// ReadChoice reads the initial mode choice.
func ReadChoice() int {
	var choice int
	for {
		PrintMessageln("Choose a mode:")
		PrintMessageln("1. Enter a custom map")
		PrintMessageln("2. Generate a random map")
		PrintMessage("Enter your choice: ")
		_, err := fmt.Scanln(&choice)
		if err == nil && (choice == 1 || choice == 2) {
			return choice
		}
		PrintMessageln("Error: invalid input. Press Enter to go back.")
		var discard string
		fmt.Scanln(&discard) // clear invalid token
	}
}

// ReadSize reads two integers h and w with basic validation (h,w >= 3).
func ReadSize() (int, int) {
	var h, w int
	for {
		PrintMessageln("Enter size (h w): ")
		_, err := fmt.Scanf("%d %d\n", &h, &w)
		if err == nil && h >= 3 && w >= 3 {
			return h, w
		}
		PrintMessageln("Error: invalid input. Press enter to go back.")
		var discard string
		fmt.Scanln(&discard)
	}
}

// ReadBombs reads an integer number of bombs (>=2 and < h*w).
func ReadBombs(h, w int) int {
	var b int
	for {
		PrintMessageln("Enter number of bombs (>=2): ")
		_, err := fmt.Scanf("%d\n", &b)
		if err == nil && b >= 2 && b < h*w {
			return b
		}
		PrintMessageln("Error: invalid input. Press Enter to go back.")
		var discard string
		fmt.Scanln(&discard)
	}
}

// ReadCustomField reads h rows each with exactly w characters ('.' or '*').
func ReadCustomField(h, w int) [][]rune {
	field := make([][]rune, h)
	for i := 0; i < h; i++ {
		for {
			PrintMessageln("Enter row " + Itoa(i+1) + " (exactly " + Itoa(w) + "chars, '.' or '*'): ")
			var line string
			_, err := fmt.Scanf("%s\n", &line)
			if err != nil || len(line) != w {
				PrintMessageln("row length mismatch")
				continue
			}
			ok := true
			row := make([]rune, w)
			for j, ch := range line {
				if ch != BombChar && ch != EmptyChar {
					ok = false
					break
				}
				row[j] = ch
			}
			if !ok {
				PrintMessageln("invalid character")
				continue
			}
			field[i] = row
			break
		}
	}
	return field
}

// ReadCoordinates reads two integers representing row and column (1-based) and validates them.
func ReadCoordinates(maxH, maxW int) (int, int) {
	var x, y int
	for {
		PrintMessage("Enter coordinates(row then col): ")
		_, err := fmt.Scanf("%d %d\n", &x, &y)
		if err == nil && x >= 1 && x <= maxH && y >= 1 && y <= maxW {
			return x, y
		}
		PrintMessageln("Invalid input. Press Enter to go back. ")
		var discard string
		fmt.Scanf("%s\n", &discard)
	}
}
