package main

import (
	"fmt"
	"os"

	"github.com/mlynam/advent-of-go/puzzles"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("invalid input")
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
		fmt.Println("expected command like '2020/day1' -puzzle=sum2020")
	}
}
