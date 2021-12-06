package day_6

import (
	"encoding/csv"
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

	startingNums, err := parse(f)
	if err != nil {
		return err
	}

	err = produceCounts(startingNums, 80)
	if err != nil {
		return err
	}

	err = produceCounts(startingNums, 256)
	if err != nil {
		return err
	}

	return nil
}

func produceSlice(fish []uint8, generations int) error {
	for i := 0; i < generations; i++ {
		newFish := make([]uint8, 0)
		for ndx, ttb := range fish {
			if ttb == 0 {
				fish[ndx] = 6
				newFish = append(newFish, 8)
			} else {
				fish[ndx]--
			}
		}
		fish = append(fish, newFish...)
		//fmt.Printf("Gen: %d | Num: %d\n", i+1, len(fish))
	}
	fmt.Println(len(fish))
	return nil
}

func produceCounts(fish []uint8, generations int) error {
	genToBirth := make([]int64, 9)
	var total int
	for _, f := range fish {
		genToBirth[f]++
		total++
	}
	for i := 0; i < generations; i++ {
		births := genToBirth[0]
		total += int(births)

		genToBirth = append(genToBirth[1:9], births)
		genToBirth[6] = genToBirth[6] + births
	}

	fmt.Println(total)
	return nil
}

func parse(r io.Reader) ([]uint8, error) {
	csvR := csv.NewReader(r)
	line, err := csvR.Read()
	if err != nil {
		return nil, err
	}

	nums := make([]uint8, 0)
	for _, s := range line {
		num, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return nil, err
		}
		nums = append(nums, uint8(num))
	}

	return nums, nil
}
