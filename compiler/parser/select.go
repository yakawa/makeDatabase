package parser

import (
	"errors"

	"github.com/yakawa/makeDatabase/common/ast"
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
			//logger.Infof("%#+v", p.currentToken)
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
			p.readToken()
			f, e := p.parseFrom()
			if e != nil {
				return sc, e
			}
			sc.FromClause = f
		}
		if p.currentToken.Type == token.K_WHERE {
			//p.parseWhere()
		}
		if p.currentToken.Type == token.K_GROUP {
			//p.parseGroupBy()
		}
		if p.currentToken.Type == token.K_WINDOW {
			//p.parseWindow()
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
		p.readToken()
		return
	} else if p.currentToken.Type == token.IDENT && p.peekToken().Type == token.PERIOD {
		tmp := p.currentToken
		p.readToken()
		p.readToken()
		if p.currentToken.Type == token.ASTERISK {
			rc.Asterisk = true
			rc.TableName = tmp.Literal
			rc.Asterisk = true
			p.readToken()
			return
		} else if p.currentToken.Type == token.IDENT {
			if p.peekToken().Type == token.PERIOD {
				schema := tmp.Literal
				table := p.currentToken.Literal
				p.readToken()
				p.readToken()
				if p.currentToken.Type == token.ASTERISK {
					rc.SchemaName = schema
					rc.TableName = table
					rc.Asterisk = true
					p.readToken()
					return
				} else if p.currentToken.Type == token.IDENT {
					p.rewindToken()
					p.rewindToken()
					p.rewindToken()
					p.rewindToken()
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
			return rc, errors.New("Parse Error: Invalid Token")
		}
		rc.Alias = p.currentToken.Literal
		p.readToken()
	} else if p.currentToken.Type == token.IDENT {
		rc.Alias = p.currentToken.Literal
		p.readToken()
	}
	return
}
