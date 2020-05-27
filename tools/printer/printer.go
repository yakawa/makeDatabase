package printer

import (
	"bytes"
	"fmt"

	"github.com/yakawa/makeDatabase/common/ast"
)

func PrintSQL(a *ast.SQL) string {
	var out bytes.Buffer
	sp := " "
	out.WriteString(fmt.Sprintf("SQL:\n"))
	if a.SelectStatement != nil {
		out.WriteString(fmt.Sprintf("%sSELECTStatement:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printSelectStatement(a.SelectStatement, sp+" ")))
	}
	return out.String()
}

func printSelectStatement(ss *ast.SelectStatement, sp string) string {
	var out bytes.Buffer
	if ss.WithClause != nil {
		out.WriteString(fmt.Sprintf("%sWITH:\n", sp))
	}
	if ss.SelectClause != nil {
		out.WriteString(fmt.Sprintf("%sSELECT:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printSelectClause(ss.SelectClause, sp+" ")))
		if ss.SelectClause.FromClause != nil {
			out.WriteString(fmt.Sprintf("%sFROM:\n", sp))
			out.WriteString(fmt.Sprintf("%s\n", printFromClause(ss.SelectClause.FromClause, sp+" ")))
		}
	}
	return out.String()
}

func printSelectClause(sc *ast.SelectClause, sp string) string {
	var out bytes.Buffer
	if sc.IsAll {
		out.WriteString(fmt.Sprintf("%sALL: true\n", sp))
	} else {
		out.WriteString(fmt.Sprintf("%sALL: false\n", sp))
	}
	if sc.IsDistinct {
		out.WriteString(fmt.Sprintf("%sDISTINCT: true\n", sp))
	} else {
		out.WriteString(fmt.Sprintf("%sDISTINCT: false\n", sp))
	}
	if len(sc.ResultColumns) != 0 {
		out.WriteString(fmt.Sprintf("%sResultColumns:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printResultColumns(sc.ResultColumns, sp+" ")))
	}
	return out.String()
}

func printResultColumns(rc []ast.ResultColumn, sp string) string {
	var out bytes.Buffer
	for _, c := range rc {
		if len(c.SchemaName) != 0 {
			out.WriteString(fmt.Sprintf("%s- Schema: %s\n", sp, c.SchemaName))
		}
		if len(c.TableName) != 0 {
			out.WriteString(fmt.Sprintf("%s  Table: %s\n", sp, c.TableName))
		}
		if c.Asterisk == true {
			out.WriteString(fmt.Sprintf("%s  Column: *\n", sp))
		}

		if c.Expr != nil {
			out.WriteString(fmt.Sprintf("%s  Expression: Expr\n", sp))
		}
		if len(c.Alias) != 0 {
			out.WriteString(fmt.Sprintf("%s  Alias: %s\n", sp, c.Alias))
		}
	}
	return out.String()
}

func printFromClause(fr *ast.FromClause, sp string) string {
	var out bytes.Buffer
	if len(fr.ToS.Schema) != 0 {
		out.WriteString(fmt.Sprintf("%sSchema: %s\n", sp, fr.ToS.Schema))
	}
	if len(fr.ToS.TableName) != 0 {
		out.WriteString(fmt.Sprintf("%sTable: %s\n", sp, fr.ToS.TableName))
	}
	if len(fr.ToS.Alias) != 0 {
		out.WriteString(fmt.Sprintf("%sAlias: %s\n", sp, fr.ToS.Alias))
	}
	if fr.ToS.JoinClause != nil {
		out.WriteString(fmt.Sprintf("%sJOIN:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printJoin(fr.ToS.JoinClause, sp+" ")))
	}
	return out.String()
}

func printJoin(j *ast.JoinClause, sp string) string {
	var out bytes.Buffer
	if j.Natural {
		out.WriteString(fmt.Sprintf("%s- NATRUAL: true\n", sp))
	} else {
		out.WriteString(fmt.Sprintf("%s- NATURAL: false\n", sp))
	}
	if j.Left {
		out.WriteString(fmt.Sprintf("%s  LEFT: true\n", sp))
	} else {
		out.WriteString(fmt.Sprintf("%s  LEFT: false\n", sp))
	}
	if j.Right {
		out.WriteString(fmt.Sprintf("%s  RIGHT: true\n", sp))
	} else {
		out.WriteString(fmt.Sprintf("%s  RIGHT: false\n", sp))
	}
	if j.Inner {
		out.WriteString(fmt.Sprintf("%s  INNER: true\n", sp))
	} else {
		out.WriteString(fmt.Sprintf("%s  INNER: false\n", sp))
	}
	if j.Cross {
		out.WriteString(fmt.Sprintf("%s  CROSS: true\n", sp))
	} else {
		out.WriteString(fmt.Sprintf("%s  CROSS: false\n", sp))
	}

	return out.String()
}
