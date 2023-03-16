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

type MetaCommandResult bool

const (
	META_COMMAND_SUCCESS              MetaCommandResult = true
	META_COMMAND_UNRECOGNIZED_COMMAND MetaCommandResult = true
)

type PrepareResult bool

const (
	PREPARE_SUCCESS                PrepareResult = true
	PREPARE_UNRECOGNIZED_STATEMENT PrepareResult = true
)

type StatementType string

const (
	STATEMENT_INSERT StatementType = "INSERT"
	STATEMENT_SELECT StatementType = "SELECT"
)

type Statement struct {
	StatementType StatementType
}

type ExitStatus bool

const (
	EXIT_SUCCESS ExitStatus = true
)

// Todo: continue will skip prepare_statement but the db > won't be triggered, redo the loop of print db> but now now.

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to SQLITE in GO CLI v0.0.01")
	fmt.Print("db > ")

	for scanner.Scan() {
		text := scanner.Text()
		var prompt inputBuffer
		var statement Statement

		prompt = inputBuffer{input: text, inputLen: len(text)}
		if strings.Compare(prompt.input, ".exit") == 0 {
			exit(EXIT_SUCCESS)
		} else {
			if strings.Compare(prompt.input, ".") == 0 {
				fmt.Println(prompt)
				switch do_meta_command(&prompt) {
				case (META_COMMAND_SUCCESS):
					fmt.Println("META COMMAND SUCCESS", prompt)
				case (META_COMMAND_UNRECOGNIZED_COMMAND):
					fmt.Printf("Unrecognized command '%s'\n", prompt.input)
				}
			}
		}

		switch prepare_statement(&prompt, &statement) {
		case (PREPARE_SUCCESS):
		case (PREPARE_UNRECOGNIZED_STATEMENT):
			fmt.Printf("Unrecognized command '%s'\n", prompt.input)
		}

		execute_statement(&statement)
		fmt.Println("Executed")

		fmt.Printf("db > ")

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func do_meta_command(prompt *inputBuffer) MetaCommandResult {
	if strings.Compare(prompt.input, ".exit") == 0 {
		exit(EXIT_SUCCESS)
	}
	return META_COMMAND_UNRECOGNIZED_COMMAND

}

func prepare_statement(prompt *inputBuffer, statement *Statement) PrepareResult {

	fmt.Println(*prompt, *statement)
	if strings.Compare(prompt.input, "insert") == 0 {
		statement.StatementType = STATEMENT_INSERT
		return PREPARE_SUCCESS
	}

	if strings.Compare(prompt.input, "select") == 0 {
		statement.StatementType = STATEMENT_SELECT
		return PREPARE_SUCCESS
	}

	return PREPARE_UNRECOGNIZED_STATEMENT

}

func execute_statement(Statement *Statement) {

	switch Statement.StatementType {

	case STATEMENT_INSERT:
		fmt.Println("This is where we would do an INSERT")
	case STATEMENT_SELECT:
		fmt.Println("This is where we would do an SELECT")

	}
}

func exit(command ExitStatus) {

	if command == EXIT_SUCCESS {
		os.Exit(1)
	}
}
