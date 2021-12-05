package day_4

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Solve(pathToFile string) error {
	f, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer f.Close()

	nums, boards, err := parse(f)
	if err != nil {
		return err
	}

	part1And2(nums, boards)

	return nil
}

func part1And2(nums []int, boards []board) {
	first := true
	firstWin := map[int]struct{}{}
	winningAnswers := make([]int, 0)
	for _, num := range nums {
		for bNdx := range boards {
			boards[bNdx].Mark(num)
			if boards[bNdx].isBingo() {
				sum := boards[bNdx].sumUnmarked()
				if first {
					first = false
					fmt.Printf("First Winning board #%d | Winning num: %d | Sum: %d | Answer: %d\n", bNdx, num, sum, sum*num)
				}
				if _, ok := firstWin[bNdx]; !ok {
					firstWin[bNdx] = struct{}{}
					winningAnswers = append(winningAnswers, sum*num)
				}
			}
		}
	}
	fmt.Printf("Last winning answer: %d\n", winningAnswers[len(winningAnswers)-1])
}

func parse(r io.Reader) ([]int, []board, error) {
	scanner := bufio.NewScanner(r)

	if !scanner.Scan() {
		return nil, nil, errors.New("err no scan")
	}
	numsLine := scanner.Text()
	numsStr := strings.Split(numsLine, ",")
	nums := make([]int, 0)
	for _, s := range numsStr {
		num, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return nil, nil, err
		}
		nums = append(nums, int(num))
	}
	// remove empty line between called numbers and boards
	if !scanner.Scan() {
		_ = scanner.Text()
	}

	boards := make([]board, 0, 3)
	squareNums := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			squares := [25]square{}
			for i, s := range squareNums {
				n, err := strconv.ParseInt(s, 10, 32)
				if err != nil {
					return nil, nil, err
				}
				squares[i].number = int(n)
			}
			boards = append(boards, board{squares: squares})
			squareNums = make([]string, 0)
			continue
		}

		line = scanner.Text()
		parts := strings.Split(line, " ")
		for _, part := range parts {
			if part != "" {
				squareNums = append(squareNums, part)
			}
		}
	}

	return nums, boards, nil
}
