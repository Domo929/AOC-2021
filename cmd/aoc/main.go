package main

import (
	"fmt"

	"github.com/Domo929/AOC-2021/days/day_1"
	"github.com/Domo929/AOC-2021/days/day_2"
	"github.com/Domo929/AOC-2021/days/day_3"
	"github.com/Domo929/AOC-2021/days/day_4"
	"github.com/Domo929/AOC-2021/days/day_5"
)

var (
	solvers = []func(string) error{
		func(s string) error { return nil },
		day_1.Solve,
		day_2.Solve,
		day_3.Solve,
		day_4.Solve,
		day_5.Solve,
	}

	day   = 5
	isAct = true
)

func main() {
	suffix := "_ex.tsv"
	if isAct {
		suffix = "_act.tsv"
	}
	filename := fmt.Sprintf("inputs/day_%d%s", day, suffix)
	err := solvers[day](filename)
	if err != nil {
		panic(err)
	}
}
