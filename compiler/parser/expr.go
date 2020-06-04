package parser

import (
	"strconv"

	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/errors"
	"github.com/yakawa/makeDatabase/common/token"
	"github.com/yakawa/makeDatabase/logger"
)

func (p *parser) parseExpression() (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Expression")
	defer logger.Tracef("Parse: Expression End")

	expr, err = p.parseExpr(LOWEST)
	p.readToken()

	return
}

func (p *parser) parseExpr(pre int) (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Expr")
	defer logger.Tracef("Parse: Expr End")

	unary := p.unaryParseFunction[p.currentToken.Type]
	left, e := unary()
	if e != nil {
		return expr, e
	}

	for p.peekToken().Type != token.EOS && pre < p.peekPrecedence() {
		binary := p.binaryParseFunction[p.peekToken().Type]
		if binary == nil {
			return left, nil
		}
		p.readToken()
		left, err = binary(left)
	}
	return left, nil
}

func (p *parser) parseIdent() (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Ident")
	defer logger.Tracef("Parse: Ident End")

	expr = &ast.Expression{}
	if !(p.peekToken().Type == token.PERIOD || p.peekToken().Type == token.LEFTPAREN) {
		c := &ast.ColumnName{}
		c.Column = p.currentToken.Literal
		expr.ColumnName = c
		return
	}
	if p.peekToken().Type == token.PERIOD {
		c := &ast.ColumnName{}
		tmp := p.currentToken
		p.readToken()
		p.readToken()
		if p.peekToken().Type == token.PERIOD {
			c.Schema = tmp.Literal
			c.Table = p.currentToken.Literal
			p.readToken()
			p.readToken()
			c.Column = p.currentToken.Literal
			expr.ColumnName = c
			return
		} else {
			c.Table = tmp.Literal
			c.Column = p.currentToken.Literal
			expr.ColumnName = c
			return
		}
	}
	if p.peekToken().Type == token.LEFTPAREN {
		e, er := p.parseFunctionExpr()
		if er != nil {
			return expr, er
		}
		expr = e
	}
	return
}

func (p *parser) parseLiteral() (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Literal")
	defer logger.Tracef("Parse: Literal End")

	expr = &ast.Expression{}
	expr.Literal = &ast.Literal{}

	switch p.currentToken.Type {
	case token.K_CURRENT_DATE:
		expr.Literal.CurrentDate = true
	case token.K_CURRENT_TIME:
		expr.Literal.CurrentTime = true
	case token.K_CURRENT_TIMESTAMP:
		expr.Literal.CurrentTimestamp = true
	case token.STRING:
		expr.Literal.IsString = true
		expr.Literal.String = p.currentToken.Literal
	case token.K_NULL:
		expr.Literal.Null = true
	case token.K_TRUE:
		expr.Literal.True = true
	case token.K_FALSE:
		expr.Literal.False = true
	}
	return
}

func (p *parser) parseNumber() (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Number")
	defer logger.Tracef("Parse: Number End")

	expr = &ast.Expression{}
	expr.Literal = &ast.Literal{}
	n := &ast.Numeric{}
	isPeriod := false
	for _, c := range p.currentToken.Literal {
		if c == '.' {
			isPeriod = true
			break
		}
	}
	if isPeriod {
		i, e := strconv.ParseFloat(p.currentToken.Literal, 64)
		if e != nil {
			return expr, e
		}
		n.Float = true
		n.FL = i
	} else {
		i, e := strconv.ParseInt(p.currentToken.Literal, 10, 32)
		if e != nil {
			return expr, e
		}
		n.SignedInt = true
		n.SI = int(i)
	}
	expr.Literal.Numeric = n
	return
}

func (p *parser) parseGroupedExpr() (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Grouped")
	defer logger.Tracef("Parse: Grouped End")
	p.readToken()

	e, er := p.parseExpr(LOWEST)
	if er != nil {
		return expr, er
	}
	p.readToken()
	if p.currentToken.Type != token.COMMA {
		return e, nil
	}
	expr = &ast.Expression{}
	expr.Expr = append(expr.Expr, *e)
	for {
		p.readToken()
		e, er := p.parseExpr(LOWEST)
		if er != nil {
			return expr, er
		}
		expr.Expr = append(expr.Expr, *e)
		p.readToken()
		if p.currentToken.Type != token.COMMA {
			break
		}
	}
	return
}

func (p *parser) parsePrefixExpr() (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Prefix")
	defer logger.Tracef("Parse: Prefix End")

	expr = &ast.Expression{}
	u := &ast.UnaryOpe{}
	u.Operator = p.currentToken.Type

	p.readToken()

	u.Expr, err = p.parseExpr(PREFIX)
	if err != nil {
		return
	}

	expr.PrefixOpe = u
	return
}

