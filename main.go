package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: <file.bf>")
		os.Exit(1)
	}

	code, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if err := compile(code); err != nil {
		fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
		os.Exit(1)
	}
}

func compile(code []byte) error {
	arr := make([]byte, 30000)
	ptr := 0
	brs := mapBrackets(string(code))
	for pc := 0; pc < len(code); pc++ {
		switch code[pc] {
		case '>': ptr++
		case '<': ptr--
		case '+': arr[ptr]++
		case '-': arr[ptr]--
		case '.': fmt.Printf("%c",arr[ptr])
		case ',': fmt.Scanf("%c",&arr[ptr])
		case '[':
			if arr[ptr] == 0 {
				pc = brs[pc]
			}
		case ']':
			if arr[ptr] != 0 {
				pc = brs[pc]
			}
		case ' ', '\t', '\n', '\r':
			continue
		default:
			return fmt.Errorf("%c is not exist.", code[pc])
		}
	}
	return nil
}

func mapBrackets(program string) map[int]int {
	stack := []int{}
	pairs := make(map[int]int)

	for i, ch := range program {
		if ch == '[' {
			stack = append(stack, i)
		} else if ch == ']' {
			open := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pairs[open] = i
			pairs[i] = open
		}
	}
	return pairs
}