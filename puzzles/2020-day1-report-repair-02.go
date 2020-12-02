package puzzles

import (
	"fmt"
	"sort"

	"github.com/mlynam/advent-of-go/shared"
)

type reportRepairWithThree struct{}

func (*reportRepairWithThree) Solve(puzzle *shared.Puzzle) {
	shared := reportRepair{}
	values, err := shared.loadInput(puzzle.InputFile())
	if err != nil {
		return
	}

	sort.Ints(values)

	for i, value := range values {
		if i+1 == len(values) {
			continue
		}

		diff := 2020 - (value + values[i+1])
		searching := values[i:]
		idx := sort.SearchInts(searching, diff)

		for _, potential := range searching[idx : idx+1] {
			if potential+value+values[i+1] == 2020 {
				answer := potential * value * values[i+1]
				fmt.Printf("%v * %v * %v = %v\n", value, values[i+1], potential, answer)
			}
		}
	}
}

// NewReportRepairWithThreeSolver solves part2 of day1
func NewReportRepairWithThreeSolver() *shared.Puzzle {
	solver := &reportRepairWithThree{}
	return shared.NewPuzzle("2020-day1-report-repair-02", solver)
}
