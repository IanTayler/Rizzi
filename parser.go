package main

import (
	"./mijn"
	"fmt"
	"log"
)
type TreeType int

const (
	BINOP	 TreeType = iota
	UNOP
	NUMBER
	NOOP
	COMPOUND
	VAR
	ASS
	CONTROL
	ARGUMENT
	
)
// A struct defining a Parser
// Perhaps too OOP-ish.
type Parser struct {
	currentToken 	Token
	lexer			Lexer
}

// Some types to use for trees:
type AST interface {
	Type()	TreeType
	Op()	string
	Left()	AST
	Right()	AST
	Value()	int
}

//-------------- BinOp struct --------------
type BinOp struct {
	left 	AST
	token 	Token
	op		string
	right	AST
}

func newBinOp(l AST, tok Token, r AST) *BinOp {
	result := BinOp{}
	result.left 	= l
	result.token	= tok
	result.op		= tok.value.(string)
	result.right	= r
	return &result
}

func (bo *BinOp) Type() TreeType {
	return BINOP
}

func (bo *BinOp) Op() string {
	return bo.op
}

func (bo *BinOp) Left() AST {
	return bo.left
}

func (bo *BinOp) Right() AST {
	return bo.right
}

func (bo *BinOp) Value() int {
	log.Println("Warning: trying to acess a BinOp's non-existing value. Returning 0")
	return 0
}

//--------------	UnOp struct --------------
type UnOp struct {
	child	AST
	token 	Token
	op		string
}

func newUnOp(ch AST, tok Token) *UnOp {
	result := UnOp{}
	result.child 	= ch
	result.token	= tok
	result.op		= tok.value.(string)
	return &result
}

func (uo *UnOp) Type() TreeType {
	return UNOP
}

func (uo *UnOp) Op() string {
	return uo.op
}

func (uo *UnOp) Left() AST {
	return uo.child
}

func (uo *UnOp) Right() AST {
	log.Println("Warning: tried to get an UnOp's right child. Returning its only child. Use Left() instead.")
	return uo.child
}

func (uo *UnOp) Value() int {
	log.Println("Warning: trying to acess an UnOp's non-existing value. Returning 0")
	return 0
}

//-------------- Number struct --------------
type Number struct {
	token	Token
	value	int
}

func newNumber(tok Token) *Number {
	result := Number{}
	result.token = tok
	result.value = tok.value.(int)
	return &result
}

func (nu *Number) Type() TreeType {
	return NUMBER
}

func (nu *Number) Op() string {
	log.Println("Error: trying to get a Number's operator. Returning empty string")
	return ""
}

func (nu *Number) Left() AST {
	log.Println("Error: trying to access a Number's child. Returning itself.")
	return nu
}

func (nu *Number) Right() AST {
	log.Println("Error: trying to access a Number's child. Returning itself.")
	return nu
}

func (nu *Number) Value() int {
	return nu.value
}
//-------------- NoOp struct --------------
type NoOp struct {
	
}
func newNoOp() *NoOp {
	return &NoOp{}
}

func (no *NoOp) Type() TreeType {
	return NOOP
}

func (no *NoOp) Op() string {
	log.Println("Error: trying to get a NoOp's operator. Returning empty string")
	return ""
}

func (no *NoOp) Left() AST {
	log.Println("Error: trying to access a NoOp's child. Returning itself.")
	return no
}

func (no *NoOp) Right() AST {
	log.Println("Error: trying to access a NoOp's child. Returning itself.")
	return no
}

func (no *NoOp) Value() int {
	log.Println("Error: trying to access a NoOp's value. Returning 0.")
	return 0
}
//-------------- Compound struct --------------
type Compound struct {
	left 	AST
	right 	AST
}

func newCompound(l AST, r AST) *Compound {
	result := Compound{}
	result.left 	= l
	result.right	= r
	return &result
}

func (co *Compound) Type() TreeType {
	return COMPOUND
}

func (co *Compound) Op() string {
	return "compound"
}

func (co *Compound) Left() AST {
	return co.left
}

func (co *Compound) Right() AST {
	return co.right
}

func (co *Compound) Value() int {
	log.Println("Warning: trying to acess a Compound's non-existing value. Returning 0")
	return 0
}

//-------------- Assign struct --------------
type Assign struct {
	left 	AST
	right	AST
	op		string
	token	Token
}

func newAssign(l AST, tok Token, r AST) *Assign {
	result := Assign{}
	result.left 	= l
	result.token	= tok
	result.op		= tok.value.(string)
	result.right	= r
	return &result
}

func (as *Assign) Type() TreeType {
	return ASS
}

func (as *Assign) Op() string {
	return as.op
}

func (as *Assign) Left() AST {
	return as.left
}

func (as *Assign) Right() AST {
	return as.right
}

func (as *Assign) Value() int {
	log.Println("Warning: trying to acess an Assign's non-existing value. Returning 0")
	return 0
}

//-------------- Var struct --------------
type Var struct {
	token Token
	value string
}

func newVar(tok Token) *Var {
	result := Var{}
	result.token = tok
	result.value = tok.value.(string)
	return &result
}

