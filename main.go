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
				in = delete(in)
				continue
			}
			i = in[len(in)-1]

		case '.':
			fmt.Print(string(slices[pointer]))
		}
	}
}

func delete(a []int) []int {
	if len(a) <= 1 {
		return nil
	}
	c := []int{}
	for i := 0; i < len(a)-1; i++ {
		c = append(c, a[i])
	}
	return c
}
