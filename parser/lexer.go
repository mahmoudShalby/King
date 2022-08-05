package parser

import (
	"bytes"
	"fmt"
)

type Lexer struct {
	text        string
	textLength  uint
	textPointer uint
	currentItem rune
	Tokens      []Token
}

func (l *Lexer) Init(text string) {
	l.text = text
	l.textLength = uint(len(text))
	l.currentItem = rune(text[0])
	l.collectTokens()
	l.printTokens()
}

func (l *Lexer) next() {
	l.textPointer++
	if l.textPointer < l.textLength {
		l.currentItem = rune(l.text[l.textPointer])
	} else {
		l.currentItem = 0
	}
}

func (l *Lexer) appendToken(T TokenType, V *bytes.Buffer) {
	l.Tokens = append(l.Tokens, Token{T, V})
}

func (l *Lexer) isCurrentItemLetter() bool {
	return 'a' <= l.currentItem && l.currentItem <= 'z' || 'A' <= l.currentItem && l.currentItem <= 'Z'
}

func (l *Lexer) collectName() {
	var result bytes.Buffer
	result.WriteRune(l.currentItem)
	l.next()
	for l.currentItem != 0 && (l.isCurrentItemLetter() || l.isCurrentItemNumber() || l.currentItem == ' ') {
		if l.currentItem == ' ' {
			l.next()
			if l.isCurrentItemLetter() {
				result.WriteRune(' ')
			} else {
				break
			}
		}
		result.WriteRune(l.currentItem)
		l.next()
	}
	l.appendToken(NAME, &result)
}

func (l *Lexer) isCurrentItemNumber() bool {
	return '0' <= l.currentItem && l.currentItem <= '9'
}

func (l *Lexer) collectNumber() {
	var result bytes.Buffer
	isResultHasDot := false
	result.WriteRune(l.currentItem)
	l.next()
	for l.isCurrentItemNumber() || l.currentItem == '.' {
		if l.currentItem == '.' {
			if isResultHasDot {
				break
			} else {
				isResultHasDot = true
			}
		}
		result.WriteRune(l.currentItem)
		l.next()
	}
	if isResultHasDot {
		l.appendToken(FLOAT, &result)
	} else {
		l.appendToken(INT, &result)
	}
}

func (l *Lexer) collectNewLine() {
	l.appendToken(NEWLINE, &bytes.Buffer{})
	l.next()
	for l.currentItem == '\n' {
		l.next()
	}
	fmt.Println(l.currentItem, '\n')
}

func (l *Lexer) collectString() {
	var result bytes.Buffer
	l.next()
	for l.currentItem != '"' {
		result.WriteRune(l.currentItem)
		l.next()
	}
	l.next()
	l.appendToken(STRING, &result)
}

func (l *Lexer) collectTokens() {
	for l.currentItem != 0 {
		switch {
		case l.isCurrentItemLetter():
			l.collectName()
		case l.isCurrentItemNumber():
			l.collectNumber()
		case l.currentItem == '\n':
			l.collectNewLine()
		case l.currentItem == '"':
			l.collectString()
		default:
			l.next()
		}
	}
}

func (l *Lexer) printTokens() {
	fmt.Println("Tokens:")
	for index := 0; index < len(l.Tokens); index++ {
		fmt.Printf("\t%d:\t%s\n", l.Tokens[index].T, l.Tokens[index].V.String())
	}
}
