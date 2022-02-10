package be

import (
	"bytes"
	"strings"
)

// Node ..
type Node interface {
	TokenLiteral() string
	ToString() string
}

// Statement ...
type Statement interface {
	Node
	statementNode()
}

// Expression ...
type Expression interface {
	Node
	expressionNode()
}

// Program ...
type Program struct {
	Statements []Statement
}

// ReturnStatement ...
type ReturnStatement struct {
	Token       Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral ...
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// ToString ///
func (rs *ReturnStatement) ToString() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.ToString())
	}

	return out.String()
}

// TokenLiteral ...
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// ToString ...
func (p *Program) ToString() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.ToString())
	}

	return out.String()
}

// ExpressionStatement ...
type ExpressionStatement struct {
	Token      Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral ...
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// ToString ...
func (es *ExpressionStatement) ToString() string {
	if es.Expression != nil {
		return es.Expression.ToString()
	}

	return ""
}

// Identifier ...
type Identifier struct {
	Token Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral ...
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// ToString ...
func (i *Identifier) ToString() string { return i.Value }

// BooleanLiteral ...
type BooleanLiteral struct {
	Token Token
	Value bool
}

func (b *BooleanLiteral) expressionNode() {}

// TokenLiteral ...
func (b *BooleanLiteral) TokenLiteral() string { return b.Token.Literal }

// ToString ...
func (b *BooleanLiteral) ToString() string { return b.Token.Literal }

// IntegerLiteral ...
type IntegerLiteral struct {
	Token Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral ...
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

// ToString ...
func (il *IntegerLiteral) ToString() string { return il.Token.Literal }

// PrefixExpression ...
type PrefixExpression struct {
	Token    Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral ...
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// ToString ...
func (pe *PrefixExpression) ToString() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.ToString())
	out.WriteString(")")

	return out.String()
}

// InfixExpression ...
type InfixExpression struct {
	Token    Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral ...
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

// ToString ...
func (ie *InfixExpression) ToString() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.ToString())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.ToString())
	out.WriteString(")")

	return out.String()
}

// StringLiteral ...
type StringLiteral struct {
	Token Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}

// TokenLiteral ...
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// ToString ...
func (sl *StringLiteral) ToString() string { return sl.Token.Literal }

// ArrayLiteral ...
type ArrayLiteral struct {
	Token    Token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}

// TokenLiteral ...
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }

// ToString ...
func (al *ArrayLiteral) ToString() string {
	var out bytes.Buffer

	elements := []string{}

	for _, el := range al.Elements {
		elements = append(elements, el.ToString())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
