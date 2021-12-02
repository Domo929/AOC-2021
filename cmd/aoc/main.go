package main

import (
	"github.com/Domo929/AOC-2021/days/day_1"
	"github.com/Domo929/AOC-2021/days/day_2"
)

var (
	files = []string{
		"inputs/day_1_ex.tsv",
		"inputs/day_1_act.tsv",
		"inputs/day_2_ex.tsv",
		"inputs/day_2_act.tsv",
	}

	solvers = []func(string) error{
		func(s string) error { return nil },
		day_1.Solve,
		day_2.Solve,
	}

	day = 2
)

func main() {
	err := solvers[day](files[3])
	if err != nil {
		panic(err)
	}
}
