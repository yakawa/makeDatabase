package parser

import (
	"errors"

	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/token"
	"github.com/yakawa/makeDatabase/logger"
)

func (p *parser) parseWithClause() (wc *ast.WithClause, err error) {
	logger.Tracef("Parse: WITH Clause")
	defer logger.Tracef("Parse: WITH Clause End")

	wc = &ast.WithClause{}
	p.readToken()
	if p.currentToken.Type == token.K_RECURSIVE {
		wc.IsRecursive = true
		p.readToken()
	}
	for {
		switch p.peekToken().Type {
		case token.EOS, token.SEMICOLON:
			err = errors.New("Parse Error Unexpected Terminate")
			return
		case token.K_SELECT, token.K_VALUES:
			return
		default:
			for {
				cte, e := p.parseCTE()
				if e != nil {
					return wc, e
				}
				wc.CTE = append(wc.CTE, cte)
				p.readToken()
				if p.currentToken.Type != token.COMMA {
					break
				}
				p.readToken()
			}
		}
	}
}

func (p *parser) parseCTE() (cte ast.CTE, err error) {
	logger.Tracef("Parse: CTE")
	defer logger.Tracef("Parse: CTE End")

	cte = ast.CTE{}
	if p.currentToken.Type != token.IDENT {
		err = errors.New("Parse Error Invalid Token")
		return
	}
	cte.TableName = p.currentToken.Literal
	p.readToken()

	if p.currentToken.Type == token.LEFTPAREN {
		p.readToken()
		for {
			if p.currentToken.Type != token.IDENT {
				return cte, errors.New("Parse Error Invalid Token")
			}
			cte.ColumnNames = append(cte.ColumnNames, p.currentToken.Literal)
			p.readToken()
			if p.currentToken.Type == token.RIGHTPAREN {
				p.readToken()
				break
			}
			if p.currentToken.Type != token.COMMA {
				return cte, errors.New("Parse Error Invalid Token")
			}
			p.readToken()
		}
	}

	if p.currentToken.Type != token.K_AS {
		return cte, errors.New("Parse Error Invalid Token")
	}
	p.readToken()
	if p.currentToken.Type != token.LEFTPAREN {
		return cte, errors.New("Parse Error Invalid Token")
	}
	p.readToken()
	ss, err := p.parseSelectStatement()
	if err != nil {
		return cte, err
	}
	cte.SelectStatement = ss

	if p.currentToken.Type != token.RIGHTPAREN {
		return cte, errors.New("Parse Error Invalid Token")
	}
	p.readToken()

	return
}