func (v *Var) Type() TreeType {
	return VAR
}

func (v *Var) Op() string {
	if VERBOSE { fmt.Println("A Var's Op() is its identifier.")}
	return v.value
}

func (v *Var) Left() AST {
	log.Println("Error: trying to access a Var's child. Returning itself.")
	return v
}

func (v *Var) Right() AST {
	log.Println("Error: trying to access a Var's child. Returning itself.")
	return v
}

func (v *Var) Value() int {
	log.Println("Error: trying to acess a Var's value. Var() method returns an integer, while a Var's value is a string. Returning 0.")
	return 0
}

//-------------- Control struct --------------
type Control struct {
	left	AST
	right	AST
	op		string
	token	Token
}

func newControl(l AST, tok Token, r AST) *Control {
	result := Control{}
	result.left	 = l
	result.token = tok
	result.op	 = tok.value.(string)
	result.right = r
	return &result
}

func (c *Control) Type() TreeType {
	return CONTROL
}

func (c *Control) Op() string {
	return c.op
}

func (c *Control) Left() AST {
	return c.left
}

func (c *Control) Right() AST {
	return c.right
}

func (c *Control) Value() int {
	log.Println("Error: trying to acess a Control's Value(). Returning 0.")
	return 0
}

//-------------- Argument struct --------------
type Argument struct {
	op		string
}
func newArgument(ident string) *Argument {
	result := Argument{}
	result.op	 = ident
	return &result
}

func (ar *Argument) Type() TreeType {
	return ARGUMENT
}

func (ar *Argument) Op() string {
	if VERBOSE { fmt.Println("An Argument's Op() is its identifier.")}
	return ar.op
}

func (ar *Argument) Left() AST {
	log.Println("Error: trying to access an Argument's child. Returning itself.")
	return ar
}

func (ar *Argument) Right() AST {
	log.Println("Error: trying to access an Argument's child. Returning itself.")
	return ar
}

func (ar *Argument) Value() int {
	log.Println("Error: trying to access an Argument's Value(). Returning 0.")
	return 0
}
//------------------------------------
//-------------- PARSER --------------
//------------------------------------
// A function to create a Parser nicely
func newParser(l Lexer) *Parser {
	result := Parser{}
	result.lexer = l
	result.currentToken = result.lexer.getNextToken()
	return &result
}

// Parser method that checks whether the current token is of a certain syntactic category and, if it is, goes to the next token.
func (p *Parser) eat(catg Cat) {
	if p.currentToken.category ==  catg {
		if VERBOSE { fmt.Println("Successfully eaten ", catg); fmt.Println("Current token :", p.currentToken) }
		if catg != EOF {
			p.currentToken = p.lexer.getNextToken()
		}
	} else {
		// Handle errors here in the future
		log.Println("Things went wrong with eat, I expected a ", catg, "and got a ", p.currentToken.category)
	}
}

// Parser method that peeks the next token, if it exists.
func (p *Parser) peekToken() Token {
	var peekToken Token
	if p.currentToken.category != EOF {
		if VERBOSE { fmt.Println("We're running peekToken at pos: ", p.lexer.pos)}
		oldPos := p.lexer.pos
		peekToken = p.lexer.getNextToken()
		if VERBOSE { fmt.Println("peekToken() is giving this token inside the if: ", peekToken)}
		p.lexer.pos = oldPos
		if VERBOSE { fmt.Println("If things went right these two numbers should be the same", oldPos, p.lexer.pos);
					 fmt.Println("Current token :", p.currentToken)}
	}
	if VERBOSE { fmt.Println("peekToken() is giving this token: ", peekToken)}
	return peekToken
}

// Parser method that evaluates a sub-factor (that is, exponential expressions).
func (p *Parser) subfactor() AST {
	currToken := p.currentToken
	var result AST
	if currToken.category == INTEGER {
		p.eat(INTEGER)
		result = newNumber(currToken)
	} else if currToken.category == ID {
		result = p.variable()
	}
	return result
}

// Parser method that evaluates a factor
// Note the wrong precedence of the operator exp
func (p *Parser) factor() AST {
	var result AST
	switch p.currentToken.category {
		case INTEGER, ID:
			result = p.subfactor()
			
			for p.currentToken.category == OP && mijn.IsSubfactorOp(p.currentToken.value.(string)) {
				op := p.currentToken
				p.eat(OP)
				
				result = newBinOp(result, op, p.subfactor())
			}
		case LPAR:
			p.eat(LPAR)
			
			result = p.expr()
			
			p.eat(RPAR)
		
		case OP:
			op := p.currentToken
			p.eat(OP)
			
			result = newUnOp(p.factor(), op)
		
		/*case ID:
			varName := p.currentToken
			p.eat(ID)
			
			result = newVar(varName)*/
			
		default:
			log.Println("Syntactic error. Expected LPAR or INTEGER. Returning 0")
	}
	if VERBOSE { fmt.Println("This is what we're returning in factor(): ", result); fmt.Println("Current token :", p.currentToken)}
	return result
}

