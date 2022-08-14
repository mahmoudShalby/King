package parser

import "bytes"

type StatementType uint8

const (
	NAME StatementType = iota
)

type Statement struct {
	T StatementType
	V []string
}

type Parser struct {
	tokens          []Token
	tokensLength     uint
	tokensPointer 	 uint
	currentItem   	 Token
	emptyItem				 Token
	statements       []Statement
	currentLine			 uint
}

func (p *Parser) Init(text string) {
	var l Lexer
	l.init(text)
	p.tokens = l.tokens
	p.tokensLength = uint(len(p.tokens))
	p.currentItem = p.tokens[0]
	var emptyBytesBuffer bytes.Buffer
	p.emptyItem = Token{NOTHING, emptyBytesBuffer}
	p.collectStatements()
}

func (p *Parser) next() {
	p.tokensPointer++
	if p.tokensPointer < p.tokensLength {
		p.currentItem = p.tokens[p.tokensPointer]
	} else {
		p.currentItem = p.emptyItem
	}
}

func (p *Parser) collectName() {
	p.next()
	if p.currentItem.T == WORD {
		name := p.currentItem.V.String()
	}
}

func (p *Parser) collectStatements() {
	for p.currentItem != p.emptyItem {
		switch {
		case p.currentItem.T == KEYWORD:
			if p.currentItem.V.String() == "name" {
				p.collectName()
			}
		}
	}
}