func (p *parser) parseBinaryExpr(left *ast.Expression) (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Binary")
	defer logger.Tracef("Parse: Binary End")

	expr = &ast.Expression{}
	b := &ast.BinaryOpe{}
	b.Expr1 = left
	b.Operator = p.currentToken.Type
	pre := p.getCurrentPrecedence()

	p.readToken()

	r, e := p.parseExpr(pre)
	if e != nil {
		return expr, e
	}
	b.Expr2 = r

	expr.BinaryOpe = b
	return
}

func (p *parser) parseFunctionExpr() (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Function")
	defer logger.Tracef("Parse: Function End")

	expr = &ast.Expression{}
	f := &ast.Function{}
	f.Name = p.currentToken.Literal
	p.readToken()
	p.readToken()

	if p.currentToken.Type == token.ASTERISK {
		f.Asterisk = true
		p.readToken()
	} else {
		if p.currentToken.Type == token.K_DISTINCT {
			f.Distinct = true
			p.readToken()
		}
		for p.currentToken.Type != token.RIGHTPAREN {
			e, er := p.parseExpr(LOWEST)
			if er != nil {
				return expr, er
			}
			f.Args = append(f.Args, *e)
			p.readToken()
			if p.currentToken.Type == token.COMMA {
				p.readToken()
			}
		}
	}
	if p.currentToken.Type != token.RIGHTPAREN {
		return expr, errors.NewErrParseInvalidToken(p.currentToken)
	}

	if p.peekToken().Type == token.K_FILTER {
		p.readToken()
		p.readToken()
		if p.currentToken.Type != token.LEFTPAREN {
			return expr, errors.NewErrParseInvalidToken(p.currentToken)
		}
		p.readToken()
		if p.currentToken.Type != token.K_WHERE {
			return expr, errors.NewErrParseInvalidToken(p.currentToken)
		}
		p.readToken()
		e, er := p.parseExpr(LOWEST)
		if er != nil {
			return expr, er
		}
		f.FilterExpr = e
		p.readToken()
		if p.currentToken.Type != token.RIGHTPAREN {
			return expr, errors.NewErrParseInvalidToken(p.currentToken)
		}
	}

	if p.peekToken().Type == token.K_OVER {
		p.readToken()
		o := &ast.OverClause{}
		p.readToken()
		if p.currentToken.Type == token.IDENT {
			o.WindowName = p.currentToken.Literal
		} else {
			if p.currentToken.Type != token.LEFTPAREN {
				return expr, errors.NewErrParseInvalidToken(p.currentToken)
			}
			p.readToken()

			d, er := p.parseWindowDefinition()
			if er != nil {
				return expr, er
			}
			o.BaseWindowName = d.BaseWindowName
			o.PartitionExpr = d.PartitionExpr
			o.OrderBy = d.OrderExpr
			o.FrameSpec = d.Frame
		}
		f.OverClause = o
	}
	expr.Function = f
	return
}

func (p *parser) parseCastExpr() (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Cast")
	defer logger.Tracef("Parse: Cast End")

	expr = &ast.Expression{}
	c := &ast.CastOpe{}

	p.readToken()
	if p.currentToken.Type != token.LEFTPAREN {
		return expr, errors.NewErrParseInvalidToken(p.currentToken)
	}
	p.readToken()
	e, er := p.parseExpr(LOWEST)
	if er != nil {
		return expr, er
	}
	c.Expr = e
	p.readToken()
	if p.currentToken.Type != token.K_AS {
		return expr, errors.NewErrParseInvalidToken(p.currentToken)
	}
	p.readToken()
	c.TypeName = p.currentToken.Literal
	p.readToken()
	if p.currentToken.Type == token.LEFTPAREN {
		p.readToken()
		n := 1
		if p.currentToken.Type == token.MINUSSIGN || p.currentToken.Type == token.PLUSSIGN {
			if p.currentToken.Type == token.MINUSSIGN {
				n *= -1
			}
			p.readToken()
		}
		if p.currentToken.Type != token.NUMBER {
			return expr, errors.NewErrParseInvalidToken(p.currentToken)
		}
		nn, er := strconv.ParseInt(p.currentToken.Literal, 10, 32)
		if er != nil {
			return expr, errors.NewErrParseInvalidToken(p.currentToken)
		}
		c.N1 = n * int(nn)
		c.IsN1 = true
		p.readToken()

		if p.currentToken.Type == token.COMMA {
			p.readToken()
			n := 1
			if p.currentToken.Type == token.MINUSSIGN || p.currentToken.Type == token.PLUSSIGN {
				if p.currentToken.Type == token.MINUSSIGN {
					n *= -1
				}
				p.readToken()
			}
			if p.currentToken.Type != token.NUMBER {
				return expr, errors.NewErrParseInvalidToken(p.currentToken)
			}
			nn, er := strconv.ParseInt(p.currentToken.Literal, 10, 32)
			if er != nil {
				return expr, errors.NewErrParseInvalidToken(p.currentToken)
			}
			c.N2 = n * int(nn)
			c.IsN2 = true
			p.readToken()
		}
		if p.currentToken.Type != token.RIGHTPAREN {
			return expr, errors.NewErrParseInvalidToken(p.currentToken)
		}
		p.readToken()
	}
	if p.currentToken.Type != token.RIGHTPAREN {
		return expr, errors.NewErrParseInvalidToken(p.currentToken)
	}
	expr.Cast = c
	return
}

