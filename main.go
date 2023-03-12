package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type inputBuffer struct {
	input    string
	inputLen int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to SQLITE in GO CLI v0.0.01")
	fmt.Print("db > ")
	for scanner.Scan() {
		text := scanner.Text()
		prompt := inputBuffer{input: text, inputLen: len(text)}
		if strings.Compare(prompt.input, "exit") == 0 {
			os.Exit(1)
		} else {
			fmt.Printf("Unrecognized command '%s'\n", prompt.input)
		}

		fmt.Print("db >")

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
