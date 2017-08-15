package ast

import (
	"github.com/mlschneid/monkey/token"
)

// Node is base type
type Node interface {
	TokenLiteral() string
}

// Statement is type of 'Node'
type Statement interface {
	Node
	statementNode()
}

// Expression is type of 'Node'
type Expression interface {
	Node
	expressionNode()
}

// Program is root node
type Program struct {
	Statements []Statement
}

// TokenLiteral implementation makes 'Program' extend 'Node'
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// LetStatement (e.g token.LET)
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// TokenLiteral : LetStatement extends 'Statement'
func (ls *LetStatement) statementNode() {}

// TokenLiteral : LetStatement extends Node
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier (e.g token.IDENT)
type Identifier struct {
	Token token.Token
	Value string
}

// ExpressionNode : Identifier extends 'Expression'
func (i *Identifier) expressionNode() {}

// TokenLiteral : Identifier extends 'Expression'
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
