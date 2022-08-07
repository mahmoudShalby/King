package parser

import "bytes"

// Name that will get type KEYWORD by lexer
var keywords = [...]string{
	"if",
	"is",
	"or",
	"and",
	"loop",
	"in",
	"at",
	"on",
	"while",
}

// Lexer's base object
type Lexer struct {
	text        string
	textLength  uint
	textPointer uint
	currentItem rune
	Tokens      []Token
}

// The initialize function of lexer just call it to use lexer
func (l *Lexer) Init(text string) {
	l.text = text
	l.textLength = uint(len(text))
	l.currentItem = rune(text[0])
	l.collectTokens()
}

// Set the next character of (*Lexer).text in (*Lexer).currentItem and increment the (*Lexer).textPointer
func (l *Lexer) next() {
	l.textPointer++
	if l.textPointer < l.textLength {
		l.currentItem = rune(l.text[l.textPointer])
	} else {
		l.currentItem = 0
	}
}

// Append new token
func (l *Lexer) appendToken(T TokenType, V bytes.Buffer) {
	l.Tokens = append(l.Tokens, Token{T, V})
}

// Check if (*Lexer).currentItem is letter
func (l *Lexer) isCurrentItemLetter() bool {
	return 'a' <= l.currentItem && l.currentItem <= 'z' || 'A' <= l.currentItem && l.currentItem <= 'Z'
}

// Append name token -type> got by the value -value> @param name
func (l *Lexer) appendNameToken(name string) {
	var new_result bytes.Buffer
	new_result.WriteString(name)
	var t TokenType
	for _, keyword := range keywords {
		if name == keyword {
			t = KEYWORD
			break
		}
	}
	if t == 0 {
		if name == "true" || name == "false" {
			t = BOOL
		} else {
			t = NAME
		}
	}
	l.appendToken(t, new_result)
}

// Collect names then append it with (*Lexer).appendNameToken
func (l *Lexer) collectName() {
	var result bytes.Buffer
	result.WriteRune(l.currentItem)
	l.next()
	for l.currentItem != 0 && (l.isCurrentItemLetter() || l.currentItem == ' ') {
		if l.currentItem == ' ' {
			if result.Len() != 0 {
				l.appendNameToken(result.String())
				result.Reset()
			}
			l.next()
			continue
		}
		result.WriteRune(l.currentItem)
		l.next()
	}
	if result.Len() != 0 {
		l.appendNameToken(result.String())
	}
}

// Check if (*Lexer).currentItem is number
func (l *Lexer) isCurrentItemNumber() bool {
	return '0' <= l.currentItem && l.currentItem <= '9'
}

// Collect numbers
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
		l.appendToken(FLOAT, result)
	} else {
		l.appendToken(INT, result)
	}
}

// Check if (*Lexer).currentItem is whitespace
func (l *Lexer) isCurrentItemWhitespace() bool {
	return l.currentItem == ' ' || l.currentItem == '\n' || l.currentItem == '\t'
}

// Collect newlines -value> just one newline and ignore others to find any other thing
func (l *Lexer) collectNewLine() {
	l.appendToken(NEWLINE, bytes.Buffer{})
	l.next()
	for l.currentItem == '\n' {
		l.next()
	}
}

// Collect tabes -value> number of tabs
func (l *Lexer) collectTab() {
	var value uint8 = 1
	l.next()
	for l.currentItem == '\t' {
		value++
		l.next()
	}
	var value_as_buffer bytes.Buffer
	value_as_buffer.WriteRune(rune('0' + value))
	l.appendToken(TAB, value_as_buffer)
}

// Collect Strings
func (l *Lexer) collectString() {
	var result bytes.Buffer
	l.next()
	for l.currentItem != '"' {
		result.WriteRune(l.currentItem)
		l.next()
	}
	l.next()
	l.appendToken(STRING, result)
}

// Collect any thing isn't letter, number or whitespace
func (l *Lexer) collectPunctuation() {
	var result bytes.Buffer
	l.next()
	for !(l.isCurrentItemLetter() || l.isCurrentItemNumber() || l.isCurrentItemWhitespace()) {
		result.WriteRune(l.currentItem)
		l.next()
	}
	l.appendToken(PUNCTUATION, result)
}

// Collect tokens from (*Lexer).text
func (l *Lexer) collectTokens() {
	for l.currentItem != 0 {
		switch {
		case l.isCurrentItemLetter():
			l.collectName()
		case l.isCurrentItemNumber():
			l.collectNumber()
		case l.currentItem == '\n':
			l.collectNewLine()
		case l.currentItem == '\t':
			l.collectTab()
		case l.currentItem == '"':
			l.collectString()
		default:
			l.collectPunctuation()
		}
	}
}

// UnComment this and run this function at the end of (*Lexer).Init function
// if you want to see tokens after collect it

// Print tokens that collected by collectTokens
// func (l *Lexer) printTokens() {
// 	fmt.Println("\x1b[1;32mTokens:\x1b[0m")
// 	for _, token := range l.Tokens {
// 		fmt.Printf("\x1b[1;37m%v: %v\x1b[0m\n", token.T, token.V.String())
// 	}
// }
