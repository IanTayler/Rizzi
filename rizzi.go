package main

/*******************************************
 * 
 * TODO:
 * a. Finish implementing parser.go, as in part 9.
 * ai. Parser.exprSet()
 * aii.  Parser.print()
 * 
 * ****************************************/

import (
	"bufio"
	"os"
	"io/ioutil"
	
	//"unicode"
	"strconv"
	
	"fmt"
	"log"
	
	//"./mijn" // Some stuff defined by me.
	
	//"bytes"
	
	// "reflect" // Uncomment when there are tests!
)

// SET TO true TO GET A VERBOSE OUTPUT THAT EXPLAINS WHAT THE PROGRAM IS DOING.
// USEFUL FOR DEBUGGING
const VERBOSE = false

// This type is used for syntactic categories.
type Cat int

// The possible categories.
const (
	INTEGER Cat	= iota	// 0
	OP					// 1
	ID					// 2
	LPAR				// 3
	RPAR				// 4
	MAIN				// 5
	END					// 6
	PRINT				// 7
	IF					// 8
	THEN				// 9
	FOR					// 10
	DO					// 11
	ASSIGN				// 12
	COMMA				// 13
	LBRACKET			// 14
	RBRACKET			// 15
	EOF					// 16 -- End of file
	EOL					// 17 -- End of line
)

// Each token has syntactic information (category) and semantic information (value).
// The value can be anything for now. I should implement a "value" interface.
type Token struct {
	category 	Cat
	value 		interface{}
}

// Reads arithmetic expressions and prints their evaluation.
func main() {
	if len(os.Args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("rizzi> ")
		for scanner.Scan() {
			lexer 		:= newLexer(scanner.Text() + "\x00")
			parser		:= newParser(*lexer)
			interpreter	:= newInterpreter(*parser)
		// TESTS
/*			if VERBOSE {
				thisToken 	:= lexer.getNextToken()
				for thisToken.category != EOF {
					fmt.Println(thisToken)
					thisToken = lexer.getNextToken()
				}
			}
*/		// END TESTS
			result 		:= interpreter.interpret()
			fmt.Println(result)
			fmt.Println(interpreter.globalScope)
			fmt.Printf("rizzi> ")
		}
	} else {
		text, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		lexer := newLexer(string(text) + "\x00")
		parser := newParser(*lexer)
		interpreter := newInterpreter(*parser)
		var setter int
		setted := false
		for i := 2; i < len(os.Args) - 1; i++ {
			if os.Args[i] == "-a" {
				setter, err = strconv.Atoi(os.Args[i+1])
				setted = true
			}
		}
		
		if setted == false {
			fmt.Println("You didn't provide an argument. Provide one now.")
			fmt.Printf("arg> ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			setter, err = strconv.Atoi(scanner.Text())
		}
		
		interpreter.setter = setter
		result := interpreter.interpret()
		fmt.Println(result)
		fmt.Println(interpreter.globalScope)
	}
}
