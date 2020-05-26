package parser

import (
	"errors"

	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/token"
	"github.com/yakawa/makeDatabase/logger"
)

type parser struct {
	tokens []token.Token

	currentToken token.Token
	pos          int
}

// Parse
func Parse(tokens []token.Token) (*ast.SQL, error) {
	p := new(tokens)
	a, err := p.parse()
	return a, err
}

func new(tokens []token.Token) *parser {
	return &parser{
		tokens: tokens,
	}
}

func (p *parser) parse() (a *ast.SQL, err error) {
	logger.Tracef("Parser Start!")
	defer logger.Tracef("Parse End!")

	p.readToken()
	a = &ast.SQL{}
	for p.currentToken.Type != token.EOS {
		switch p.currentToken.Type {
		case token.SEMICOLON, token.EOS:
			return
		case token.K_WITH, token.K_SELECT:
			ss, err := p.parseSelectStatement()
			if err != nil {
				return a, err
			}
			a.SelectStatement = ss
		default:
			err = errors.New("Parse Error: Unknown Token")
			return
		}
		p.readToken()
	}
	return
}

func (p *parser) parseSelectStatement() (ss *ast.SelectStatement, err error) {
	logger.Tracef("Parse: SELECT Statement")
	defer logger.Tracef("Parse: SELECT Statement End")

	ss = &ast.SelectStatement{}
	switch p.currentToken.Type {
	case token.K_WITH:
		wc, e := p.parseWithClause()
		if e != nil {
			return ss, e
		}
		ss.WithClause = wc
	case token.K_SELECT, token.K_VALUES:
		sc, e := p.parseSelectClause()
		if e != nil {
			return ss, e
		}
		ss.SelectClause = sc
	case token.K_ORDER:
	case token.K_LIMIT:
	default:
		err = errors.New("Parse Error Invalid Token")
		return
	}
	return
}

func (p *parser) readToken() {
	if p.pos >= len(p.tokens) {
		p.currentToken = token.Token{Type: token.EOS}
	}
	p.currentToken = p.tokens[p.pos]
	p.pos++
	return
}

func (p *parser) peekToken() (t token.Token) {
	if p.pos+1 >= len(p.tokens) {
		return token.Token{Type: token.EOS}
	}
	t = p.tokens[p.pos+1]
	return
}

func (p *parser) peek2Token() (t token.Token) {
	if p.pos+2 >= len(p.tokens) {
		return token.Token{Type: token.EOS}
	}
	t = p.tokens[p.pos+2]
	return
}
