package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	slices := []int{0}
	pointer := 0
	in := []int{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '+':
			slices[pointer]++
		case '-':
			slices[pointer]--
		case '>':
			pointer++
			if pointer == len(slices) {
				slices = append(slices, 0)
			}
		case '<':
			pointer--
		case '[':
			in = append(in, i)
		case ']':
			if slices[pointer] == 0 {
				in = in[:len(in)-1]
			} else {
				i = in[len(in)-1]
			}

		case '.':
			fmt.Print(string(slices[pointer]))
		}
	}
}
