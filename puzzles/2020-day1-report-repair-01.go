package puzzles

import (
	"fmt"
	"sort"

	"github.com/mlynam/advent-of-go/shared"
)

type reportRepairWithTwo struct{}

func (r *reportRepairWithTwo) Solve(puzzle *shared.Puzzle) {
	shared := ReportRepair{}
	input, err := shared.loadInput(puzzle.InputFile())
	if err != nil {
		return
	}

	sort.Ints(input)

	for i, value := range input {
		diff := 2020 - value
		searching := input[i:]
		idx := sort.SearchInts(searching, diff)

		for _, potential := range searching[idx : idx+1] {
			if potential+value == 2020 {
				answer := potential * value
				fmt.Printf("%v * %v = %v\n", value, potential, answer)
			}
		}
	}
}

// NewReportRepairWithTwoSolver solves part1 of day1
func NewReportRepairWithTwoSolver() *shared.Puzzle {
	solver := &reportRepairWithTwo{}
	return shared.NewPuzzle("2020-day1-report-repair-01", solver)
}
