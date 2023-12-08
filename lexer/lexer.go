package lexer

import (
	"Quark/token"
)

// TokenizeInputString is the primary method for tokenizing the input string.
// It skips whitespace and identifies the correct token type,
// returning a single token each time it's called.
func (l *Lexer) TokenizeInputString() token.Token {
	var tok token.Token

	l.skipWhitespace()
	// Defining switch case for all the possible tokens. Where tokens are keywords, operators, identifiers, etc.
	// Among which switch case is aimed for operators.
	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			ch := l.char
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.char)}
		} else {
			tok = newToken(token.ASSIGN, l.char)
		}
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
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '!':
		if l.peekChar() == '=' {
			ch := l.char
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.char)}
		} else {
			tok = newToken(token.BANG, l.char)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.char)
	case '/':
		tok = newToken(token.SLASH, l.char)
	case '%':
		tok = newToken(token.MOD, l.char)
	case '<':
		tok = newToken(token.LT, l.char)
	case '>':
		tok = newToken(token.GT, l.char)
	case ':':
		tok = newToken(token.COLON, l.char)
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

// Lexer struct defines the structure for our lexical analyzer.
type Lexer struct {
	input        string // The input string being analyzed
	position     int    // Current position in the input (points to current char)
	char         byte   // Current char under examination.
	readPosition int    // Next reading position in the input (after current char). The default value is 0.
}

// New initializes a new Lexer with the input and prepares it to parse.
// It returns a pointer to the Lexer.
// In simple terms, it's a constructor for the Lexer and returns a pointer to it.
func New(inputString string) *Lexer {
	l := &Lexer{input: inputString}
	l.readChar() // Initialize reading the first character
	//fmt.Println(l.char," ",l.position," ",l.readPosition)
	//Output is: 108   0   1 for inputString = "let five = 5; at position 0. As the readPosition is 1, the next character is e.
	//Then, the readPosition is incremented to 2. Output is: 101   1   2. The next character is t.
	return l
}

// readNumber reads in a number from the current position
// until a non-digit character is encountered.
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
//
//	For example, if the input string is "let five = 5;", the readIdentifier() method will return "let" as the identifier.
func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

// isLetter checks whether a character is a letter or underscore,
// which are allowed in identifiers.
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
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

// peekChar returns the next character in the input string without advancing the read position.
// It's used to look ahead in the input to see what the next character is.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
