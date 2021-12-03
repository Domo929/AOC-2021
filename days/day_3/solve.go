package day_3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(pathToFile string) error {
	f, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer f.Close()

	inputs, err := parse(f)
	if err != nil {
		return err
	}

	err = part1(inputs)
	if err != nil {
		return err
	}

	err = part2(inputs)
	if err != nil {
		return err
	}

	return nil
}

func part1(inputs []string) error {
	binLength, mask := maskAndLength(inputs[0])

	zeroCounts := make([]int, binLength)
	oneCounts := make([]int, binLength)

	for _, line := range inputs {
		for i, bin := range line {
			switch bin {
			case '1':
				oneCounts[i]++
			case '0':
				zeroCounts[i]++
			}
		}
	}

	var gammaStr string
	for i := 0; i < binLength; i++ {
		if zeroCounts[i] > oneCounts[i] {
			gammaStr += "0"
		} else {
			gammaStr += "1"
		}
	}

	gamma, err := strconv.ParseUint(gammaStr, 2, binLength)
	if err != nil {
		return err
	}
	epsilon := ^gamma & mask

	fmt.Printf("gamma: %d | epsilon: %d | multible: %d\n", gamma, epsilon, gamma*epsilon)
	return nil
}

func part2(inputs []string) error {
	oxygenMatch := func(zeros, ones int) string {
		mostCommon := "1"
		if zeros > ones {
			mostCommon = "0"
		}
		return mostCommon
	}

	co2Match := func(zeros, ones int) string {
		leastCommon := "0"
		if ones < zeros {
			leastCommon = "1"
		}
		return leastCommon
	}

	oxygen := matchPart2(inputs, oxygenMatch)
	co2 := matchPart2(inputs, co2Match)

	oxyNum, err := strconv.ParseUint(oxygen, 2, 12)
	if err != nil {
		return err
	}
	co2Num, err := strconv.ParseUint(co2, 2, 12)
	if err != nil {
		return err
	}

	fmt.Printf("Oxygen: %d | CO2: %d | multible: %d\n", oxyNum, co2Num, oxyNum*co2Num)

	return nil
}

func matchPart2(inputs []string, determineMask func(int, int) string) string {
	filtered := append(make([]string, 0), inputs...)
	for i := 0; true; i++ {
		if len(filtered) == 1 {
			return filtered[0]
		}
		var zeros int
		var ones int
		for _, s := range filtered {
			char := string(s[i])
			switch char {
			case "1":
				ones++
			case "0":
				zeros++
			}
		}

		match := determineMask(zeros, ones)

		newFiltered := make([]string, 0)
		for _, s := range filtered {
			if string(s[i]) == match {
				newFiltered = append(newFiltered, s)
			}
		}

		filtered = newFiltered
	}
	return ""
}

func parse(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func maskAndLength(input string) (int, uint64) {
	binLength := len(input)

	var mask uint64
	for i := 0; i < binLength; i++ {
		mask = mask | 1<<i
	}

	return binLength, mask
}
