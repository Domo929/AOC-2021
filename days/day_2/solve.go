package day_2

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Direction string

const (
	Forward Direction = "forward"
	Down    Direction = "down"
	Up      Direction = "up"
)

type Movement struct {
	Direction Direction
	Amount    int
}

var (
	ErrUnknownMovement = errors.New("unknown movement error")
)

func Solve(pathToFile string) error {
	f, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer f.Close()

	movements, err := parse(f)
	if err != nil {
		return err
	}

	err = part1(movements)
	if err != nil {
		return err
	}

	err = part2(movements)
	if err != nil {
		return err
	}

	return nil
}

func part1(movs []Movement) error {
	depth, hor := 0, 0
	for _, mov := range movs {
		switch mov.Direction {
		case Forward:
			hor += mov.Amount
		case Down:
			depth += mov.Amount
		case Up:
			depth -= mov.Amount
		default:
			return ErrUnknownMovement
		}
	}

	fmt.Printf("Part 1 - Total Depth: %d | Total Horizontal: %d | Mulitple: %d\n", depth, hor, depth*hor)

	return nil
}

func part2(movs []Movement) error {
	depth, hor, aim := 0, 0, 0
	for _, mov := range movs {
		switch mov.Direction {
		case Forward:
			hor += mov.Amount
			depth += aim * mov.Amount
		case Down:
			aim += mov.Amount
		case Up:
			aim -= mov.Amount
		default:
			return ErrUnknownMovement
		}
	}

	fmt.Printf("Part 2 - Total Depth: %d | Total Horizontal: %d | Mulitple: %d\n", depth, hor, depth*hor)

	return nil
}

func parse(r io.Reader) ([]Movement, error) {
	csvR := csv.NewReader(r)
	csvR.Comma = ' '

	lines, err := csvR.ReadAll()
	if err != nil {
		return nil, err
	}

	movements := make([]Movement, 0, len(lines))
	for _, line := range lines {
		mov := Movement{}
		dirStr, amountStr := line[0], line[1]
		switch dirStr {
		case "forward":
			mov.Direction = Forward
		case "up":
			mov.Direction = Up
		case "down":
			mov.Direction = Down
		default:
			return nil, ErrUnknownMovement
		}

		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			return nil, err
		}
		mov.Amount = amount

		movements = append(movements, mov)
	}

	return movements, nil
}
