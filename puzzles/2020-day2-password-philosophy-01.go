package puzzles

import (
	"fmt"
	"strings"

	"github.com/mlynam/advent-of-go/shared"
)

type passwordPhilosophyPart1 struct{}

func (*passwordPhilosophyPart1) Solve(puzzle *shared.Puzzle) {
	shared := passwordPhilosophy{}
	passwords, err := shared.loadPasswords(puzzle.InputFile())
	if err != nil {
		return
	}

	validCount := 0
	for _, metadata := range passwords {
		count := strings.Count(metadata.password, metadata.token)
		if count >= metadata.min && count <= metadata.max {
			validCount++
		}
	}

	fmt.Printf("Found %v valid passwords.", validCount)
}

// NewPasswordPhilosophyPart1Solver finds valid passwords
func NewPasswordPhilosophyPart1Solver() *shared.Puzzle {
	solver := passwordPhilosophyPart1{}
	return shared.NewPuzzle("2020-day2-password-philosophy-01", &solver)
}
