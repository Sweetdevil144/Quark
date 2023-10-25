package lexer

import "Quark/token"

// Lexer struct defines the structure for our lexical analyzer.
type Lexer struct {
	input        string // The input string being analyzed
	position     int    // Current position in the input (points to current char)
	char         byte   // Current char under examination
	readPosition int    // Next reading position in the input (after current char)
}

// New initializes a new Lexer with the input and prepares it to parse.
func New(inputString string) *Lexer {
	l := &Lexer{input: inputString}
	l.readChar() // Initialize reading the first character
	return l
}

// TokenizeInputString is the primary method for tokenizing the input string.
// It skips whitespace and identifies the correct token type,
// returning a single token each time it's called.
func (l *Lexer) TokenizeInputString() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return tok
}

// isLetter checks whether a character is a letter or underscore,
// which are allowed in identifiers.
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// readIdentifier reads in an identifier from the current position
// until a non-letter character is encountered.
func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

// newToken creates a new token of the given token type and literal value.
func newToken(tokenType token.TypeOfToken, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readChar advances the current read position in the input string.
// It handles the end of input by setting 'char' to 0, which signifies EOF.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
		// EOF
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// skipWhitespace advances the read position past any whitespace characters.
func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}
