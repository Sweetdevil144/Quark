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

	// TRUE FALSE are boolean values.
	TRUE  = "TRUE"
	FALSE = "FALSE"

	// IF ELSE are conditional statements.
	IF   = "IF"
	ELSE = "ELSE"

	// RETURN is a keyword for return statements.
	RETURN = "RETURN"

	// WHILE is a keyword for while loops.
	WHILE = "WHILE"

	// FOR is a keyword for the for loops.
	FOR = "FOR"

	// BREAK is a keyword for the break statement.
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	SWITCH   = "SWITCH"
	CASE     = "CASE"

	// LT GT LTE GTE are comparison operators.
	LT  = "<"
	GT  = ">"
	LTE = "<="
	GTE = ">="

	COLON = "COLON"
)

var keywords = map[string]TypeOfToken{
	"fn":  FUNCTION,
	"let": LET,

	// Added keywords for boolean values.
	"true":  TRUE,
	"false": FALSE,

	// Added keywords for conditional statements.
	"if":   IF,
	"else": ELSE,

	// Added keyword for return statements.
	"return": RETURN,

	// Added keyword for while loops.
	"while": WHILE,

	// Added keyword for for loops.
	"for": FOR,

	// Added keyword for switch case statements.
	"break":    BREAK,
	"continue": CONTINUE,
	"switch":   SWITCH,
	"case":     CASE,

	// Added keywords for comparison operators.
	"<":  LT,
	">":  GT,
	"<=": LTE,
	">=": GTE,
}

func LookUpIdentifier(ident string) TypeOfToken {
	tok, ok := keywords[ident]
	if ok {
		return tok
	}
	return IDENT
}
