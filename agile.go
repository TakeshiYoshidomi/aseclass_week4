package main

type Board struct {
	tokens []int
}

func (b *Board) put(x, y int, u string) {
       if u == "o" {
             b.tokens[x + 3*y] = 1
       }
}
func (b *Board) get(x, y int) string {
        if b.tokens[x + 3*y] == 1 {
             reture "o"
        }
}
