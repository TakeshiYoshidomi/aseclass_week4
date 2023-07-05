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
		sum := 0
		sum = b.tokens[i*3] + b.tokens[i*3+1] + b.tokens[i*3+2]
		if sum == 3 {
			return "o win"
		} else if sum == -3 {
			return "x win"
		}
		i++
	}
	return "other"
}

func main() {
}