func (p *parser) parseExistsExpr() (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Exists")
	defer logger.Tracef("Parse: Exists End")

	r := &ast.ExistsOpe{}

	if p.currentToken.Type == token.K_NOT {
		r.Not = true
		p.readToken()
	}

	p.readToken()
	if p.currentToken.Type != token.LEFTPAREN {
		return expr, errors.NewErrParseInvalidToken(p.currentToken)
	}
	p.readToken()

	if !(p.currentToken.Type == token.K_WITH || p.currentToken.Type == token.K_SELECT || p.currentToken.Type == token.K_VALUES) {
		return expr, errors.NewErrParseInvalidToken(p.currentToken)
	}

	ss, er := p.parseSelectStatement()
	if er != nil {
		return expr, er
	}
	r.SelectStatement = ss
	if p.currentToken.Type != token.RIGHTPAREN {
		return expr, errors.NewErrParseInvalidToken(p.currentToken)
	}

	expr = &ast.Expression{}
	expr.Exists = r
	return
}

func (p *parser) parseCaseExpr() (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Case")
	defer logger.Tracef("Parse: Case End")

	p.readToken()

	r := &ast.CaseOpe{}

	if p.currentToken.Type != token.K_WHEN {
		ex, er := p.parseExpr(LOWEST)
		if er != nil {
			return expr, er
		}
		r.CaseExpr = ex
		p.readToken()
	}

	for {
		w := &ast.WhenThen{}
		if p.currentToken.Type != token.K_WHEN {
			return expr, errors.NewErrParseInvalidToken(p.currentToken)
		}
		p.readToken()
		ex, er := p.parseExpr(LOWEST)
		if er != nil {
			return expr, er
		}
		w.WhenExpr = ex

		p.readToken()

		if p.currentToken.Type != token.K_THEN {
			return expr, errors.NewErrParseInvalidToken(p.currentToken)
		}
		p.readToken()

		ex, er = p.parseExpr(LOWEST)
		if er != nil {
			return expr, er
		}
		w.ThenExpr = ex
		r.WhenThen = append(r.WhenThen, *w)
		p.readToken()
		if p.currentToken.Type != token.K_WHEN {
			break
		}
	}
	if p.currentToken.Type == token.K_ELSE {
		p.readToken()
		ex, er := p.parseExpr(LOWEST)
		if er != nil {
			return expr, er
		}
		r.ElseExpr = ex
		p.readToken()
	}
	if p.currentToken.Type != token.K_END {
		return expr, errors.NewErrParseInvalidToken(p.currentToken)
	}

	expr = &ast.Expression{}
	expr.Case = r
	return
}

func (p *parser) parseCollateExpr(left *ast.Expression) (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Collate")
	defer logger.Tracef("Parse: Collate End")

	expr = &ast.Expression{}

	c := &ast.CollateOpe{}
	c.Expr = left
	p.readToken()
	c.Name = p.currentToken.Literal

	expr.Collate = c
	return
}

func (p *parser) parseStringFunc(left *ast.Expression) (expr *ast.Expression, err error) {
	logger.Tracef("Parse: StringFunc")
	defer logger.Tracef("Parse: StringFunc End")

	expr = &ast.Expression{}

	r := &ast.StringOpe{}
	r.Expr1 = left
	if p.currentToken.Type == token.K_LIKE {
		r.Like = true
	}
	if p.currentToken.Type == token.K_GLOB {
		r.Glob = true
	}
	if p.currentToken.Type == token.K_REGEXP {
		r.Regexp = true
	}
	if p.currentToken.Type == token.K_MATCH {
		r.Match = true
	}
	p.readToken()
	ex, er := p.parseExpr(LOWEST)
	if er != nil {
		return expr, er
	}
	r.Expr2 = ex
	if p.peekToken().Type == token.K_ESCAPE {
		p.readToken()
		p.readToken()
		ex, er := p.parseExpr(LOWEST)
		if er != nil {
			return expr, er
		}
		r.EscapeExpr = ex
	}
	expr.String = r

	return
}

