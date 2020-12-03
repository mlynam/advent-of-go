package puzzles

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
)

type tobogganTrajectory struct {
	mx int
	my int
}

func (*tobogganTrajectory) loadMap(file *string) ([][]bool, error) {
	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	reader := strings.NewReader(string(bytes))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var rows [][]bool

	for scanner.Scan() {
		var row []bool
		line := scanner.Text()
		for _, rune := range line {
			row = append(row, rune == '#')
		}

		rows = append(rows, row)
	}

	return rows, nil
}
