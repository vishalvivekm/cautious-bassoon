package token

type TokenType string 

type Token struct {
	Type    TokenType
	Literal string
}

// TokenTypes
const (
	ILLEGAL = "ILLEGAl"
	EOF = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // x, y, add ...
	INT = "INT" // 34342

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)
var keywords = map[string]TokenType {
		"fn": FUNCTION,
		"let": LET,
	}
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
			return tok
		}
		return IDENT
	}
