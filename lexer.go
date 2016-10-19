package main

import (
	"./mijn"
	"bytes"
	"unicode"
	"strconv"
	"log"
	"fmt"
)

type Lexer struct {
	text 			string
	pos 			int
}

func newLexer(txt string) *Lexer {
	result := Lexer{}
	result.text = txt
	result.pos = 0
	return &result
}

// Lexer method that gets the current char.
func (l *Lexer) getCurrChar() byte {
	return l.text[l.pos]
}

// Lexer method that peeks the next char.
func (l *Lexer) peek() byte {
	peekPos := l.pos + 1
	if peekPos > len(l.text) - 1 {
		return '\x00'
	} else {
		return l.text[peekPos]
	}
}

// Lexer method that skips all following whitespace
func (l *Lexer) skipWhite() {
	for mijn.IsBlank(l.text[l.pos]) && l.pos < len(l.text) - 1 {
		l.pos += 1
	}
}

// Lexer method that gets a string containing an integer starting in the current char.
func (l *Lexer) getIntegerStr() string {
	var buffer bytes.Buffer
		
	for unicode.IsDigit(rune(l.getCurrChar())) {
		
		buffer.WriteString(string(l.getCurrChar()))
		
		if l.pos < len(l.text) - 1 {
			l.pos += 1
		} else {
			break
		}
	}
	
	if buffer.String() == "" {
		log.Println("Error: there wasn't an integer for getIntegerStr()")
	}
	
	return buffer.String()
}

// Lexer method that gets a parenthesis.
func (l *Lexer) getParen() string {
	if l.getCurrChar() == '(' {
		l.pos += 1
		return "("
	} else if l.getCurrChar() == ')' {
		l.pos += 1
		return ")"
	}
	
	log.Println("Error: there wasn't a parenthesis for getParen()")
	return ""
}

// Lexer method that gets an operator starting in the current char.
func (l *Lexer) getOperator() string {
	var buffer bytes.Buffer
	
	for unicode.IsLetter(rune(l.getCurrChar())) || unicode.IsNumber(rune(l.getCurrChar())) {
		
		buffer.WriteString(string(l.getCurrChar()))
		
		if l.pos < len(l.text) -1 {
			l.pos += 1
		} else {
			break
		}
	}
	
	if buffer.String() != "" {
		return buffer.String()
	} else {
		opChar := l.getCurrChar()
		l.pos += 1
		return string(opChar)
	}
}

// Lexer method that gets the next token starting in its current position.
func (l *Lexer) getNextToken() Token {
	
	var currChar byte
	
	for l.pos < len(l.text) - 1 {
		
		currChar = l.getCurrChar()
		
		if mijn.IsBlank(currChar) {
			l.skipWhite()
			continue
		}
	
		if unicode.IsDigit(rune(currChar)) {
			representedNumber, _ := strconv.Atoi(l.getIntegerStr())
			mToken := Token{INTEGER, representedNumber}
			return mToken
		
		} else if unicode.IsLetter(rune(currChar)) || mijn.IsOp(currChar) {
			if VERBOSE { fmt.Println("We recognise this character: ", string(currChar), " as one that can be an operator, an ID, or something of the sort")}
			opr := l.getOperator()
			
			var mToken Token
			
			if mijn.IsOpStr(opr) {
				mToken = Token{OP, opr}
			} else if mijn.IsSpecStr(opr) {
				switch opr {
					case "main", "m":	mToken = Token{MAIN, "main"}
					case "end", "e":	mToken = Token{END, "end"}
					case "print":		mToken = Token{PRINT, "print"}
					case "if":			mToken = Token{IF, "if"}
					case "then":		mToken = Token{THEN, "then"}
					case "=":			mToken = Token{ASSIGN, "="}
					case "for":			mToken = Token{FOR, "for"}
					case "do":			mToken = Token{DO, "do"}
					case ",":			mToken = Token{COMMA, ","}
					case "[":			mToken = Token{LBRACKET, "["}
					case "]":			mToken = Token{RBRACKET, "]"}
					default:			log.Println("Mismatch between IsSpecStr and the list of special strings.")
				}
			} else {
				mToken = Token{ID, opr}
			}
			
			return mToken
		
		} else if mijn.IsPar(currChar) {
			var mToken Token
			if currChar == '(' {
				mToken = Token{LPAR, l.getParen()}
			} else {
				mToken = Token{RPAR, l.getParen()}
			}
			return mToken
		
		} else {
			log.Println("Things went wrong with getNextToken()")
			return Token{}
		}
	}
	
	return Token{EOF, ""}
}
