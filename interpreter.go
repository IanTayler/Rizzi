package main

import (
	"./mijn"
	"fmt"
	"log"
)

// A struct defining an interpreter.
// Perhaps too OOP-ish.
type Interpreter struct {
	currentToken 	Token
	lexer			Lexer
}

// Interpreter method that checks whether the current token is of a certain syntactic category and, if it is, goes to the next token.
func (i *Interpreter) eat(catg Cat) {
	if i.currentToken.category ==  catg {
		fmt.Println("Successfully eaten ", catg)
		if catg != EOF {
			i.currentToken = i.lexer.getNextToken()
		}
	} else {
		// Handle errors here in the future
		log.Println("Things went wrong with eat, I expected a ", catg, "and got a ", i.currentToken.category)
	}
}

// Interpreter method that evaluates a sub-factor (that is, exponential expressions).
func (i *Interpreter) subfactor() int {
	currToken := i.currentToken
	i.eat(INTEGER)
	return currToken.value.(int)
}

// Interpreter method that evaluates a factor
// Note the wrong precedence of the operator exp
func (i *Interpreter) factor() int {
	result := i.subfactor()
	
	for i.currentToken.category == OP && mijn.IsSubfactorOp(i.currentToken.value.(string)) {
		op := i.currentToken
		i.eat(OP)
		
		right := i.currentToken
		i.eat(INTEGER)
		
		result = mijn.OpToFunc(op.value.(string))(result, right.value.(int))
	}
	return result
}

// Interpreter method that evaluates a term
func (i *Interpreter) term() int {
	result := i.factor()
	for i.currentToken.category == OP && mijn.IsFactorOp(i.currentToken.value.(string)) {
		op := i.currentToken
		i.eat(OP)
		
		result = mijn.OpToFunc(op.value.(string))(result, i.factor())
	}
	
	return result
}

// Interpreter method that evaluates an expression.
func (i *Interpreter) expr() int {
	i.currentToken = i.lexer.getNextToken()
	var result int
	
	switch i.currentToken.category {
		case INTEGER:
			result = i.term()
			
			for i.currentToken.category == OP {
				
				op := i.currentToken
				i.eat(OP)
				
				result = mijn.OpToFunc(op.value.(string))(result, i.term())
			}
			
			i.eat(EOF)
			
			return result
		
		case OP:
			op := i.currentToken
			i.eat(OP)
			
			number := i.currentToken
			i.eat(INTEGER)
			
			i.eat(EOF)
			
			return mijn.UnOpToFunc(op.value.(string))(number.value.(int))
			
		default:
			log.Println("First token should be an INTEGER or an OP. Returning 0.")
			return 0
			
	}
}
