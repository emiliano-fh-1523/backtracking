package main

import "fmt"

func can_parse(input string) bool {
	input = input
	return false
}

func main() {
	inputs := []string{
		"a+a*a+(a*a)",
		"a++a",
		"a",
		"a+a",
		"a(a)",
		"(a+a)+a",
	}

	for _, input := range inputs {
		fmt.Printf("Parsing input: %s -- %t\n", input, can_parse(input))
	}
}
