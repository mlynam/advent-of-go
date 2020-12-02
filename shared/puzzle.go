package shared

import (
	"flag"
	"os"
)

// Solveable instance
type Solveable interface {
	Solve(*Puzzle)
}

// Puzzle describes and intiailizes a puzzle solver
type Puzzle struct {
	name      string
	inputFile *string
	flagSet   *flag.FlagSet
	solver    Solveable
}

// NewPuzzle to solve
func NewPuzzle(name string, solver Solveable) *Puzzle {
	set := flag.NewFlagSet(name, flag.ExitOnError)
	inputFile := set.String("inputFile", "empty", "Path to the file containing the input data for the puzzle")

	puzzle := &Puzzle{
		name:      name,
		inputFile: inputFile,
		solver:    solver,
		flagSet:   set,
	}

	return puzzle
}

// Name for the puzzle
func (puzzle *Puzzle) Name() string {
	return puzzle.name
}

// Solve the puzzle
func (puzzle *Puzzle) Solve() {
	puzzle.flagSet.Parse(os.Args[2:])
	puzzle.solver.Solve(puzzle)
}

// InputFile for the puzzle
func (puzzle *Puzzle) InputFile() *string {
	return puzzle.inputFile
}
