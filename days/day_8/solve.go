package day_8

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	Digits  []string
	Display []string
}

func Solve(pathToInput string) error {
	f, err := os.Open(pathToInput)
	if err != nil {
		return err
	}
	defer f.Close()

	entries, err := parse(f)
	if err != nil {
		return err
	}

	err = part1(entries)
	if err != nil {
		return err
	}

	err = part2(entries)
	if err != nil {
		return err
	}

	return nil
}

func part1(entries []Entry) error {
	counts := make(map[int]int)
	for _, entry := range entries {
		for _, s := range entry.Display {
			counts[decodeUnique(s)]++
		}
	}

	num1 := counts[1]
	num4 := counts[4]
	num7 := counts[7]
	num8 := counts[8]

	fmt.Printf("One: %d | Four: %d | Seven: %d | Eight: %d | Total: %d\n", num1, num4, num7, num8, num1+num4+num7+num8)
	return nil
}

func part2(entries []Entry) error {
	var sum int64
	for _, entry := range entries {
		ref := decodeDigits(entry.Digits)
		var numStr string
		for _, s := range entry.Display {
			n, err := match(ref, s)
			if err != nil {
				return err
			}
			numStr += fmt.Sprint(n)
		}
		num, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			return err
		}
		sum += num
	}

	fmt.Println("total sum:", sum)
	return nil
}

func parse(r io.Reader) ([]Entry, error) {
	entries := make([]Entry, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()

		parts := strings.Split(text, " | ")

		entries = append(entries, Entry{
			Digits:  strings.Split(parts[0], " "),
			Display: strings.Split(parts[1], " "),
		})
	}

	return entries, nil
}

func decodeUnique(digit string) int {
	switch len(digit) {
	case 2:
		return 1
	case 4:
		return 4
	case 3:
		return 7
	case 7:
		return 8
	default:
		return -1
	}
}

func decodeDigits(digits []string) map[string]int {
	ref := make(map[string]int)
	lens := make(map[int][]string)
	inv := make(map[int]string)
	for _, digit := range digits {
		val := decodeUnique(digit)
		ref[digit] = val
		if val != -1 {
			inv[val] = digit
		}

		l := len(digit)
		if lens[l] == nil {
			lens[l] = make([]string, 0)
		}
		lens[l] = append(lens[l], digit)
	}

	for _, digit := range lens[6] {
		if containsAll(digit, inv[4]) {
			ref[digit] = 9
			continue
		}
		if containsAll(digit, inv[1]) {
			ref[digit] = 0
			continue
		}
		ref[digit] = 6
	}

	for _, digit := range lens[5] {
		if containsAll(digit, inv[1]) {
			ref[digit] = 3
			continue
		}
		if containsAll(digit, extractDiff(inv[4], inv[1])) {
			ref[digit] = 5
			continue
		}
		ref[digit] = 2
	}

	return ref
}

func containsAll(root, ref string) bool {
	for _, r := range ref {
		if !strings.ContainsRune(root, r) {
			return false
		}
	}
	return true
}

func extractDiff(four, one string) string {
	var diff string
	for _, fr := range four {
		if !strings.ContainsRune(one, fr) {
			diff += string(fr)
		}
	}
	return diff
}

func match(ref map[string]int, s string) (int, error) {
	for digit, i := range ref {
		if len(digit) == len(s) && containsAll(s, digit) {
			return i, nil
		}
	}
	return 0, errors.New("unmatched")
}
