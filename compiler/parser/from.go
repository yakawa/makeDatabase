package parser

import (
	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/errors"
	"github.com/yakawa/makeDatabase/common/token"
	"github.com/yakawa/makeDatabase/logger"
)

func (p *parser) parseFromClause() (fr *ast.FromClause, err error) {
	logger.Tracef("Parse: FROM Clause %#+v", p.currentToken)
	defer logger.Tracef("Parse: FROM Clause End")

	p.readToken()

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
			} else {
				if p.currentToken.Type != token.K_JOIN {
					return fr, errors.NewErrParseInvalidToken(p.currentToken)
				}
				p.readToken()
			}
		}

		tos, er := p.parseToS()
		if er != nil {
			return fr, er
		}
		ts.Schema = tos.Schema
		ts.TableName = tos.TableName
		ts.Subquery = tos.Subquery

		if p.peekToken().Type == token.K_AS {
			p.readToken()
		}
		if p.peekToken().Type == token.IDENT {
			p.readToken()
			ts.Alias = p.currentToken.Literal
		}

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

		p.readToken()
		if p.currentToken.Type != token.COMMA && !(p.currentToken.Type == token.K_LEFT || p.currentToken.Type == token.K_RIGHT || p.currentToken.Type == token.K_INNER || p.currentToken.Type == token.K_CROSS || p.currentToken.Type == token.K_NATURAL) {

			break
		}

		if p.currentToken.Type == token.COMMA {
			p.readToken()
		}
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
		if p.currentToken.Type != token.RIGHTPAREN {
			return ts, errors.NewErrParseInvalidToken(p.currentToken)
		}
	} else {
		if p.peekToken().Type == token.PERIOD {
			tmp1 := p.currentToken.Literal
			p.readToken()
			p.readToken()
			if p.peekToken().Type == token.PERIOD {
				tmp2 := p.currentToken.Literal
				p.readToken()
				p.readToken()
				ts.Schema = tmp1
				ts.Database = tmp2
			} else {
				ts.Database = tmp1
			}
		}
		ts.TableName = p.currentToken.Literal
	}
	return
}
