package game

import (
	"math/rand"

	a "main/internal/algo"
	"main/internal/input"
	u "main/internal/utils"
)

// Game holds the state of a minesweeper game.
// nice
type Game struct {
	h, w        int
	field       [][]rune // '.' or '*'
	numbers     [][]int  // -1 for bomb, 0..8 for counts
	revealed    [][]bool
	moves       int
	bombs       int
	firstMove   bool
	revealedCnt int
}

// New creates an empty game with no bombs placed yet.
func New(h, w int) *Game {
	g := &Game{
		h:         h,
		w:         w,
		field:     make([][]rune, h),
		numbers:   make([][]int, h),
		revealed:  make([][]bool, h),
		firstMove: true,
	}
	for i := 0; i < h; i++ {
		g.field[i] = make([]rune, w)
		g.numbers[i] = make([]int, w)
		g.revealed[i] = make([]bool, w)
		for j := 0; j < w; j++ {
			g.field[i][j] = u.EmptyChar
			g.numbers[i][j] = 0
			g.revealed[i][j] = false
		}
	}
	return g
}

// SetField sets the board from a user-provided field (must be h x w).
func (g *Game) SetField(field [][]rune) {
	g.field = field
	if nums, err := a.CalcNumbers(field); err == nil {
		g.numbers = nums
	}
}

// PlaceRandomBombs randomly places `bombCount` bombs on the board.
func (g *Game) PlaceRandomBombs(bombCount int) {
	// ensure field is initialized
	if len(g.field) != g.h {
		g.field = make([][]rune, g.h)
		for i := 0; i < g.h; i++ {
			g.field[i] = make([]rune, g.w)
			for j := 0; j < g.w; j++ {
				g.field[i][j] = u.EmptyChar
			}
		}
	}

	placed := 0
	for placed < bombCount {
		x := rand.Intn(g.h)
		y := rand.Intn(g.w)
		if g.field[x][y] != u.BombChar {
			g.field[x][y] = u.BombChar
			placed++
		}
	}
	// compute numbers
	if nums, err := a.CalcNumbers(g.field); err == nil {
		g.numbers = nums
	}
}

// SetBombCount sets the number of bombs (useful when importing a custom field).
func (g *Game) SetBombCount(b int) { g.bombs = b }

// IncMoves increments the number of moves (called by main loop).
func (g *Game) IncMoves() { g.moves++ }

// Height and Width accessors
func (g *Game) Height() int { return g.h }
func (g *Game) Width() int  { return g.w }

// CountBombs returns the number of bombs inside a provided field.
func CountBombs(field [][]rune) int {
	cnt := 0
	for i := range field {
		for j := range field[i] {
			if field[i][j] == u.BombChar {
				cnt++
			}
		}
	}
	return cnt
}

// revealCascade reveals cell x,y and cascades if zero-adjacent bombs.
func (g *Game) revealCascade(x, y int) {
	if x < 0 || x >= g.h || y < 0 || y >= g.w {
		return
	}
	if g.revealed[x][y] {
		return
	}
	g.revealed[x][y] = true
	if g.numbers[x][y] != -1 {
		g.revealedCnt++
	}
	// if this cell is zero, reveal neighbors
	if g.numbers[x][y] == 0 {
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				g.revealCascade(x+dx, y+dy)
			}
		}
	}
}

// relocateBomb moves a bomb from x,y to a random empty cell (used on first move).
func (g *Game) relocateBomb(x, y int) {
	// remove bomb at x,y
	g.field[x][y] = u.EmptyChar
	empties := make([][2]int, 0)
	for i := 0; i < g.h; i++ {
		for j := 0; j < g.w; j++ {
			if g.field[i][j] == u.EmptyChar && !(i == x && j == y) {
				empties = append(empties, [2]int{i, j})
			}
		}
	}
	if len(empties) == 0 {
		// put it back if nowhere to move
		g.field[x][y] = u.BombChar
		return
	}
	r := rand.Intn(len(empties))
	nx, ny := empties[r][0], empties[r][1]
	g.field[nx][ny] = u.BombChar
	// recompute numbers after moving the bomb
	if nums, err := a.CalcNumbers(g.field); err == nil {
		g.numbers = nums
	}
}

// Step processes a move at zero-based coordinates and returns true if game over.
func (g *Game) Step(x, y int) bool {
	if x < 0 || x >= g.h || y < 0 || y >= g.w {
		return false
	}
	if g.firstMove {
		g.firstMove = false
		if g.numbers[x][y] == -1 {
			g.relocateBomb(x, y)
		}
	}
	// If selected cell is a bomb after relocation check -> game over
	if g.numbers[x][y] == -1 {
		// reveal bomb cell
		g.revealed[x][y] = true
		return true
	}
	if !g.revealed[x][y] {
		g.revealCascade(x, y)
	}
	return false
}

// Won returns true when all non-bomb cells are revealed.
func (g *Game) Won() bool {
	totalSafe := g.h*g.w - g.bombs
	return g.revealedCnt >= totalSafe
}

// RevealAllSafe opens all non-bomb cells (for final display on win).
func (g *Game) RevealAllSafe() {
	for i := 0; i < g.h; i++ {
		for j := 0; j < g.w; j++ {
			if !g.revealed[i][j] && g.field[i][j] != u.BombChar {
				g.revealed[i][j] = true
			}
		}
	}
}

// Stats prints simple statistics about the game.
func (g *Game) Stats() {
	u.PrintMessageln("Your statistics:")
	u.PrintMessageln("- Field size: " + u.Itoa(g.h) + "x" + u.Itoa(g.w) + "\n")
	u.PrintMessageln("- Number of bombs: " + u.Itoa(g.bombs) + "\n")
	u.PrintMessageln("- Number of moves: " + u.Itoa(g.moves) + "\n")
}

// Play runs the main loop (uses input.ReadCoordinates)
func Play(g *Game) {
	for {
		g.Print(false)
		x, y := input.ReadCoordinates(g.Height(), g.Width())
		g.IncMoves()
		gameOver := g.Step(x-1, y-1)
		if gameOver {
			g.Print(true)
			u.PrintMessageln("Game Over!")
			g.Stats()
			return
		}
		if g.Won() {
			g.RevealAllSafe()
			g.Print(false)
			u.PrintMessageln("You Win!")
			g.Stats()
			return
		}
	}
}
