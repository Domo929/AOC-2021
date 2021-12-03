package main

import (
	"github.com/Domo929/AOC-2021/days/day_1"
	"github.com/Domo929/AOC-2021/days/day_2"
	"github.com/Domo929/AOC-2021/days/day_3"
)

var (
	files = []string{
		"inputs/day_1_ex.tsv", //0
		"inputs/day_1_act.tsv",
		"inputs/day_2_ex.tsv",
		"inputs/day_2_act.tsv",
		"inputs/day_3_ex.tsv",
		"inputs/day_3_act.tsv", //5
	}

	solvers = []func(string) error{
		func(s string) error { return nil },
		day_1.Solve,
		day_2.Solve,
		day_3.Solve,
	}

	day = 3
)

func main() {
	err := solvers[day](files[5])
	if err != nil {
		panic(err)
	}
}
