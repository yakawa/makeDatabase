package parser

import (
	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/token"
)

func (p *parser) parseExpression() (expr *ast.Expression, err error) {
	switch p.currentToken.Type {
	case token.K_CAST:
	case token.K_CASE:
	case token.K_NOT:
	case token.IDENT:
	default:
	}
	return
}

func (p *parser) parseIdent() (expr *ast.Expression, err error) {
	return
}

func (p *parser) parseNumber() (expr *ast.Expression, err error) {
	return
}

func (p *parser) parseBinaryExpr() (expr *ast.Expression, err error) {
	return
}