// Parser method that evaluates a term
func (p *Parser) term() AST {
	var result AST
	switch p.currentToken.category {
		case INTEGER:
			result = p.factor()
			for p.currentToken.category == OP && mijn.IsFactorOp(p.currentToken.value.(string)) {
				op := p.currentToken
				p.eat(OP)
				
				result = newBinOp(result, op, p.factor())
			}
		case LPAR, OP, ID:
			result = p.factor()
			
		default:
			log.Println("Syntactic error. Expected LPAR or INTEGER. Returning 0")
	}
	if VERBOSE { fmt.Println("This is what we're returning in term(): ", result)}
	return result
}

// Parser method that evaluates an expression.
func (p *Parser) expr() AST {
	var result AST
	
	switch p.currentToken.category {
		case INTEGER:
			result = p.term()
			
			for p.currentToken.category == OP {
				
				op := p.currentToken
				p.eat(OP)
				
				result = newBinOp(result, op, p.term())
			}
			
		case LPAR, OP, ID:
			result = p.term()
		
		default:
			log.Println("First token should be an INTEGER, an OP or a LPAR. Returning 0.")
			
	}
	if VERBOSE { fmt.Println("This is what we're returning in expr(): ", result); fmt.Println("Current token :", p.currentToken)}
	return result
}

// Parser method to start evaluating an expression.
func (p *Parser) parse() AST {	
	result := p.program()
	if p.currentToken.category != EOF {
		log.Println("Syntax error.")
	}
	return result
}

//----------------------------------------------------
//---------------- EXPERIMENTAL STUFF ----------------
//----------------------------------------------------

func (p *Parser) program() AST {
	p.eat(MAIN)
	p.eat(LBRACKET)
	argm := p.argument()
	p.eat(RBRACKET)
	progr := p.statementList()
	p.eat(END)
	
	result := newCompound(argm, progr)
	return result
}

func (p *Parser) statementList() AST {
	firstNode := p.statement()
	
	var secondNode AST
	var result AST
	
	if firstNode.Type() != NOOP {
		secondNode = p.statementList()
		result = newCompound(firstNode, secondNode)
	} else {
		result = p.empty()
	}
	
	if VERBOSE { fmt.Println("This is what we're returning in statementList(): ", result); fmt.Println("Current token :", p.currentToken)}
	return result
}

// Parser method statement(). We use a goto to be able to catch IDs both as expressions (they can be used as integers) and as assignments.
// If the next token is a "=", then we fallthrough to a default assignment case.
func (p *Parser) statement() AST {
	var result AST
	switch p.currentToken.category {
		case RPAR, END, EOF, EOL:				//Not a statement. Eat the category and return a NoOp
			//p.eat(p.currentToken.category)
			result = p.empty()
			
		case IF, PRINT, FOR:		//ORDER
			result = p.order()
			
		case INTEGER, OP, LPAR: 	//EXPR
			
			result = p.expr()
		
		case ID:
			if p.peekToken().category == ASSIGN {
				if VERBOSE { fmt.Println("The next token was an assigment so we assume this ID is used in an assignment")
							 fmt.Println("Current token :", p.currentToken)}
				goto keepLooking
			} else {
				result = p.expr()
			}
			break
		keepLooking:
			fallthrough
			
		default:					//ASSIGNMENT
			if VERBOSE { fmt.Println("We fell-through in the switch inside statement(), so we assume this is an assignment.") }
			result = p.assignmentS()
	}
	if VERBOSE { fmt.Println("This is what we're returning in statement(): ", result);
				 fmt.Println("Current token :", p.currentToken)}
	return result
}

func (p *Parser) assignmentS() AST {
	var result AST
	
	left := p.variable()
	token := p.currentToken
	
	p.eat(ASSIGN)
	
	right := p.expr()
	
	result = newAssign(left, token, right)
	
	if VERBOSE { fmt.Println("This is what we're returning in assignmentS(): ", result);
				 fmt.Println("Current token :", p.currentToken)}
	return result
}

func (p *Parser) variable() AST {
	var result AST
	
	result = newVar(p.currentToken)
	p.eat(ID)
	
	if VERBOSE { fmt.Println("This is what we're returning in variable(): ", result);
				 fmt.Println("Current token :", p.currentToken)}
	return result
}

func (p *Parser) empty() AST {
	return newNoOp()
}

func (p *Parser) order() AST {
	var result AST
	switch p.currentToken.category {
		case IF:
			token := p.currentToken
			p.eat(IF)
			
			left := p.expr()
			p.eat(THEN)
			
			right := p.statement()
			p.eat(END)
			
			result = newControl(left, token, right)
			
		case FOR:
			token := p.currentToken
			p.eat(FOR)
			
			left := p.expr()
			p.eat(DO)
			
			right := p.statementList()
			p.eat(END)
			
			result = newControl(left, token, right)
			
		//case PRINT: We'll leave this for later
	}
	if VERBOSE { fmt.Println("This is what we're returning in order(): ", result);
				 fmt.Println("Current token :", p.currentToken)}
	return result
}

func (p *Parser) argument() AST {
	var result AST
	currToken := p.currentToken
	p.eat(ID)
	result = newArgument(currToken.value.(string))
	return result
}
