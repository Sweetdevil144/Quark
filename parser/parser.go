package parser

import (
	"Quark/ast"
	"Quark/lexer"
	"Quark/token"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.TokenizeInputString()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{l: lex}
	p.nextToken()
	p.nextToken()

	return p
}
