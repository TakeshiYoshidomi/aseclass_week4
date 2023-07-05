package main

type Board struct {
	tokens []int
}

func (b *Board) put(x, y int, u string) {
	if b.tokens[x+3*y] != 0 {
		return
	} else if u == "o" {
		b.tokens[x+3*y] = 1
	} else {
		b.tokens[x+3*y] = -1
	}

}
func (b *Board) get(x, y int) string {
	if b.tokens[x+3*y] == 1 {
		return "o"
	}
	if b.tokens[x+3*y] == -1 {
		return "x"
	}
	return "other"
}

func (b *Board) judge() string {
	i := 0
	for i <= 2 {
		sum_w := 0
		sum_w = b.tokens[i*3] + b.tokens[i*3+1] + b.tokens[i*3+2]
		if sum_w == 3 {
			return "o win"
		} else if sum_w == -3 {
			return "x win"
		}
		sum_h := 0
		sum_h = b.tokens[i] + b.tokens[i+3] + b.tokens[i+6]
		if sum_h == 3 {
			return "o win"
		} else if sum_h == -3 {
			return "x win"
		}
		i++
	}
	sum_d1 := 0
	sum_d1 = b.tokens[0] + b.tokens[4] + b.tokens[8]
	if sum_d1 == 3 {
		return "o win"
	} else if sum_d1 == -3 {
		return "x win"
	}
	sum_d2 := 0
	sum_d2 = b.tokens[2] + b.tokens[4] + b.tokens[6]
	if sum_d2 == 3 {
		return "o win"
	} else if sum_d2 == -3 {
		return "x win"
	}
	return "other"
}

func main() {
}
