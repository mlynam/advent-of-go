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
	case "2020-day3-toboggan-trajectory-01":
		puzzles.NewTobogganTrajectorySolver(3, 1).Solve()
	case "2020-day3-toboggan-trajectory-02":
		puzzles.NewTobogganTrajectorySolver(1, 1).Solve()
		puzzles.NewTobogganTrajectorySolver(3, 1).Solve()
		puzzles.NewTobogganTrajectorySolver(5, 1).Solve()
		puzzles.NewTobogganTrajectorySolver(7, 1).Solve()
		puzzles.NewTobogganTrajectorySolver(1, 2).Solve()
	case "2020-day4-passport-processing-01":
		puzzles.NewPassportProcessingSolver(false).Solve()
	case "2020-day4-passport-processing-02":
		puzzles.NewPassportProcessingSolver(true).Solve()
	default:
		fmt.Println(invalidInputErr)
	}
}
