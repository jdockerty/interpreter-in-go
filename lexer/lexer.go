package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // Current position at the input string
	readPosition int  // Next character/token, i.e. `position` + 1
	char         byte // Current character that is being examined, this is a concrete representation of `position`
}

func New(in string) *Lexer {
	l := &Lexer{input: in}
	l.readChar()
	return l
}

// readChar is used to move our position to the next character.
// Incrementing our position and readPosition within the lexer.
func (l *Lexer) readChar() {

	// If at greater than input, this is end of file.
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Monkey doesn't care about whitespace, so we can skip it entirely.
	l.skipWhitespace()

	// Find our current token
	switch l.char {
	case '=':
        if l.peekChar() == '=' {
            ch := l.char
            l.readChar()
            // We need to combine our token here to '==' after incrementing our position
            // Otherwise it would return [...ASSIGN, ASSIGN] as 2 tokens, which is not correct.
            tok = token.Token{Type: token.EQUAL, Literal: string(ch) + string(l.char)}
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
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '!':
		if l.peekChar() == '=' {
            ch := l.char
            l.readChar()
            tok = token.Token{Type: token.NOTEQUAL, Literal: string(ch) + string(l.char)}
		} else {
			tok = newToken(token.EXCLAIM, l.char)
		}
	case '*':
		tok = newToken(token.MULTIPLY, l.char)
	case '/':
		tok = newToken(token.FSLASH, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case '<':
		tok = newToken(token.LT, l.char)
	case '>':
		tok = newToken(token.GT, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	// Move position and readPosition
	l.readChar()
	return tok
}

// Return the next character, without incrementing our position or readPosition
// This is useful for multi-character operators like != and ==
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.char) {
		l.readChar()
	}

	return l.input[position:l.position]

}

// Reads an identifier and advance the lexer until we reach a non-letter character.
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// Helper for identifying whether a character is a letter or not
// Note the use of '_' in here too, this means that this tiny function is what enables
// the variable name of 'foo_bar', as _ is a valid character. We could even allow
// the use of ! and ? by altering this too.
func isLetter(ch byte) bool { return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' }

func isDigit(ch byte) bool { return '0' <= ch && ch <= '9' }

func newToken(tt token.TokenType, char byte) token.Token {
	return token.Token{Type: tt, Literal: string(char)}
}
