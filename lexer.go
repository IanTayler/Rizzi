package main

import (
	"./mijn"
	"bytes"
	"unicode"
	"strconv"
	"log"
)

type Lexer struct {
	text 			string
	pos 			int
}

// Lexer method that gets the current char.
func (l *Lexer) getCurrChar() byte {
	return l.text[l.pos]
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
	
	for unicode.IsLetter(rune(l.getCurrChar())) {
		
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
			opr := l.getOperator()
			
			var mToken Token
			
			if mijn.IsOpStr(opr) {
				mToken = Token{OP, opr}
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
