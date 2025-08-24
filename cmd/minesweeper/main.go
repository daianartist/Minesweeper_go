package main

import (
	"main/internal/game"
	"main/internal/input"
	"main/internal/utils"
)

func main() {
	// Step 1: Read initial choice
	choice := input.ReadChoice()

	var h, w, bombs int
	var field [][]rune

	// Step 2: Depending on choice, read map or generate random
	if choice == 1 {
		h, w = input.ReadSize()
		utils.PrintMessageln("Please be informed that you should plant at least 2 bombs in order to play the game")
		field = input.ReadCustomField(h, w)
		bombs = game.CountBombs(field)
		if h < 3 || w < 3 || bombs < 2 {
			utils.PrintMessageln("Error: invalid input. Press enter to go back.")
			return
		}
		g := game.New(h, w)
		g.SetField(field)
		g.SetBombCount(bombs)
		game.Play(g)
	} else {
		h, w = input.ReadSize()
		bombs = input.ReadBombs(h, w)
		if h < 3 || w < 3 || bombs < 2 {
			utils.PrintMessageln("Error: invalid input. Press enter to go back.")
			return
		}
		g := game.New(h, w)
		g.PlaceRandomBombs(bombs)
		g.SetBombCount(bombs)
		game.Play(g)
	}
}
