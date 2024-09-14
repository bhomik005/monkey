package token

// A token is represented by the type and its value

type TokenType string

type Token struct {
	Type    TokenType // type of the token like "Keyword"
	Literal string    // actual content of the token "let"
}
