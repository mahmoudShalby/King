package parser

import "bytes"

type TokenType uint8

const (
	NOTHING TokenType = iota
	KEYWORD
	NAME
	INT
	FLOAT
	STRING
	BOOL
	PUNCTUATION
	NEWLINE
	TAB
)

type Token struct {
	T TokenType
	V *bytes.Buffer
}
