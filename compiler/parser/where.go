package parser

import (
	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/logger"
)

func (p *parser) parseWhereClause() (w *ast.WhereClause, err error) {
	logger.Tracef("Parse: WHERE Clause")
	defer logger.Tracef("Parse: WHERE Clause End")

	w = &ast.WhereClause{}
	ex, er := p.parseExpression()
	if er != nil {
		return w, er
	}
	w.Expr = ex
	return
}
