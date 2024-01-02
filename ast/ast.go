package ast

import (
	"Quark/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token token.Token
	Value string
}

type LetStatement struct {
	Token token.Token // token.LET here
	Name  *Identifier
	Value Expression
}

func (letStatement *LetStatement) statementNode()       {}
func (letStatement *LetStatement) TokenLiteral() string { return letStatement.Token.Literal }

func (ident *Identifier) expressionNode()      {}
func (ident *Identifier) TokenLiteral() string { return ident.Token.Literal }
