package puzzles

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type ReportRepair struct{}

func (solver *ReportRepair) loadInput(file *string) ([]int, error) {
	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	reader := strings.NewReader(string(bytes))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var input []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}

		input = append(input, x)
	}

	return input, nil
}
