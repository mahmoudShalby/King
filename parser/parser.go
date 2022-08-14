package parser

import "fmt"

type StatementType uint8

const (
	NAME StatementType = iota
)

type Statement struct {
	T StatementType
	V []string
}

type Parser struct {
	tokens        []Token
	tokensLength  uint
	tokensPointer uint
	currentItem   Token
	statements    []Statement
	currentLine   uint
}

// Print tokens that collected by collectTokens
func (l *Lexer) printTokens() {
	fmt.Println("\033[1;32mTokens:\033[0m")
	for _, token := range l.tokens {
		fmt.Printf("\033[1;37m%v: %v\033[0m\n", token.T, token.V.String())
	}
}

func (p *Parser) Init(text string) {
	var l Lexer
	l.init(text)
	l.printTokens()
	p.tokens = l.tokens
	p.tokensLength = uint(len(p.tokens))
	p.currentItem = p.tokens[0]
	p.currentLine = 1
	p.collectStatements()
}

func (p *Parser) printError(msg string) {
	fmt.Printf("Error in line %d:\n\t\033[31m%s\033[0m", p.currentLine, msg)
}

func (p *Parser) next() {
	p.tokensPointer++
	if p.tokensPointer < p.tokensLength {
		p.currentItem = p.tokens[p.tokensPointer]
		if p.currentItem.T == NEWLINE {
			p.currentLine++
		}
	} else {
		p.currentItem.T = NOTHING
	}
}

func (p *Parser) collectName() {
	p.next()
	if p.currentItem.T == WORD {
		name := p.currentItem.V.String()
		println(name)
	} else {
		p.printError("You should write a name for the word")
	}
}

func (p *Parser) collectStatements() {
	for p.currentItem.T != NOTHING {
		switch {
		case p.currentItem.T == KEYWORD:
			if p.currentItem.V.String() == "name" {
				p.collectName()
			}
		}
	}
}
