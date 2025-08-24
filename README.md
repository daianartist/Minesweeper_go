# Minesweeper Game in Go

This is an ASCII-based implementation of the classic Minesweeper game written in Go. The program supports both user-defined and randomly generated game grids, adhering to standard Minesweeper rules with additional features like colored number output and robust input validation.

## Features

- **Game Modes**:
  - **Custom Map**: Users can input their own grid with bombs (`*`) and empty cells (`.`).
  - **Random Map Generation**: Generates a grid with randomly placed bombs based on user-specified dimensions and bomb count.
- **Game Mechanics**:
  - Validates grid size (minimum 3x3) and bomb count (minimum 2).
  - Handles first-move bomb relocation to ensure a safe start.
  - Implements cascade reveal for empty cells with no adjacent bombs.
  - Displays game statistics (field size, number of bombs, moves) on win or loss.
- **Visuals**:
  - ASCII-based grid with clear borders and column/row labels.
  - Colored numbers for adjacent bomb counts to enhance readability.
- **Error Handling**:
  - Robust input validation for grid dimensions, bomb counts, and coordinates.
  - Graceful handling of invalid inputs with clear error messages.

## Requirements

- **Go**: Version 1.16 or higher.
- **Dependencies**: Uses the `github.com/alem-platform/ap` package for ASCII printing and `math/rand` for random bomb placement.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/alem-platform/ap
   ```
2. Navigate to the project directory:
   ```bash
   cd minesweeper
   ```
3. Ensure the `ap` package is installed:
   ```bash
   go get github.com/alem-platform/ap
   ```

## Usage

1. Run the program:
   ```bash
   go run main.go
   ```
2. Choose a game mode:
   - Enter `1` for a custom map and input the grid size and layout.
   - Enter `2` for a random map, then specify the grid size and number of bombs.
3. Play the game by entering coordinates (row, column) to reveal cells.
4. The game ends with a "You Win" or "Game Over" message, displaying statistics.

### Example Input/Output

**Start the game**:
```
Choose a mode:
1. Enter a custom map
2. Generate a random map
Enter your choice: 2
Enter the size of the grid: 5 5
Enter number of bombs (>=2): 5
```

**Initial Grid**:
```
      1       2       3       4       5
   _______________________________________
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
1 |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
2 |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
3 |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
4 |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
5 |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
Enter coordinates:
```

**After entering coordinates `1 1`**:
```
      1       2       3       4       5
   _______________________________________
  |       |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
1 |   1   |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |_______|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
2 |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
3 |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
4 |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
5 |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
Enter coordinates:
```

**Game Over Example**:
```
Enter coordinates: 2 2
      1       2       3       4       5
   _______________________________________
  |       |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
1 |   1   |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |_______|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|       |XXXXXXX|XXXXXXX|XXXXXXX|
2 |XXXXXXX|   *   |XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|_______|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
3 |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
4 |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
5 |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
  |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
Game Over!
Your statistics:
- Field size: 5x5
- Number of bombs: 5
- Number of moves: 2
```

## Code Structure

- **`main.go`**: Entry point, handles user choice (custom or random map) and initializes the game.
- **`game/game.go`**: Core game logic, including game state management, bomb placement, and move processing.
- **`game/print.go`**: Handles grid printing with ASCII formatting and colored number output.
- **`algo/algo.go`**: Utility functions for grid validation and bomb adjacency calculations.
- **`input/input.go`**: Manages user input for grid size, bomb count, and coordinates with validation.
- **`utils/utils.go`**: Helper functions for printing and string conversion.

## Development Process

The project was developed using the **Double Diamond** approach:
1. **Discover**: Analyzed Minesweeper rules and edge cases, exploring user interaction needs and ASCII display constraints.
2. **Define**: Defined core requirements (grid validation, bomb relocation, cascade reveal) and bonus features (random map, colored numbers).
3. **Develop**: Implemented modular code with separate packages for game logic, input, and printing, ensuring reusability and readability.
4. **Deliver**: Tested extensively with various grid sizes, bomb placements, and invalid inputs to ensure robust functionality.

## Future Improvements

- Add difficulty levels for random map generation (e.g., easy, medium, hard).
- Enhance color customization for numbers.
- Implement a replay feature to start a new game without restarting the program.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.# Minesweeper_go
