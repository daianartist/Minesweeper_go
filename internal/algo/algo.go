package algo

import (
	"errors"
	u "main/internal/utils"
)

// InBounds checks coordinates inside the grid.
func InBounds(r, c, rows, cols int) bool {
	return r >= 0 && r < rows && c >= 0 && c < cols
}

// CalcNumbers takes a rune field ('.' empty, '*' bomb) and returns an int grid:
// -1 for bombs, 0..8 for number of adjacent bombs.
func CalcNumbers(field [][]rune) ([][]int, error) {
	rows := len(field)
	if rows == 0 {
		return nil, errors.New("empty field")
	}
	cols := len(field[0])
	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		if len(field[i]) != cols {
			return nil, errors.New("inconsistent row lengths")
		}
		grid[i] = make([]int, cols)
	}

	dr := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dc := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if field[i][j] == u.BombChar {
				grid[i][j] = -1
				continue
			}
			cnt := 0
			for d := 0; d < 8; d++ {
				r := i + dr[d]
				c := j + dc[d]
				if InBounds(r, c, rows, cols) && field[r][c] == u.BombChar {
					cnt++
				}
			}
			grid[i][j] = cnt
		}
	}
	return grid, nil
}

// FromRaw converts raw string rows to [][]rune field; validates characters.
func FromRaw(raw []string, rows, cols int) ([][]rune, error) {
	if len(raw) != rows {
		return nil, errors.New("raw row count mismatch")
	}
	field := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		line := raw[i]
		if len(line) != cols {
			return nil, errors.New("raw row length mismatch")
		}
		row := make([]rune, cols)
		for j, ch := range line {
			if rune(ch) != u.BombChar && rune(ch) != u.EmptyChar {
				return nil, errors.New("invalid character in raw")
			}
			row[j] = rune(ch)
		}
		field[i] = row
	}
	return field, nil
}
