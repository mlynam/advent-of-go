package puzzles

import (
	"fmt"

	"github.com/mlynam/advent-of-go/shared"
)

type passwordPhilosophyPart2 struct{}

func (*passwordPhilosophyPart2) Solve(puzzle *shared.Puzzle) {
	shared := passwordPhilosophy{}
	passwords, err := shared.loadPasswords(puzzle.InputFile())
	if err != nil {
		return
	}

	validCount := 0
	for _, metadata := range passwords {
		positionCount := 0
		if string(metadata.password[metadata.min-1]) == metadata.token {
			positionCount++
		}

		if string(metadata.password[metadata.max-1]) == metadata.token {
			positionCount++
		}

		if positionCount == 1 {
			validCount++
		}
	}

	fmt.Printf("Found %v valid passwords.", validCount)
}

// NewPasswordPhilosophyPart2Solver finds valid passwords
func NewPasswordPhilosophyPart2Solver() *shared.Puzzle {
	solver := passwordPhilosophyPart2{}
	return shared.NewPuzzle("2020-day2-password-philosophy-02", &solver)
}
