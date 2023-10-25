package lexer

import (
	"Quark/token"
	"testing"
)

// TestNextToken verifies the lexical analysis for a given input string.
// It ensures that each token is correctly identified and lexed.
func TestNextToken(t *testing.T) {
	// Test input string containing various tokens.
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
		x + y;
	};
	
	let result = add(five, ten);
	`

	// tests defines a slice of test cases, each containing the expected token type and literal value.
	tests := []struct {
		expectedType    token.TypeOfToken
		expectedLiteral string
	}{
		{token.LET, "let"}, {token.IDENT, "five"}, {token.ASSIGN, "="}, {token.INT, "5"},
		{token.SEMICOLON, ";"}, {token.LET, "let"}, {token.IDENT, "ten"}, {token.ASSIGN, "="},
		{token.INT, "10"}, {token.SEMICOLON, ";"}, {token.LET, "let"}, {token.IDENT, "add"},
		{token.ASSIGN, "="}, {token.FUNCTION, "fn"}, {token.LPAREN, "("}, {token.IDENT, "x"},
		{token.COMMA, ","}, {token.IDENT, "y"}, {token.RPAREN, ")"}, {token.LBRACE, "{"},
		{token.IDENT, "x"}, {token.PLUS, "+"}, {token.IDENT, "y"}, {token.SEMICOLON, ";"},
		{token.RBRACE, "}"}, {token.SEMICOLON, ";"}, {token.LET, "let"}, {token.IDENT, "result"},
		{token.ASSIGN, "="}, {token.IDENT, "add"}, {token.LPAREN, "("}, {token.IDENT, "five"},
		{token.COMMA, ","}, {token.IDENT, "ten"}, {token.RPAREN, ")"}, {token.SEMICOLON, ";"},
	}

	l := New(input)

	// Iterating through each test case to validate the token types and literals.
	for i, tt := range tests {
		tok := l.NextToken()

		// Checking for the correct token type.
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		// Checking for the correct literal value.
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
