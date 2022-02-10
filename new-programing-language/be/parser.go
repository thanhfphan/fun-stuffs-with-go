package be

import (
	"fmt"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
)

var precedences = map[TokenType]int{
	EQUAL:        EQUALS,
	NOT_EQUAL:    EQUALS,
	LESS_THAN:    LESSGREATER,
	GREATER_THAN: LESSGREATER,
	PLUS:         SUM,
	MINUS:        SUM,
	SLASH:        PRODUCT,
	ASTERISK:     PRODUCT,
}

type prefixParseFn func() Expression
type infixParseFn func(Expression) Expression

// Parser ...
type Parser struct {
	l              *Lexer
	errors         []string
	currentToken   Token
	peekToken      Token
	prefixParseFns map[TokenType]prefixParseFn
	infixParseFns  map[TokenType]infixParseFn
}

// NewParser ...
func NewParser(l *Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[TokenType]prefixParseFn)
	p.registerPrefix(IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(INTEGER, p.parseIntegerLiteral)
	p.registerPrefix(STRING, p.parseStringLiteral)
	p.registerPrefix(BANG, p.parsePrefixExpression)
	p.registerPrefix(MINUS, p.parsePrefixExpression)
	p.registerPrefix(TRUE, p.parseBoolean)
	p.registerPrefix(FALSE, p.parseBoolean)
	p.registerPrefix(LEFT_PAREN, p.parseGroupedExpression)

	p.infixParseFns = make(map[TokenType]infixParseFn)
	p.registerInfix(PLUS, p.parseInfixExpression)
	p.registerInfix(MINUS, p.parseInfixExpression)
	p.registerInfix(SLASH, p.parseInfixExpression)
	p.registerInfix(ASTERISK, p.parseInfixExpression)
	p.registerInfix(EQUAL, p.parseInfixExpression)
	p.registerInfix(NOT_EQUAL, p.parseInfixExpression)
	p.registerInfix(LESS_THAN, p.parseInfixExpression)
	p.registerInfix(GREATER_THAN, p.parseInfixExpression)
	p.registerInfix(LEFT_PAREN, p.parseInfixExpression)

	p.nextToken()
	p.nextToken()

	return p
}

// ParseProgram ...
func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for !p.currentTokenIs(EOF) {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.currentToken.Type {
	case RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ExpressionStatement {
	stmt := &ExpressionStatement{Token: p.currentToken}
	stmt.Expression = p.parseExpression(LOWEST)

	return stmt
}

func (p *Parser) parseReturnStatement() *ReturnStatement {
	stmt := &ReturnStatement{Token: p.currentToken}

	p.nextToken()

	stmt.ReturnValue = p.parseExpression(LOWEST)

	return stmt
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) currentTokenIs(t TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.errors = append(p.errors, fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type))
	return false
}

func (p *Parser) registerPrefix(tokenType TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokentype TokenType, fn infixParseFn) {
	p.infixParseFns[tokentype] = fn
}

func (p *Parser) parseIdentifier() Expression {
	return &Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) parseIntegerLiteral() Expression {
	literal := &IntegerLiteral{Token: p.currentToken}

	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("could not parse %q as integer", p.currentToken.Literal))
		return nil
	}
	literal.Value = value

	return literal
}

func (p *Parser) parseStringLiteral() Expression {
	return &StringLiteral{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) currentPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) parseExpression(precedence int) Expression {
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		p.errors = append(p.errors, fmt.Sprintf("not prefix function %s found", p.currentToken.Type))
		return nil
	}

	leftExp := prefix()

	for precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parsePrefixExpression() Expression {
	experssion := &PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
	}

	p.nextToken()

	experssion.Right = p.parseExpression(PREFIX)

	return experssion
}

func (p *Parser) parseInfixExpression(left Expression) Expression {
	exp := &InfixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
		Left:     left,
	}

	precedence := p.currentPrecedence()
	p.nextToken()
	exp.Right = p.parseExpression(precedence)

	return exp
}

func (p *Parser) parseBoolean() Expression {
	return &BooleanLiteral{Token: p.currentToken, Value: p.currentTokenIs(TRUE)}
}

func (p *Parser) parseGroupedExpression() Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(RIGHT_PAREN) {
		return nil
	}

	return exp
}
