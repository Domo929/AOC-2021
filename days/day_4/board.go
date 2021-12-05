package day_4

type square struct {
	number int
	marked bool
}

type board struct {
	squares [25]square
}

func (b *board) Mark(num int) {
	for sNdx, s := range b.squares {
		if s.number == num {
			b.squares[sNdx].marked = true
		}
	}
}

func (b board) isBingo() bool {
	//rows
	for i := 0; i < 20; i += 5 {
		if isAllMarked(b.squares[i : i+5]...) {
			return true
		}
	}
	//cols
	for i := 0; i < 5; i++ {
		colSqs := make([]square, 0)
		for j := 0; j < 25; j += 5 {
			colSqs = append(colSqs, b.squares[i+j])
		}
		if isAllMarked(colSqs...) {
			return true
		}
	}

	////diag forward slash
	//if isAllMarked(b.squares[20], b.squares[16], b.squares[12], b.squares[8], b.squares[4]) {
	//	return true
	//}
	//
	////diag back slash
	//if isAllMarked(b.squares[0], b.squares[6], b.squares[12], b.squares[18], b.squares[24]) {
	//	return true
	//}

	return false
}

func (b board) sumUnmarked() int {
	var sum int
	for _, s := range b.squares {
		if !s.marked {
			sum += s.number
		}
	}
	return sum
}

func isAllMarked(squares ...square) bool {
	for _, s := range squares {
		if !s.marked {
			return false
		}
	}
	return true
}
