package parser

import (
	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/token"
	"github.com/yakawa/makeDatabase/logger"
)

func (p *parser) parseGroupByClause() (gb *ast.GroupByExpression, err error) {
	logger.Tracef("Parse: GroupBy Clause")
	defer logger.Tracef("Parse: GroupBy Clause End")

	gb = &ast.GroupByExpression{}

	for {
		ex, er := p.parseExpression()
		if er != nil {
			return gb, er
		}
		gb.GroupingExpr = append(gb.GroupingExpr, *ex)
		if p.currentToken.Type != token.COMMA {
			break
		}
		p.readToken()
	}
	logger.Infof("Ha: %#+v, %#+v", p.currentToken, p.peekToken())

	if p.currentToken.Type == token.K_HAVING {
		p.readToken()
		logger.Infof("Ha: %#+v, %#+v", p.currentToken, p.peekToken())

		ex, er := p.parseExpression()
		if er != nil {
			return gb, er
		}
		gb.HavingExpr = ex
	}

	return
}
