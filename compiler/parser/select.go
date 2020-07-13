package parser

import (
	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/errors"
	"github.com/yakawa/makeDatabase/common/token"
	"github.com/yakawa/makeDatabase/logger"
)

func (p *parser) parseSelectClause() (sc *ast.SelectClause, err error) {
	logger.Tracef("Parse: SELECT Clause")
	defer logger.Tracef("Parse: SELECT Clause End")

	sc = &ast.SelectClause{}
	if p.currentToken.Type == token.K_SELECT {
		p.readToken()
		if p.currentToken.Type == token.K_DISTINCT || p.currentToken.Type == token.K_ALL {
			if p.currentToken.Type == token.K_DISTINCT {
				sc.IsDistinct = true
			} else {
				sc.IsAll = true
			}
			p.readToken()
		}
		for {
			r, e := p.parseResultColumn()
			if e != nil {
				return sc, e
			}
			sc.ResultColumns = append(sc.ResultColumns, r)

			if p.currentToken.Type != token.COMMA {
				break
			}
			p.readToken()
		}

		if p.currentToken.Type == token.EOS || p.currentToken.Type == token.SEMICOLON {
			return
		}
		if p.currentToken.Type == token.K_FROM {
			f, e := p.parseFromClause()
			if e != nil {
				return sc, e
			}
			sc.FromClause = f
		}

		if p.currentToken.Type == token.K_WHERE {
			p.readToken()
			w, e := p.parseWhereClause()
			if e != nil {
				return sc, e
			}
			sc.WhereClause = w
		}
		if p.currentToken.Type == token.K_GROUP {
			p.readToken()
			if p.currentToken.Type != token.K_BY {
				return sc, errors.NewErrParseInvalidToken(p.currentToken)
			}
			p.readToken()
			g, e := p.parseGroupByClause()
			if e != nil {
				return sc, e
			}
			sc.GroupByExpression = g
		}
		if p.currentToken.Type == token.K_WINDOW {
			p.readToken()
			w, e := p.parseWindowClause()
			if e != nil {
				return sc, e
			}
			sc.WindowExpression = w
		}
	}
	return
}

func (p *parser) parseResultColumn() (rc ast.ResultColumn, err error) {
	logger.Tracef("Parse: Result Columns")
	defer logger.Tracef("Parse: Result Columns End")

	rc = ast.ResultColumn{}

	if p.currentToken.Type == token.ASTERISK {
		rc.Asterisk = true
		return
	} else if p.currentToken.Type == token.IDENT && p.peekToken().Type == token.PERIOD {
		// Schema.Database.Table.Column
		tmp1 := p.currentToken.Literal
		p.readToken()
		p.readToken()
		if p.currentToken.Type == token.ASTERISK {
			rc.TableName = tmp1
			rc.Asterisk = true
			return
		} else if p.currentToken.Type == token.IDENT {
			if p.peekToken().Type == token.PERIOD {
				tmp2 := p.currentToken.Literal
				p.readToken()
				p.readToken()
				if p.currentToken.Type == token.ASTERISK {
					rc.DatabaseName = tmp1
					rc.TableName = tmp2
					rc.Asterisk = true
					return
				} else if p.currentToken.Type == token.IDENT {
					if p.peekToken().Type == token.PERIOD {
						tmp3 := p.currentToken.Literal
						p.readToken()
						p.readToken()
						if p.currentToken.Type == token.ASTERISK {
							rc.SchemaName = tmp1
							rc.DatabaseName = tmp2
							rc.TableName = tmp3
							rc.Asterisk = true
							return
						} else if p.currentToken.Type == token.IDENT {
						}
					}
				}
			} else {
				p.rewindToken()
				p.rewindToken()
			}
		}
	}

	expr, e := p.parseExpression()
	if e != nil {
		return rc, e
	}
	rc.Expr = expr

	if p.currentToken.Type == token.K_AS {
		p.readToken()
		if p.currentToken.Type != token.IDENT {
			return rc, errors.NewErrParseInvalidToken(p.currentToken)
		}
		rc.Alias = p.currentToken.Literal
		p.readToken()
	} else if p.currentToken.Type == token.IDENT {
		rc.Alias = p.currentToken.Literal
		p.readToken()
	}
	return
}
