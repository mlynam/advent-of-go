package main

import (
	"fmt"
	"os"

	"github.com/mlynam/advent-of-go/puzzles"
)

func main() {
	invalidInputErr := "expected command like\n\t$ advent-of-go 2020-day2-password-philosophy-02 -inputFile .\\inputs\\day2input"

	if len(os.Args) < 2 {
		fmt.Println(invalidInputErr)
		return
	}

	reportRepairWithTwo := puzzles.NewReportRepairWithTwoSolver()
	reportReportWIthThree := puzzles.NewReportRepairWithThreeSolver()
	passwordPhilosophyPartOne := puzzles.NewPasswordPhilosophyPart1Solver()
	passwordPhilosophyPartTwo := puzzles.NewPasswordPhilosophyPart2Solver()

	switch os.Args[1] {
	case reportRepairWithTwo.Name():
		reportRepairWithTwo.Solve()
		break
	case reportReportWIthThree.Name():
		reportReportWIthThree.Solve()
		break
	case passwordPhilosophyPartOne.Name():
		passwordPhilosophyPartOne.Solve()
		break
	case passwordPhilosophyPartTwo.Name():
		passwordPhilosophyPartTwo.Solve()
	default:
		fmt.Println(invalidInputErr)
	}
}
