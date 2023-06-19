package token

// TokenType allows multiple types to be defined.
// Note that this isn't the most performative, but for the book example it is okay.
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // Unknown token to our lexer
	EOF     = "EOF"     // Tells the parser is can stop

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	EXCLAIM  = "!"
	MULTIPLY = "*"
	FSLASH   = "/"

	LT = "<"
	GT = ">"

    EQUAL = "=="
    NOTEQUAL = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var (
	keywords = map[string]TokenType{
		"fn":     FUNCTION,
		"let":    LET,
		"true":   TRUE,
		"false":  FALSE,
		"if":     IF,
		"else":   ELSE,
		"return": RETURN,
	}
)

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok // If valid in table, return the token type
	}
	return IDENT // Return general identifier constant.
}
