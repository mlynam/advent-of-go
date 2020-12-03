package puzzles

import (
	"fmt"

	"github.com/mlynam/advent-of-go/shared"
)

// Solve the puzzle
func (*tobogganTrajectory) Solve(puzzle *shared.Puzzle) {
	solver := puzzle.Solver().(*tobogganTrajectory)
	rows, err := solver.loadMap(puzzle.InputFile())
	if err != nil {
		fmt.Println(err)
		return
	}

	rowCount := len(rows)
	var (
		x int
		y int
	)

	treeCount := 0
	for y = 0; y < rowCount; y += solver.my {
		row := rows[y]
		if row[x%len(row)] {
			treeCount++
		}

		x += solver.mx
	}

	fmt.Printf("%v trees encountered.\n", treeCount)
}

// NewTobogganTrajectorySolver solves the toboggan trajectory puzzle
func NewTobogganTrajectorySolver(mx int, my int) *shared.Puzzle {
	solver := tobogganTrajectory{mx: mx, my: my}
	return shared.NewPuzzle("2020-day3-toboggan-trajectory-01", &solver)
}
