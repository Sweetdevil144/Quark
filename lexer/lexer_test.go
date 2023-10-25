package lexer

import (
	"Quark/token"
	"testing"
)

// Updated test cases to cover additional scenarios.
func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9;

	while (true) {
		break;
	}

	for (let x = 0; x < 10; x = x + 1) {
		continue;
	}

	switch (x) {
	case 10:
		return;
	}
	`

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

		// Newly added test cases
		{token.IF, "if"}, {token.LPAREN, "("}, {token.INT, "5"}, {token.LT, "<"},
		{token.INT, "10"}, {token.RPAREN, ")"}, {token.LBRACE, "{"}, {token.RETURN, "return"},
		{token.TRUE, "true"}, {token.SEMICOLON, ";"}, {token.RBRACE, "}"}, {token.ELSE, "else"},
		{token.LBRACE, "{"}, {token.RETURN, "return"}, {token.FALSE, "false"},
		{token.SEMICOLON, ";"}, {token.RBRACE, "}"},

		{token.WHILE, "while"}, {token.LPAREN, "("}, {token.TRUE, "true"}, {token.RPAREN, ")"},
		{token.LBRACE, "{"}, {token.BREAK, "break"}, {token.SEMICOLON, ";"}, {token.RBRACE, "}"},

		{token.FOR, "for"}, {token.LPAREN, "("}, {token.LET, "let"}, {token.IDENT, "x"},
		{token.ASSIGN, "="}, {token.INT, "0"}, {token.SEMICOLON, ";"}, {token.IDENT, "x"},
		{token.LT, "<"}, {token.INT, "10"}, {token.SEMICOLON, ";"}, {token.IDENT, "x"},
		{token.ASSIGN, "="}, {token.IDENT, "x"}, {token.PLUS, "+"}, {token.INT, "1"},
		{token.RPAREN, ")"}, {token.LBRACE, "{"}, {token.CONTINUE, "continue"}, {token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.SWITCH, "switch"}, {token.LPAREN, "("}, {token.IDENT, "x"}, {token.RPAREN, ")"},
		{token.LBRACE, "{"}, {token.CASE, "case"}, {token.INT, "10"}, {token.COLON, ":"},
		{token.RETURN, "return"}, {token.SEMICOLON, ";"}, {token.RBRACE, "}"},

		{token.EOF, ""},
	}

	l := New(input)
	// Testing each token in the input string.
	for i, tt := range tests {
		tok := l.TokenizeInputString()

		// Validating the token type.
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		// Validating the literal value.
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
