package be

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"
	STRING     = "STRING"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LESS_THAN    = "<"
	GREATER_THAN = ">"

	EQUAL     = "=="
	NOT_EQUAL = "!="

	// Delimiters
	COMMA = ","

	LEFT_PAREN    = "("
	RIGHT_PAREN   = ")"
	LEFT_BRACKET  = "["
	RIGHT_BRACKET = "]"

	// Keywords
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	RETURN = "RETURN"

	AND = "AND"
	OR  = "OR"
)

// Token ...
type Token struct {
	Type    TokenType
	Literal string
}

// NewToken ...
func NewToken(tokenType TokenType, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}

var keywords = map[string]TokenType{
	"true":   TRUE,
	"false":  FALSE,
	"and":    AND,
	"or":     OR,
	"return": RETURN,
}

// LookupIdent ...
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENTIFIER
}
