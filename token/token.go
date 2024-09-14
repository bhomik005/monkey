package token

// A token is represented by the type and its value

type TokenType string

type Token struct {
	Type    TokenType // type of the token like "Keyword"
	Literal string    // actual content of the token "let"
}

// Defining the possible token types as constants
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, x, y, foobar, ....
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
