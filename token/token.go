package token

// TypeOfToken defines the type of lexical tokens in the programming language.
// TypeOfToken here is an alias for string.
type TypeOfToken string

// Token represents a lexical token with a specific type and literal value.
type Token struct {
	Type    TypeOfToken // Type indicates the category of the token.
	Literal string      // Literal is the textual representation of the token.
}

// Definition of token types, including both language constructs and syntactical markers.
const (
	// ILLEGAL represents an illegal or unknown token; used for error handling in parsing.
	ILLEGAL = "ILLEGAL"

	// EOF signifies "end of file", used to indicate that no more tokens are available for reading.
	EOF = "EOF"

	// IDENT and INT are token types for identifiers and integer literals.
	IDENT = "IDENT"
	INT   = "INT"

	// ASSIGN and PLUS represent the assignment operator "=" and the addition operator "+", respectively.
	ASSIGN = "="
	PLUS   = "+"

	// COMMA and SEMICOLON are DELIMITERS used in the syntax for separating elements.
	COMMA     = ","
	SEMICOLON = ";"

	// FUNCTION and LET are keywords in the language, denoting function declarations and variable assignments.
	FUNCTION = "FUNCTION"
	LET      = "LET"

	// LPAREN RPAREN LBRACE RBRACE are Parenthetical tokens used to denote grouped expressions and code blocks.
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
)
