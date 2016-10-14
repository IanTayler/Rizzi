package main

import (
	"bufio"
	"os"
	
	"unicode"
	"strconv"
	
	"fmt"
	"log"
	
	"./mijn" // Some stuff defined by me.
	
	"bytes"
	
	// "reflect" // Uncomment when there are tests!
)

// This type is used for syntactic categories.
type Cat int

// The possible categories.
const (
	INTEGER Cat	= iota
	OP
	EOF					// End of file
	EOL					// End of line
)

// Each token has syntactic information (category) and semantic information (value).
// The value can be anything for now. I should implement a "value" interface.
type token struct {
	category 	Cat
	value 		interface{}
}

// A struct defining an interpreter.
// Perhaps too OOP-ish.
type interpreter struct {
	text 			string
	pos 			int
	currentToken 	token
}

// Interpreter method that gets the current char, skipping all whitespace.
func (i *interpreter) getCurrChar() byte {
	for mijn.IsBlank(i.text[i.pos]) {
		i.pos += 1
	}
	return i.text[i.pos]
}

// Interpreter method that gets a string containing an integer starting in the current char.
func (i *interpreter) getIntegerStr() string {
	var buffer bytes.Buffer
		
	for unicode.IsDigit(rune(i.getCurrChar())) {
		
		buffer.WriteString(string(i.getCurrChar()))
		
		if i.pos < len(i.text) - 1 {
			i.pos += 1
		} else {
			break
		}
	}
	
	return buffer.String()
}

// Interpreter method that gets an operator starting in the current char.
func (i *interpreter) getOperator() string {
	var buffer bytes.Buffer
	
	for unicode.IsLetter(rune(i.getCurrChar())) {
		
		buffer.WriteString(string(i.getCurrChar()))
		
		if i.pos < len(i.text) -1 {
			i.pos += 1
		} else {
			break
		}
	}
	
	if buffer.String() != "" {
		return buffer.String()
	} else {
		opChar := i.getCurrChar()
		i.pos += 1
		return string(opChar)
	}
}

// Interpreter method that gets the next token starting in its current position.
func (i *interpreter) getNextToken() token {
	
	if i.pos > len(i.text) - 1 {
		return token{EOF, ""}
	}
	
	if unicode.IsDigit(rune(i.getCurrChar())) {
		representedNumber, _ := strconv.Atoi(i.getIntegerStr())
		mToken := token{INTEGER, representedNumber}
		return mToken
		
	} else if mijn.IsOp(i.getCurrChar()) {
		opr := i.getOperator()
		
		// fmt.Println("This is what getOperator is giving: ", opr) // test
		
		mToken := token{OP, opr}
		return mToken
		
	}
	
	// Handle errors here in the future
	log.Println("Things went wrong with getNextToken")
	
	return token{}
}

// Interpreter method that checks whether the current token is of a certain syntactic category and, if it is, goes to the next token.
func (i *interpreter) eat(catg Cat) {
	if i.currentToken.category ==  catg {
		i.currentToken = i.getNextToken()
	} else {
		// Handle errors here in the future
		log.Println("Things went wrong with eat, I expected a ", catg, "and got a ", i.currentToken.category)
	}
}

// Interpreter method that evaluates an expression.
func (i *interpreter) expr() int {
	i.currentToken = i.getNextToken()
	
	switch i.currentToken.category {
		case INTEGER:
			left := i.currentToken
			i.eat(INTEGER)
			
			op := i.currentToken
			// fmt.Println("This is the operator token: ", op)	//test
			i.eat(OP)
			
			// fmt.Println("All's fine") //test
	
			right := i.currentToken
			i.eat(INTEGER)
			
			// fmt.Println("This is the left token's value: ", left.value) //test
			// fmt.Println("This is the left token's value's type: ", reflect.TypeOf(left.value)) //test
			
			// fmt.Println("This is the operator token's value: ", op.value) //test
			// fmt.Println("This is the operator token's value's type: ", reflect.TypeOf(op.value)) //test
			
			// fmt.Println("This is the right token's value: ", right.value) //test
			// fmt.Println("This is the right token's value's type: ", reflect.TypeOf(right.value)) //test
	
			return mijn.OpToFunc(op.value.(string))(left.value.(int), right.value.(int))
		
		case OP:
			op := i.currentToken
			i.eat(OP)
			
			number := i.currentToken
			i.eat(INTEGER)
			
			return mijn.UnOpToFunc(op.value.(string))(number.value.(int))
			
		default:
			log.Println("First token should be an INTEGER or an OP. Returning 0.")
			return 0
			
	}
}

// Reads arithmetic expressions and prints their evaluation.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("calc> ")
	for scanner.Scan() {
		interp := interpreter{text: scanner.Text(), pos: 0}
		result := interp.expr()
		fmt.Println(result)
		fmt.Printf("calc> ")
	}
}
