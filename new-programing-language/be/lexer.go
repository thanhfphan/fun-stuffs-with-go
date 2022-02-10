package be

// Lexer ...
type Lexer struct {
	input        string
	position     int  // points to current char
	readPosition int  // current reading position
	ch           byte // current char under examination
}

// NewLexer ...
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken ...
func (l *Lexer) NextToken() Token {
	var tok Token
	l.skipWhitespaces()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = NewToken(EQUAL, "==")
		} else {
			tok = NewToken(ASSIGN, "=")
		}
	case '+':
		tok = NewToken(PLUS, "+")
	case '-':
		tok = NewToken(MINUS, "-")
	case '/':
		tok = NewToken(SLASH, "/")
	case '*':
		tok = NewToken(ASTERISK, "*")
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = NewToken(NOT_EQUAL, "!=")
		} else {
			tok = NewToken(BANG, "!")
		}
	case '<':
		tok = NewToken(LESS_THAN, "<")
	case '>':
		tok = NewToken(GREATER_THAN, ">")
	case ',':
		tok = NewToken(COMMA, ",")
	case '(':
		tok = NewToken(LEFT_PAREN, "(")
	case ')':
		tok = NewToken(RIGHT_PAREN, ")")
	case '[':
		tok = NewToken(LEFT_BRACKET, "[")
	case ']':
		tok = NewToken(RIGHT_BRACKET, "]")
	case '"':
		tok = NewToken(STRING, l.readString())
	case 0:
		tok = NewToken(EOF, "")
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			return Token{Type: INTEGER, Literal: l.readNumber()}
		} else {
			tok = NewToken(ILLEGAL, string(l.ch))
		}
	}

	l.readChar()

	return tok
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition = l.readPosition + 1
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *Lexer) readString() string {
	pos := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[pos:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}
