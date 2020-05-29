package parser

import (
	"errors"

	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/token"
	"github.com/yakawa/makeDatabase/logger"
)

func (p *parser) parseValuesClause() (v *ast.ValuesClause, err error) {
	logger.Tracef("Parse: Values Clause")
	defer logger.Tracef("Parse: Values Clause End")

	v = &ast.ValuesClause{}

	p.readToken()
	if p.currentToken.Type != token.LEFTPAREN {
		return v, errors.New("Parse Error Invalid Token")
	}
	p.readToken()

	for {
		ex, er := p.parseExpr(LOWEST)
		if er != nil {
			return v, er
		}
		v.Expr = append(v.Expr, *ex)

		p.readToken()
		if p.currentToken.Type != token.COMMA {
			break
		}
		p.readToken()
	}

	return
}
