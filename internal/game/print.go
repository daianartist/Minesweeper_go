package game

import (
	u "main/internal/utils"

	"github.com/alem-platform/ap"
)

// drawTopBorder prints the top border line of the board
func drawTopBorder(cols int) {
	ap.PutRune(' ')
	for i := 0; i < 8*cols-1; i++ {
		ap.PutRune('_')
	}
	ap.PutRune('\n')
}

// Print prints the board with nicely formatted cells and headers
// If revealAllBombs == true, bombs are shown regardless of revealed state.
func (g *Game) Print(revealAllBombs bool) {
	// Print column headers
	u.FreeSpace(4)
	for cols := 0; cols < g.w; cols++ {
		u.FreeSpace(3)
		ap.PutRune(rune('1' + cols))
		u.FreeSpace(4)
	}
	ap.PutRune('\n')

	// Top border
	u.FreeSpace(3)
	drawTopBorder(g.w)

	for rows := 0; rows < g.h; rows++ {
		// Top part of each row (row number not yet printed)
		u.FreeSpace(3)
		for cols := 0; cols < g.w; cols++ {
			cellRevealed := g.revealed[rows][cols]
			cellHasBomb := g.numbers[rows][cols] == -1
			showBomb := revealAllBombs && cellHasBomb

			if cellRevealed || showBomb {
				u.Row(' ') // top part is blank for revealed cells
			} else {
				u.Row('x') // unrevealed cells
			}
		}
		ap.PutRune('|')
		ap.PutRune('\n')

		// Middle part (row number + content)
		u.FreeSpace(1)
		ap.PutRune(rune('1' + rows))
		u.FreeSpace(1)
		for cols := 0; cols < g.w; cols++ {
			cellRevealed := g.revealed[rows][cols]
			cellHasBomb := g.numbers[rows][cols] == -1
			showBomb := revealAllBombs && cellHasBomb

			if cellRevealed || showBomb {
				if cellHasBomb {
					u.CharacterRow('*') // show bomb
				} else {
					cnt := g.numbers[rows][cols]
					if cnt == 0 {
						u.CharacterRow(' ')
					} else {
						u.CharacterRow(rune('0' + cnt))
					}
				}
			} else {
				u.Row('x') // unrevealed cell
			}
		}
		ap.PutRune('|')
		ap.PutRune('\n')

		// Bottom part of each row
		u.FreeSpace(3)
		for cols := 0; cols < g.w; cols++ {
			cellRevealed := g.revealed[rows][cols]
			cellHasBomb := g.numbers[rows][cols] == -1
			showBomb := revealAllBombs && cellHasBomb

			if cellRevealed || showBomb {
				u.Row('_')
			} else {
				u.Row('x')
			}
		}
		ap.PutRune('|')
		ap.PutRune('\n')
	}
}
