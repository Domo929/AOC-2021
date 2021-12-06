package day_5

import (
	"bufio"
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

	lines, err := parse(f)
	if err != nil {
		return err
	}

	err = part1(lines)
	if err != nil {
		return err
	}

	return nil
}

func part1(lines []line) error {
	part1Intersections := make(map[string]int)
	part2Intersections := make(map[string]int)
	for _, l := range lines {
		if l.isFlat() {
			pts := l.pointsVorH()
			for _, pt := range pts {
				key := fmt.Sprintf("%d|%d", pt.X, pt.Y)
				part1Intersections[key]++
				part2Intersections[key]++
			}
		}
		if l.isDiag() {
			pts := l.pointsDiag()
			for _, pt := range pts {
				key := fmt.Sprintf("%d|%d", pt.X, pt.Y)
				part2Intersections[key]++
			}
		}
	}

	var part1Count = 0
	for _, i := range part1Intersections {
		if i > 1 {
			part1Count++
		}
	}

	var part2Count = 0
	for _, i := range part2Intersections {
		if i > 1 {
			part2Count++
		}
	}

	fmt.Printf("number of part1Intersections with greater than two lines: %d\n", part1Count)
	fmt.Printf("number of part2Intersections with greater than two lines: %d\n", part2Count)

	return nil
}

func parse(r io.Reader) ([]line, error) {
	scanner := bufio.NewScanner(r)
	lines := make([]line, 0)

	for scanner.Scan() {
		text := scanner.Text()

		line, err := parseLine(text)
		if err != nil {
			return nil, err
		}

		lines = append(lines, line)
	}

	return lines, nil
}

func parseLine(text string) (line, error) {
	parts := strings.Split(text, " -> ")
	a, err := parsePoint(parts[0])
	if err != nil {
		return line{}, err
	}
	b, err := parsePoint(parts[1])
	if err != nil {
		return line{}, err
	}
	return line{
		A: a,
		B: b,
	}, nil
}

func parsePoint(text string) (point, error) {
	parts := strings.Split(text, ",")
	x, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return point{}, err
	}
	y, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		return point{}, err
	}
	return point{
		X: int(x),
		Y: int(y),
	}, nil
}
