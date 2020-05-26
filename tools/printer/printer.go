package printer

import (
	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/logger"
)

func PrintAST(a *ast.SQL) {
	logger.Infof("AST Printer")
	if a.SelectStatement != nil {
		printSelectStatement(a.SelectStatement, 0)
	}
}

func printSelectStatement(ss *ast.SelectStatement, n int) {
	sp := ""
	for i := 0; i < n; i++ {
		sp += "  "
	}
	if ss.WithClause != nil {
		logger.Infof(sp + "WITH:")
	}
	if ss.SelectClause != nil {
		logger.Infof(sp+"SELECT: %#+v", ss)
		printSelectClause(ss.SelectClause, n+1)
	}
}

func printSelectClause(sc *ast.SelectClause, n int) {
	sp := ""
	for i := 0; i < n; i++ {
		sp += "  "
	}
	if sc.IsAll == true {
		logger.Infof(sp + "ALL: true")
	}
	if sc.IsDistinct == true {
		logger.Infof(sp + "DISTINCT: true")
	}
	if len(sc.ResultColumns) != 0 {
		logger.Infof(sp + "Result Columns:")
	}
}
