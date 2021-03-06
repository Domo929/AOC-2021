package main

import (
	"fmt"

	"github.com/Domo929/AOC-2021/days/day_1"
	"github.com/Domo929/AOC-2021/days/day_2"
	"github.com/Domo929/AOC-2021/days/day_3"
	"github.com/Domo929/AOC-2021/days/day_4"
	"github.com/Domo929/AOC-2021/days/day_5"
	"github.com/Domo929/AOC-2021/days/day_6"
	"github.com/Domo929/AOC-2021/days/day_7"
	"github.com/Domo929/AOC-2021/days/day_8"
)

var (
	solvers = []func(string) error{
		func(s string) error { return nil },
		day_1.Solve,
		day_2.Solve,
		day_3.Solve,
		day_4.Solve,
		day_5.Solve,
		day_6.Solve,
		day_7.Solve,
		day_8.Solve,
	}

	day   = 8
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
