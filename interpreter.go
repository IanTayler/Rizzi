package main

import (
	"./mijn"
	"bytes"
	"strconv"
	"fmt"
	"log"
)

type NodeVisitor interface {
	visit(AST)
	genericVisit(AST)
}

type Interpreter struct {
	parser 		Parser
	globalScope	map[string]int
	setter		int
}

func newInterpreter(p Parser) *Interpreter {
	result := Interpreter{}
	result.parser = p
	result.globalScope = make(map[string]int)
	return &result
}

func (i *Interpreter) visit(n AST) int {
	nType := n.Type()
	var result int
	switch nType {
		case BINOP:		result = mijn.OpToFunc(n.Op())(i.visit(n.Left()), i.visit(n.Right()))
		case UNOP:		result = mijn.UnOpToFunc(n.Op())(i.visit(n.Left()))
		case NUMBER:	result = n.Value()
		case COMPOUND:	
			left := i.visit(n.Left());
			right := i.visit(n.Right());
			if n.Right().Type() != NOOP {
				result = right
			} else {
				result = left
			}
		case NOOP:		
			if VERBOSE {
				fmt.Println("We visited a NoOp and we're returning 0.")
			}
			result = 0
		case ASS:
			varName := n.Left().Op()
			i.globalScope[varName] = i.visit(n.Right())
			result = i.globalScope[varName]
		case VAR:
			varName 	 := n.Op()
			varValue, ok := i.globalScope[varName]
			if ok {
				result = varValue
			} else {
				log.Println(varName, " is an undeclared variable in a denotative context. Returning 0.")
				result = 0
			}
		case CONTROL:
			if n.Op() == "for" {
				for i.visit(n.Left()) != 0 {
					result = i.visit(n.Right())
				}
			} else if n.Op() == "if" {
				if i.visit(n.Left()) != 0 {
					result = i.visit(n.Right())
				}
			}
		case ARGUMENT:
			argName := n.Op()
			i.globalScope[argName] = i.setter
	}
	return result
}

func (i *Interpreter) interpret() int {
	tree := i.parser.parse()
	return i.visit(tree)
}

//------------------------------------------------------------
//-------------- PREFIX AND POSTFIX TRANSLATION --------------
//------------------------------------------------------------
func (i *Interpreter) translatePostfix() string {
	tree := i.parser.parse()
	return i.visitPostfix(tree)
}

func (i *Interpreter) visitPostfix(n AST) string {
	nType := n.Type()
	var buffer bytes.Buffer
	switch nType {
		case BINOP:	
			buffer.WriteString(i.visitPostfix(n.Left()) + " ")
			buffer.WriteString(i.visitPostfix(n.Right()) + " ")
			buffer.WriteString(n.Op() + " ")
		case UNOP:
			buffer.WriteString(i.visitPostfix(n.Left()) + " ")
			buffer.WriteString(n.Op() + " ")
		case NUMBER:
			buffer.WriteString(strconv.Itoa(n.Value()) + " ")
	}
	return buffer.String()
}

func (i *Interpreter) translateLisp() string {
	tree := i.parser.parse()
	return i.visitLisp(tree)
}

func (i *Interpreter) visitLisp(n AST) string {
	nType := n.Type()
	var buffer bytes.Buffer
	switch nType {
		case BINOP:
			buffer.WriteString("(" + n.Op() + " ")
			buffer.WriteString(i.visitLisp(n.Left()) + " ")
			buffer.WriteString(i.visitLisp(n.Right()))
			buffer.WriteString(")")
		case UNOP:
			buffer.WriteString("(" + n.Op() + " ")
			buffer.WriteString(i.visitPostfix(n.Left()))
			buffer.WriteString(")")
		case NUMBER:
			buffer.WriteString(strconv.Itoa(n.Value()))
	}
	return buffer.String()
}
