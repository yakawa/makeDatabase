package parser

import (
	"strconv"

	"github.com/yakawa/makeDatabase/common/ast"
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
	logger.Tracef("Parse: Expr %s", p.currentToken.Type.String())
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
	expr = &ast.Expression{}

	p.readToken()

	expr, err = p.parseExpr(LOWEST)
	p.readToken()
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
