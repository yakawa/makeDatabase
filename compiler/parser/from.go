package parser

import (
	"errors"

	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/token"
	"github.com/yakawa/makeDatabase/logger"
)

func (p *parser) parseFrom() (fr *ast.FromClause, err error) {
	logger.Tracef("Parse: FROM Clause")
	defer logger.Tracef("Parse: FROM Clause End")

	fr = &ast.FromClause{}
	for {
		ts, e := p.parseTableOrSubquery()
		if e != nil {
			return fr, e
		}
		fr.ToS = ts
		break
	}
	return
}

func (p *parser) parseTableOrSubquery() (ts *ast.TableOrSubquery, err error) {
	logger.Tracef("Parse: Table OR Subquery")
	defer logger.Tracef("Parse: Table OR Subquery End")

	ts = &ast.TableOrSubquery{}
	if p.currentToken.Type == token.LEFTPAREN {
		p.readToken()
		if p.currentToken.Type == token.K_WITH || p.currentToken.Type == token.K_SELECT || p.currentToken.Type == token.K_VALUES {
			ss, e := p.parseSelectStatement()
			if e != nil {
				return ts, e
			}
			ts.SelectStatement = ss
			if p.currentToken.Type != token.RIGHTPAREN {
				return ts, errors.New("Parse Error: Invalid Token")
			}
			p.readToken()
			if p.currentToken.Type == token.K_AS {
				p.readToken()
				if p.currentToken.Type != token.IDENT {
					return ts, errors.New("Parse Error: Invalid Token")
				}
				ts.Alias = p.currentToken.Literal
			} else if p.currentToken.Type == token.IDENT {
				ts.Alias = p.currentToken.Literal
			}
		} else {
			var j *ast.JoinClause
			ts2, e := p.parseTableOrSubquery()
			if e != nil {
				return ts, e
			}
			ts.TableOrSubquery = ts2
			if p.currentToken.Type == token.COMMA || p.currentToken.Type == token.K_NATURAL || p.currentToken.Type == token.K_LEFT || p.currentToken.Type == token.K_INNER || p.currentToken.Type == token.K_CROSS || p.currentToken.Type == token.K_JOIN {
				j, e = p.parseJoin()
				if e != nil {
					return ts, e
				}
				j.LeftTableOrSubquery = ts2
			}
			if j != nil {
				ts.JoinClause = j
			}
		}
	} else {
		if p.currentToken.Type != token.IDENT {
			return ts, errors.New("Parse Error: Invalid Token")
		}
		if p.peekToken().Type == token.PERIOD {
			ts.Schema = p.currentToken.Literal
			p.readToken()
			p.readToken()
		}
		ts.TableName = p.currentToken.Literal
		p.readToken()
		if p.currentToken.Type == token.K_AS {
			p.readToken()
			if p.currentToken.Type != token.IDENT {
				return ts, errors.New("Parse Error: Invalid Token")
			}
			ts.Alias = p.currentToken.Literal
		} else if p.currentToken.Type == token.IDENT {
			ts.Alias = p.currentToken.Literal
		}
	}
	return
}

func (p *parser) parseJoin() (j *ast.JoinClause, err error) {
	logger.Tracef("Parse: JOIN")
	defer logger.Tracef("Parse: JOIN End")

	j = &ast.JoinClause{}
	if p.currentToken.Type == token.COMMA {
		j.Cross = true
		p.readToken()
	} else {
		if p.currentToken.Type == token.K_NATURAL {
			j.Natural = true
			p.readToken()
		}
		if p.currentToken.Type == token.K_LEFT {
			j.Left = true
			p.readToken()
			if p.currentToken.Type == token.K_OUTER {
				p.readToken()
			}
		} else if p.currentToken.Type == token.K_INNER {
			j.Inner = true
			p.readToken()
		} else if p.currentToken.Type == token.K_CROSS {
			j.Cross = true
			p.readToken()
		}
		p.readToken()
	}
	ss3, e := p.parseTableOrSubquery()
	if e != nil {
		return j, e
	}
	j.RightTableOrSubquery = ss3
	if p.currentToken.Type == token.K_ON {
		exp, e := p.parseExpression()
		if e != nil {
			return j, e
		}
		j.Expr = exp
	} else if p.currentToken.Type == token.K_USING {
		p.readToken()
		if p.currentToken.Type != token.LEFTPAREN {
			return j, errors.New("Parse Error: Invalid Token")
		}
		p.readToken()
		for {
			j.ColumnNames = append(j.ColumnNames, p.currentToken.Literal)
			p.readToken()
			if p.currentToken.Type != token.COMMA {
				break
			}
			p.readToken()
		}
	}
	return
}
