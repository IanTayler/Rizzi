package main

import (
	"bufio"
	"os"
	
	//"unicode"
	//"strconv"
	
	"fmt"
	//"log"
	
	//"./mijn" // Some stuff defined by me.
	
	//"bytes"
	
	// "reflect" // Uncomment when there are tests!
)

// This type is used for syntactic categories.
type Cat int

// The possible categories.
const (
	INTEGER Cat	= iota
	OP
	ID
	EOF					// End of file
	EOL					// End of line
)

// Each token has syntactic information (category) and semantic information (value).
// The value can be anything for now. I should implement a "value" interface.
type Token struct {
	category 	Cat
	value 		interface{}
}

// Reads arithmetic expressions and prints their evaluation.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("calc> ")
	for scanner.Scan() {
		lexer 		:= Lexer{pos: 0}
		lexer.text 	= scanner.Text() + "\n"
		interpreter	:= Interpreter{lexer: lexer}
		result := interpreter.expr()
		fmt.Println(result)
		fmt.Printf("calc> ")
	}
}
