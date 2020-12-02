package puzzles

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passwordPhilosophy struct{}

type passwordMetadata struct {
	min      int
	max      int
	token    string
	password string
}

func (*passwordPhilosophy) loadPasswords(inputFile *string) ([]passwordMetadata, error) {
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	lineStructure := regexp.MustCompile(`^([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)$`)
	reader := strings.NewReader(string(bytes))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var metadatas []passwordMetadata

	for scanner.Scan() {
		line := scanner.Text()

		if lineStructure.MatchString(line) {
			matches := lineStructure.FindStringSubmatch(line)
			min, err := strconv.Atoi(matches[1])
			if err != nil {
				return nil, err
			}

			max, err := strconv.Atoi(matches[2])
			if err != nil {
				return nil, err
			}

			metadata := passwordMetadata{
				min:      min,
				max:      max,
				token:    matches[3],
				password: matches[4],
			}

			metadatas = append(metadatas, metadata)
		}
	}

	return metadatas, nil
}
