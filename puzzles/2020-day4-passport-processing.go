package puzzles

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/mlynam/advent-of-go/shared"
)

type passportProcessing struct {
	validate bool
}

// Solve the puzzle
func (pp *passportProcessing) Solve(p *shared.Puzzle) {
	flags, err := pp.loadPassports(p.InputFile())
	if err != nil {
		return
	}

	count := 0
	for _, flag := range flags {
		if flag {
			count++
		}
	}

	fmt.Printf("%v valid passports found.", count)
}

// NewPassportProcessingSolver to solve day 4
func NewPassportProcessingSolver(validate bool) *shared.Puzzle {
	solver := passportProcessing{validate: validate}
	return shared.NewPuzzle("2020-day4-passport-processing", &solver)
}

func (pp *passportProcessing) loadPassports(file *string) ([]bool, error) {
	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	reader := strings.NewReader(string(bytes))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var (
		byr bool
		iyr bool
		eyr bool
		hgt bool
		hcl bool
		ecl bool
		pid bool
	)

	var passportFlags []bool

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			passportFlags = append(passportFlags, byr && iyr && eyr && hgt && hcl && ecl && pid)
			byr = false
			iyr = false
			eyr = false
			hgt = false
			hcl = false
			ecl = false
			pid = false
		}

		pairs := strings.Split(line, " ")
		hclrx := regexp.MustCompile(`^#[0-9a-f]{6}$`)

		for _, pair := range pairs {
			parts := strings.Split(pair, ":")
			switch parts[0] {
			case "byr":
				if pp.validate {
					byr = numberValidator(1920, 2002, parts[1])
				} else {
					byr = true
				}
				break
			case "iyr":
				if pp.validate {
					iyr = numberValidator(2010, 2020, parts[1])
				} else {
					iyr = true
				}
				break
			case "eyr":
				if pp.validate {
					eyr = numberValidator(2020, 2030, parts[1])
				} else {
					eyr = true
				}
				break
			case "hgt":
				if pp.validate {
					if strings.HasSuffix(parts[1], "cm") {
						hgt = numberValidator(150, 193, strings.TrimSuffix(parts[1], "cm"))
					} else if strings.HasSuffix(parts[1], "in") {
						hgt = numberValidator(59, 76, strings.TrimSuffix(parts[1], "in"))
					} else {
						hgt = false
					}
				} else {
					hgt = true
				}
				break
			case "hcl":
				if pp.validate {
					hcl = hclrx.MatchString(parts[1])
				} else {
					hcl = true
				}
				break
			case "ecl":
				if pp.validate {
					ecl = validateEyeColor(parts[1])
				} else {
					ecl = true
				}
				break
			case "pid":
				if pp.validate {
					_, err := strconv.Atoi(strings.TrimLeft(parts[1], "0"))
					pid = len(parts[1]) == 9 && err == nil
				} else {
					pid = true
				}
				break
			}
		}
	}

	return passportFlags, nil
}

func numberValidator(min, max int, value string) bool {
	integer, err := strconv.Atoi(value)
	return err == nil && integer >= min && integer <= max
}

func validateEyeColor(value string) bool {
	for _, color := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if color == value {
			return true
		}
	}

	return false
}
