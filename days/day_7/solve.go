package day_7

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

func Solve(pathToFile string) error {
	f, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer f.Close()

	nums, err := parse(f)
	if err != nil {
		return err
	}

	err = part1(nums)
	if err != nil {
		return err
	}

	err = part2(nums)
	if err != nil {
		return err
	}

	return nil
}

func part1(nums []int) error {
	sorted := append(make([]int, 0), nums...)
	sort.Ints(sorted)
	mid := len(sorted) / 2

	median := sorted[mid]
	var fuel int
	for _, num := range sorted {
		fuel += int(math.Abs(float64(num - median)))
	}

	fmt.Printf("Total fuel used: %d\n", fuel)
	return nil
}

func part2(nums []int) error {
	var sum int
	for _, num := range nums {
		sum += num
	}
	mean := int(math.Round(float64(sum) / float64(len(nums))))
	var fuelMin int
	var fuel int
	var fuelMax int
	sm := make(sumMap)
	for _, num := range nums {
		fuelMin += sm.sumFuel(int(math.Abs(float64(num - mean - 1))))
		fuel += sm.sumFuel(int(math.Abs(float64(num - mean))))
		fuelMax += sm.sumFuel(int(math.Abs(float64(num - mean + 1))))
	}
	totalFuel := int(math.Min(float64(fuel), math.Min(float64(fuelMax), float64(fuelMin))))

	fmt.Printf("Total fuel used: %d\n", totalFuel)
	return nil
}

type sumMap map[int]int

func (s sumMap) sumFuel(dist int) int {
	if fuel, ok := s[dist]; ok {
		return fuel
	}

	var fuel int
	for i := 1; i <= dist; i++ {
		fuel += i
	}

	s[dist] = fuel
	return fuel
}

func parse(r io.Reader) ([]int, error) {
	csvR := csv.NewReader(r)
	line, err := csvR.Read()
	if err != nil {
		return nil, err
	}

	nums := make([]int, 0)
	for _, s := range line {
		num, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return nil, err
		}
		nums = append(nums, int(num))
	}

	return nums, nil
}
