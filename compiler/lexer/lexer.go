package lexer

import (
	"errors"

	"github.com/yakawa/makeDatabase/common/token"
)

type lexer struct {
	src        []rune
	currentPos int
	readPos    int
	pos        int
	line       int
}

// Tokenize
func Tokenize(src string) (tokens []token.Token, err error) {
	l := new(src)
	tokens, err = l.tokenize()
	return
}

func new(src string) *lexer {
	return &lexer{
		src: []rune(src),
	}
}

func (l *lexer) tokenize() (tokens []token.Token, err error) {
	ch := l.readChar()
	for ch != 0 {
		switch ch {
		case '"':
			if l.peekChar() == 0 {
				tokens = append(tokens, l.makeToken([]rune("\""), token.DOUBLEQUOTE))
			} else {
				ln := l.line
				po := l.pos

				s, e := l.readString('"')
				if e != nil {
					tokens = append(tokens, token.Token{Type: token.INVALID})
					return tokens, e
				}
				tokens = append(tokens, token.Token{Type: token.STRING, Literal: s, Pos: po, Line: ln})
			}
		case '%':
			tokens = append(tokens, l.makeToken([]rune("%"), token.PERCENT))
		case '&':
			tokens = append(tokens, l.makeToken([]rune("&"), token.AMPERSAND))
		case '\'':
			if l.peekChar() == 0 {
				tokens = append(tokens, l.makeToken([]rune("'"), token.QUOTE))
			} else {
				ln := l.line
				po := l.pos

				s, e := l.readString('\'')
				if e != nil {
					tokens = append(tokens, token.Token{Type: token.INVALID})
					return tokens, e
				}
				tokens = append(tokens, token.Token{Type: token.STRING, Literal: s, Pos: po, Line: ln})
			}
		case '(':
			tokens = append(tokens, l.makeToken([]rune("("), token.LEFTPAREN))
		case ')':
			tokens = append(tokens, l.makeToken([]rune(")"), token.RIGHTPAREN))
		case '*':
			tokens = append(tokens, l.makeToken([]rune("*"), token.ASTERISK))
		case '+':
			tokens = append(tokens, l.makeToken([]rune("+"), token.PLUSSIGN))
		case ',':
			tokens = append(tokens, l.makeToken([]rune(","), token.COMMA))
		case '-':
			tokens = append(tokens, l.makeToken([]rune("-"), token.MINUSSIGN))
		case '.':
			if isDigit(l.peekChar()) {
				ln := l.line
				po := l.pos

				n := l.readNumber(ch)
				tokens = append(tokens, token.Token{Type: token.NUMBER, Literal: n, Pos: po, Line: ln})
			} else {
				tokens = append(tokens, l.makeToken([]rune("."), token.PERIOD))
			}
		case '/':
			tokens = append(tokens, l.makeToken([]rune("/"), token.SOLIDAS))
		case ':':
			tokens = append(tokens, l.makeToken([]rune(":"), token.COLON))
		case ';':
			tokens = append(tokens, l.makeToken([]rune(";"), token.SEMICOLON))
		case '<':
			if l.peekChar() == '>' {
				l.readChar()
				tokens = append(tokens, l.makeToken([]rune("<>"), token.NOTEQUALS))
			} else if l.peekChar() == '=' {
				l.readChar()
				tokens = append(tokens, l.makeToken([]rune("<="), token.LESSTHANEQUALS))
			} else {
				tokens = append(tokens, l.makeToken([]rune("<"), token.LESSTHAN))
			}
		case '=':
			tokens = append(tokens, l.makeToken([]rune("="), token.EQUALS))
		case '>':
			if l.peekChar() == '=' {
				l.readChar()
				tokens = append(tokens, l.makeToken([]rune(">="), token.GREATERTHANEQUALS))
			} else {
				tokens = append(tokens, l.makeToken([]rune(">"), token.GREATERTHAN))
			}
		case '?':
			tokens = append(tokens, l.makeToken([]rune("?"), token.QUESTION))
		case '[':
			tokens = append(tokens, l.makeToken([]rune("["), token.LEFTBRACKET))
		case ']':
			tokens = append(tokens, l.makeToken([]rune("]"), token.RIGHTBRACKET))
		case '^':
			tokens = append(tokens, l.makeToken([]rune("^"), token.CIRCUMFLEX))
		case '_':
			tokens = append(tokens, l.makeToken([]rune("_"), token.UNDERSCORE))
		case '|':
			if l.peekChar() == '|' {
				l.readChar()
				tokens = append(tokens, l.makeToken([]rune("||"), token.CONCAT))
			} else {
				tokens = append(tokens, l.makeToken([]rune("|"), token.VERTICALBAR))
			}
		case '{':
			tokens = append(tokens, l.makeToken([]rune("{"), token.LEFTBRACE))
		case '}':
			tokens = append(tokens, l.makeToken([]rune("}"), token.RIGHTBRACE))
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			ln := l.line
			po := l.pos

			n := l.readNumber(ch)
			tokens = append(tokens, token.Token{Type: token.NUMBER, Literal: n, Pos: po, Line: ln})
		default:
			ln := l.line
			po := l.pos

			s := l.readIdentifier(ch)
			t := token.LookupKeyword(s)
			tokens = append(tokens, token.Token{Type: t, Literal: s, Pos: po, Line: ln})
		}
		for {
			ch = l.readChar()
			if ch == 0 || !(ch == ' ' || ch == '\n' || ch == '\t') {
				break
			}
		}
	}
	tokens = append(tokens, token.Token{Type: token.EOS})
	return
}

func (l *lexer) readChar() rune {
	if l.currentPos >= len(l.src) {
		return 0
	}
	r := l.src[l.currentPos]
	l.readPos = l.currentPos
	l.currentPos++
	return r
}

func (l *lexer) peekChar() rune {
	if l.currentPos >= len(l.src) {
		return 0
	}
	return l.src[l.currentPos]
}

func (l *lexer) makeToken(s []rune, t token.Type) (r token.Token) {
	r.Type = t
	r.Literal = string(s)

	r.Pos = l.pos
	r.Line = l.line

	return r
}

func (l *lexer) readString(q rune) (s string, err error) {
	s = ""
	ch := l.readChar()
	loop := true
	for loop && ch != 0 {
		if ch == q {
			if l.peekChar() == q {
				l.readChar()
				s += string(ch)
				l.readChar()
			} else {
				loop = false
			}
		} else if ch == '\n' {
			err = errors.New("INVALID String Token")
			return
		} else {
			s += string(ch)
			ch = l.readChar()
		}
	}
	return
}

func (l *lexer) readNumber(ch rune) (s string) {
	s = string(ch)
	for isDigit(l.peekChar()) {
		ch = l.readChar()
		s += string(ch)
	}
	if l.peekChar() != '.' {
		return
	}
	ch = l.readChar()
	s += string(ch)
	for isDigit(l.peekChar()) {
		ch = l.readChar()
		s += string(ch)
	}
	return
}

func (l *lexer) readIdentifier(ch rune) (s string) {
	s = string(ch)
	for !isSymbol(l.peekChar()) || (l.peekChar() == '_') {
		ch = l.readChar()
		s += string(ch)
	}
	return
}

func isSymbol(ch rune) bool {
	switch ch {
	case '"', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '_', '|', '{', '}', '\n', '\t', ' ', 0:
		return true
	}
	return false
}

func isDigit(ch rune) bool {
	if '0' <= ch && ch <= '9' {
		return true
	}
	return false
}
