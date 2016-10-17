package main

import (
	"./mijn"
	"fmt"
	"log"
)
type TreeType int

const (
	BINOP	 TreeType = iota
	NUMBER
	
)
// A struct defining an interpreter.
// Perhaps too OOP-ish.
type Parser struct {
	currentToken 	Token
	lexer			Lexer
}

// Some types to use for trees:
type AST interface {
	Type()	TreeType
}

type BinOp struct {
	left 	AST
	token 	Token
	op		string
	right	AST
}

func newBinOp(l AST, tok Token, r AST) BinOp {
	result := 		BinOp{}
	result.left 	= l
	result.token	= tok
	result.op		= tok.value.(string)
	result.right	= r
	return result
}

func (bo *BinOp) Type() TreeType {
	return BINOP
}

type Number struct {
	token	Token
	value	int
}

func newInteger(tok Token) Number {
	result := Number{}
	result.token = tok
	result.value = tok.value.(int)
	return result
}

func (nu *Number) Type() TreeType {
	return NUMBER
}

// A function to create a Parser nicely
func newParser(l Lexer) Parser {
	result := Parser{}
	result.lexer = l
	result.currentToken = result.lexer.getNextToken()
	return result
}

// Parser method that checks whether the current token is of a certain syntactic category and, if it is, goes to the next token.
func (p *Parser) eat(catg Cat) {
	if p.currentToken.category ==  catg {
		fmt.Println("Successfully eaten ", catg)
		if catg != EOF {
			p.currentToken = p.lexer.getNextToken()
		}
	} else {
		// Handle errors here in the future
		log.Println("Things went wrong with eat, I expected a ", catg, "and got a ", p.currentToken.category)
	}
}

// Parser method that evaluates a sub-factor (that is, exponential expressions).
func (p *Parser) subfactor() int {
	currToken := p.currentToken
	p.eat(INTEGER)
	return currToken.value.(int)
}

// Parser method that evaluates a factor
// Note the wrong precedence of the operator exp
func (p *Parser) factor() int {
	var result int
	switch p.currentToken.category {
		case INTEGER:
			result = p.subfactor()
			
			for p.currentToken.category == OP && mijn.IsSubfactorOp(p.currentToken.value.(string)) {
			op := p.currentToken
			p.eat(OP)
			
			right := p.currentToken
			p.eat(INTEGER)
			
			result = mijn.OpToFunc(op.value.(string))(result, right.value.(int))
			}
		case LPAR:
			p.eat(LPAR)
			
			result = p.expr()
			
			p.eat(RPAR)
		default:
			log.Println("Syntactic error. Expected LPAR or INTEGER. Returning 0")
	}
	return result
}

// Parser method that evaluates a term
func (p *Parser) term() int {
	var result int
	switch p.currentToken.category {
		case INTEGER:
			result = p.factor()
			for p.currentToken.category == OP && mijn.IsFactorOp(p.currentToken.value.(string)) {
				op := p.currentToken
				p.eat(OP)
				
				result = mijn.OpToFunc(op.value.(string))(result, p.factor())
			}
		case LPAR:
			result = p.factor()
			
		default:
			log.Println("Syntactic error. Expected LPAR or INTEGER. Returning 0")
	}
	return result
}

// Parser method that evaluates an expression.
func (p *Parser) expr() int {
	var result int
	
	switch p.currentToken.category {
		case INTEGER:
			result = p.term()
			
			for p.currentToken.category == OP {
				
				op := p.currentToken
				p.eat(OP)
				
				result = mijn.OpToFunc(op.value.(string))(result, p.term())
			}
		
		case OP:
			op := p.currentToken
			p.eat(OP)
			
			number := p.currentToken
			p.eat(INTEGER)
			
			result = mijn.UnOpToFunc(op.value.(string))(number.value.(int))
			
		case LPAR:
			result = p.factor()
		
		default:
			log.Println("First token should be an INTEGER, an OP or a LPAR. Returning 0.")
			
	}
	return result
}

// Parser method to start evaluating an expression.
func (p *Parser) parse() int {	
	return p.expr()
}
