package day_1

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(pathToInput string) error {
	f, err := os.Open(pathToInput)
	if err != nil {
		return err
	}
	defer f.Close()

	depthReadings, err := parse(f)
	if err != nil {
		return err
	}

	err = solvePart1(depthReadings)
	if err != nil {
		return err
	}

	err = solvePart2(depthReadings)
	if err != nil {
		return err
	}

	return nil

}

func parse(r io.Reader) ([]int, error) {
	csvR := csv.NewReader(r)

	readings := make([]int, 0)
	for {
		line, err := csvR.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if len(line) != 1 {
			return nil, errors.New("illegal length of input line")
		}

		i, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, err
		}
		readings = append(readings, i)
	}

	return readings, nil
}

func solvePart1(depthReadings []int) error {
	var numIncreases int
	for i, cur := range depthReadings {
		if i == 0 {
			continue
		}
		prev := depthReadings[i-1]

		if cur > prev {
			numIncreases++
		}
	}

	fmt.Printf("Simple increases: %d\n", numIncreases)

	return nil
}

func solvePart2(depthReadings []int) error {
	buckets := make([]int, 0)

	for i := 2; i < len(depthReadings); i++ {
		bucketedDepth := depthReadings[i] + depthReadings[i-1] + depthReadings[i-2]

		buckets = append(buckets, bucketedDepth)
	}

	var numIncreases int
	for i, bucket := range buckets {
		if i == 0 {
			continue
		}
		if bucket > buckets[i-1] {
			numIncreases++
		}
	}

	fmt.Printf("Bucketed increases: %d\n", numIncreases)

	return nil
}
