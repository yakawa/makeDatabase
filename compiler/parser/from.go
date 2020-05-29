package parser

import (
	"errors"

	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/token"
	"github.com/yakawa/makeDatabase/logger"
)

func (p *parser) parseFrom() (fr *ast.FromClause, err error) {
	logger.Tracef("Parse: FROM Clause %s", p.currentToken.Type.String())
	defer logger.Tracef("Parse: FROM Clause End")

	fr = &ast.FromClause{}

	for i := 0; ; i++ {
		ts := &ast.TableOrSubquery{}

		if p.currentToken.Type == token.K_NATURAL {
			ts.Natural = true
			p.readToken()
		}
		if p.currentToken.Type == token.K_LEFT || p.currentToken.Type == token.K_RIGHT || p.currentToken.Type == token.K_INNER || p.currentToken.Type == token.K_CROSS {
			if p.currentToken.Type == token.K_LEFT || p.currentToken.Type == token.K_RIGHT {
				if p.currentToken.Type == token.K_LEFT {
					ts.Left = true
				}
				if p.currentToken.Type == token.K_RIGHT {
					ts.Right = true
				}
				p.readToken()
				if p.currentToken.Type == token.K_OUTER {
					p.readToken()
				}
			}
			if p.currentToken.Type == token.K_INNER {
				ts.Inner = true
				p.readToken()
			}
			if p.currentToken.Type == token.K_CROSS {
				ts.Cross = true
				p.readToken()
			}
		}

		if i != 0 {
			if ts.Natural == false && ts.Left == false && ts.Right == false && ts.Inner == false && ts.Cross == false {
				ts.Cross = true
			}
		}

		tos, er := p.parseToS()
		if er != nil {
			return fr, er
		}
		ts.Schema = tos.Schema
		ts.TableName = tos.TableName
		ts.Subquery = tos.Subquery

		if p.peekToken().Type == token.K_ON || p.peekToken().Type == token.K_USING {
			p.readToken()
			if p.currentToken.Type == token.K_ON {
				p.readToken()
				ex, er := p.parseExpression()
				if er != nil {
					return fr, er
				}
				ts.On = ex
			}
			if p.currentToken.Type == token.K_USING {
				p.readToken()
				p.readToken()
				for {
					ts.ColumnNames = append(ts.ColumnNames, p.currentToken.Literal)
					p.readToken()
					if p.currentToken.Type != token.COMMA {
						break
					}
					p.readToken()
				}
			}
		}

		fr.ToS = append(fr.ToS, *ts)
		if p.peekToken().Type != token.COMMA {
			break
		}
		p.readToken()
		p.readToken()
	}
	return
}

func (p *parser) parseToS() (ts *ast.TableOrSubquery, err error) {
	logger.Tracef("Parse: ToS")
	defer logger.Tracef("Parse: ToS End")

	ts = &ast.TableOrSubquery{}

	if p.currentToken.Type == token.LEFTPAREN {
		p.readToken()
		if p.currentToken.Type == token.K_WITH || p.currentToken.Type == token.K_SELECT || p.currentToken.Type == token.K_VALUES {
			ss, er := p.parseSelectStatement()
			if er != nil {
				return ts, er
			}
			ts.Subquery = ss
		} else {
			tos, er := p.parseToS()
			if er != nil {
				return ts, er
			}
			ts = tos
		}
		p.readToken()
		if p.currentToken.Type != token.RIGHTPAREN {
			return ts, errors.New("Parse Error Invalid Token")
		}
	} else {
		if p.peekToken().Type == token.PERIOD {
			ts.Schema = p.currentToken.Literal
			p.readToken()
			p.readToken()
		}
		ts.TableName = p.currentToken.Literal
	}
	return
}