func (p *parser) parseNullExpr(left *ast.Expression) (expr *ast.Expression, err error) {
	logger.Tracef("Parse: NULL")
	defer logger.Tracef("Parse: NULL End")

	expr = &ast.Expression{}

	r := &ast.NullOpe{}
	r.Expr = left

	if p.currentToken.Type == token.K_NOTNULL {
		r.NotNull = true
	}
	if p.currentToken.Type == token.K_ISNULL {
		r.IsNull = true
	}

	expr.Null = r
	return
}

func (p *parser) parseIsExpr(left *ast.Expression) (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Is")
	defer logger.Tracef("Parse: Is End")

	expr = &ast.Expression{}

	r := &ast.BinaryOpe{}
	r.Expr1 = left
	p.readToken()
	if p.currentToken.Type == token.K_NOT {
		r.Operator = token.NOTEQUALS
		p.readToken()
	} else {
		r.Operator = token.EQUALS
	}
	ex, er := p.parseExpr(LOWEST)
	if er != nil {
		return expr, er
	}
	r.Expr2 = ex

	expr.BinaryOpe = r
	return
}

func (p *parser) parseBetweenExpr(left *ast.Expression) (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Between")
	defer logger.Tracef("Parse: Between End")

	r := &ast.BetweenOpe{}
	r.Expr1 = left
	p.readToken()
	ex, er := p.parseExpr(LOWEST)
	if er != nil {
		return expr, er
	}
	r.Expr2 = ex
	p.readToken()
	if p.currentToken.Type != token.K_AND {
		return expr, errors.NewErrParseInvalidToken(p.currentToken)
	}
	p.readToken()
	ex, er = p.parseExpr(LOWEST)
	if er != nil {
		return expr, er
	}
	r.Expr3 = ex
	expr = &ast.Expression{}
	expr.Between = r
	return
}

func (p *parser) parseInExpr(left *ast.Expression) (expr *ast.Expression, err error) {
	logger.Tracef("Parse: In")
	defer logger.Tracef("Parse: In End")

	r := &ast.InOpe{}
	r.Expr = left
	p.readToken()
	if p.currentToken.Type == token.LEFTPAREN {
		p.readToken()
		if p.currentToken.Type == token.K_WITH || p.currentToken.Type == token.K_SELECT || p.currentToken.Type == token.K_VALUES {
			ss, er := p.parseSelectStatement()
			if er != nil {
				return expr, er
			}
			r.SelectStatement = ss
		} else {
			for {
				ex, er := p.parseExpr(LOWEST)
				if er != nil {
					return expr, er
				}
				r.InExpr = append(r.InExpr, *ex)
				p.readToken()
				if p.currentToken.Type != token.COMMA {
					break
				}
			}
		}
	} else if p.currentToken.Type == token.IDENT {
		if p.peekToken().Type == token.PERIOD {
			r.Schema = p.currentToken.Literal
			p.readToken()
			p.readToken()
			if p.currentToken.Type != token.IDENT {
				return expr, errors.NewErrParseInvalidToken(p.currentToken)
			}
		}
		r.Table = p.currentToken.Literal
	} else {
		return expr, errors.NewErrParseInvalidToken(p.currentToken)
	}
	expr = &ast.Expression{}

	expr.In = r
	return
}

func (p *parser) parseNotExpr(left *ast.Expression) (expr *ast.Expression, err error) {
	logger.Tracef("Parse: Not")
	defer logger.Tracef("Parse: Not End")

	p.readToken()
	switch p.currentToken.Type {
	case token.K_LIKE, token.K_GLOB, token.K_REGEXP, token.K_MATCH:
		ex, er := p.parseStringFunc(left)
		if er != nil {
			return expr, er
		}
		ex.String.Not = true
		expr = ex
	case token.K_NULL:
		expr = &ast.Expression{}
		r := &ast.NullOpe{}
		r.NotNull = true
		r.Expr = left
		expr.Null = r
	case token.K_BETWEEN:
		ex, er := p.parseBetweenExpr(left)
		if er != nil {
			return expr, er
		}
		ex.Between.Not = true
		expr = ex
	case token.K_IN:
		ex, er := p.parseInExpr(left)
		if er != nil {
			return expr, er
		}
		ex.In.Not = true
		expr = ex
	}

	p.readToken()
	return
}
