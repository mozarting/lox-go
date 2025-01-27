package lexer

import (
	"fmt"
	"strconv"

	"github.com/mozarting/lox/internal/token"
)

type Lexer struct {
	source string
	tokens []token.Token

	start   int
	current int
	line    int
}

func New(input string) *Lexer {
	return &Lexer{source: input}
}

func (l *Lexer) ScanTokens() []token.Token {
	for !l.isAtEnd() {
		l.start = l.current
		l.scanToken()
	}
	return l.tokens
}

func (l *Lexer) scanToken() {
	c := l.advance()
	switch c {
	case '(':
		l.addToken(token.LEFT_PAREN)
	case ')':
		l.addToken(token.RIGHT_PAREN)
	case '{':
		l.addToken(token.LEFT_BRACE)
	case '}':
		l.addToken(token.RIGHT_BRACE)
	case ',':
		l.addToken(token.COMMA)
	case '.':
		l.addToken(token.DOT)
	case '-':
		l.addToken(token.MINUS)
	case '+':
		l.addToken(token.PLUS)
	case ';':
		l.addToken(token.SEMICOLON)
	case '*':
		l.addToken(token.STAR)
	case '!':
		if l.match('=') {
			l.addToken(token.BANG_EQUAL)
		} else {
			l.addToken(token.BANG)
		}
	case '=':
		if l.match('=') {
			l.addToken(token.EQUAL_EQUAL)
		} else {
			l.addToken(token.EQUAL)
		}
	case '<':
		if l.match('=') {
			l.addToken(token.LESS_EQUAL)
		} else {
			l.addToken(token.LESS)
		}
	case '>':
		if l.match('=') {
			l.addToken(token.GREATER_EQUAL)
		} else {
			l.addToken(token.GREATER)
		}
	case '/':
		if l.match('/') {
			for l.peek() != '\n' && !l.isAtEnd() {
				l.advance()
			}
		} else {
			l.addToken(token.SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		l.line++
	case '"':
		l.string()
	default:
		if l.isDigit(c) {
			l.number()
		} else if l.isAlpha(c) {
			l.identifier()
		} else {
			fmt.Println("ERROR: Unexpected character")
		}
	}
}

func (l *Lexer) isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func (l *Lexer) identifier() {
	for l.isAlphaNumeric(rune(l.peek())) {
		l.advance()
	}
	text := l.source[l.start:l.current]
	tokenType, ok := token.Keywords[text]
	if !ok {
		tokenType = token.IDENTIFIER
	}

	l.addToken(tokenType)
}

func (l *Lexer) isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func (l *Lexer) isAlphaNumeric(c rune) bool {
	return l.isAlpha(c) || l.isDigit(c)
}
func (l *Lexer) number() {
	for l.isDigit(rune(l.peek())) {
		l.advance()
	}
	// Look for a fractional part.
	if l.peek() == '.' && l.isDigit(l.peekNext()) {
		// Consume the "."
		l.advance()

		for l.isDigit(rune(l.peek())) {
			l.advance()
		}
	}

	str := l.source[l.start:l.current]
	value, err := strconv.ParseFloat(str, 64) // 64 specifies double-precision
	if err != nil {
		fmt.Println("Error parsing string:", err)
		return
	}
	l.addTokenWithLiteral(token.NUMBER, value)

}

func (l *Lexer) peekNext() rune {
	if l.current+1 >= len(l.source) {
		return '\x00'
	}
	return rune(l.source[l.current+1])
}

func (l *Lexer) string() {
	for l.peek() != '"' && !l.isAtEnd() {
		if l.peek() == '\n' {
			l.line++
		}
		l.advance()
	}

	if l.isAtEnd() {
		fmt.Println("Unterminated string")
		return
	}

	// The closing ".
	l.advance()

	// Trim the surrounding quotes.
	value := l.source[l.start+1 : l.current-1]
	l.addTokenWithLiteral(token.STRING, value)
}

func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.source)
}

func (l *Lexer) advance() rune {
	current := rune(l.source[l.current])
	l.current++
	return current
}

func (l *Lexer) addToken(t token.TokenType) {
	l.addTokenWithLiteral(t, "")
}

func (l *Lexer) addTokenWithLiteral(t token.TokenType, literal interface{}) {
	text := l.source[l.start:l.current]
	l.tokens = append(l.tokens, token.Token{Type: t, Lexeme: text, Literal: literal, Line: l.line})
}

func (l *Lexer) match(expected rune) bool {
	if l.isAtEnd() || rune(l.source[l.current]) != expected {
		return false
	}

	l.current++
	return true
}

func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return '\x00'
	}
	return l.source[l.current]
}
