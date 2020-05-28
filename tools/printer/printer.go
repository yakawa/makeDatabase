package printer

import (
	"bytes"
	"fmt"

	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/token"
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
	}
	if sc.IsDistinct {
		out.WriteString(fmt.Sprintf("%sDISTINCT: true\n", sp))
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
		if len(c.SchemaName) != 0 || len(c.TableName) != 0 {
			out.WriteString(fmt.Sprintf("%s- Column:\n", sp))
			if len(c.SchemaName) != 0 {
				out.WriteString(fmt.Sprintf("%s  Schema: %s\n", sp, c.SchemaName))
			}
			if len(c.TableName) != 0 {
				out.WriteString(fmt.Sprintf("%s   Table: %s\n", sp, c.TableName))
			}
		}

		if c.Expr != nil {
			out.WriteString(fmt.Sprintf("%s- Expression:\n", sp))
			out.WriteString(fmt.Sprintf("%s", printExpression(c.Expr, sp+"   ")))
			if len(c.Alias) != 0 {
				out.WriteString(fmt.Sprintf("%s   Alias: %s\n", sp, c.Alias))
			}
		}
	}
	return out.String()
}

func printOperator(t token.Type) string {
	switch t {
	case token.PLUSSIGN:
		return "+"
	case token.MINUSSIGN:
		return "-"
	case token.ASTERISK:
		return "\"*\""
	case token.SOLIDAS:
		return "/"
	default:
		return "Unknown"
	}
}

func printExpression(ex *ast.Expression, sp string) string {
	var out bytes.Buffer
	if ex.PrefixOpe != nil {
		out.WriteString(fmt.Sprintf("%sPrefixOpe:\n", sp))
		out.WriteString(fmt.Sprintf("%s Operator: %s\n", sp, printOperator(ex.PrefixOpe.Operator)))
		out.WriteString(fmt.Sprintf("%s Ope1:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printExpression(ex.PrefixOpe.Expr, sp+"  ")))
	}
	if ex.BinaryOpe != nil {
		out.WriteString(fmt.Sprintf("%sBinaryOpe:\n", sp))
		out.WriteString(fmt.Sprintf("%s Operator: %s\n", sp, printOperator(ex.BinaryOpe.Operator)))
		out.WriteString(fmt.Sprintf("%s Ope1:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printExpression(ex.BinaryOpe.Expr1, sp+"  ")))
		out.WriteString(fmt.Sprintf("%s Ope2:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printExpression(ex.BinaryOpe.Expr2, sp+"  ")))
	}
	if ex.Literal != nil {
		if ex.Literal.Numeric != nil {
			if ex.Literal.Numeric.SignedInt {
				out.WriteString(fmt.Sprintf("%sNumber: %d\n", sp, ex.Literal.Numeric.SI))
			} else if ex.Literal.Numeric.UnsignedInt {
				out.WriteString(fmt.Sprintf("%sNumber: %d\n", sp, ex.Literal.Numeric.UI))
			} else {
				out.WriteString(fmt.Sprintf("%sNumber: %f\n", sp, ex.Literal.Numeric.FL))
			}
		}
		if ex.Literal.IsString {
			out.WriteString(fmt.Sprintf("%sString: \"%s\"\n", sp, ex.Literal.String))
		}
		if ex.Literal.CurrentDate {
			out.WriteString(fmt.Sprintf("%sLiteral: CURRENT_DATE\n", sp))
		}
		if ex.Literal.CurrentTime {
			out.WriteString(fmt.Sprintf("%sLiteral: CURRENT_TIME\n", sp))
		}
		if ex.Literal.CurrentTimestamp {
			out.WriteString(fmt.Sprintf("%sLiteral: CURRENT_TIMESTAMP\n", sp))
		}
		if ex.Literal.True {
			out.WriteString(fmt.Sprintf("%sLiteral: TRUE\n", sp))
		}
		if ex.Literal.False {
			out.WriteString(fmt.Sprintf("%sLiteral: FALSE\n", sp))
		}
		if ex.Literal.Null {
			out.WriteString(fmt.Sprintf("%sLiteral: NULL\n", sp))
		}
	}
	if ex.ColumnName != nil {
		if len(ex.ColumnName.Schema) != 0 {
			out.WriteString(fmt.Sprintf("%sSchema: %s\n", sp, ex.ColumnName.Schema))
		}
		if len(ex.ColumnName.Table) != 0 {
			out.WriteString(fmt.Sprintf("%sTable: %s\n", sp, ex.ColumnName.Table))
		}
		if len(ex.ColumnName.Column) != 0 {
			out.WriteString(fmt.Sprintf("%sColumn: %s\n", sp, ex.ColumnName.Column))
		}
	}
	if ex.Function != nil {
		out.WriteString(fmt.Sprintf("%sFunction:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printFunctionExpression(ex.Function, sp+" ")))
	}
	return out.String()
}

func printFunctionExpression(f *ast.Function, sp string) string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%sName: %s\n", sp, f.Name))
	if f.Asterisk {
		out.WriteString(fmt.Sprintf("%sAsterisk: true\n", sp))
	}
	if f.Distinct {
		out.WriteString(fmt.Sprintf("%sDISTINCT: true\n", sp))
	}
	if len(f.Args) != 0 {
		out.WriteString(fmt.Sprintf("%sARGS:\n", sp))
		for _, a := range f.Args {
			out.WriteString(fmt.Sprintf("%s - arg:\n", sp))
			out.WriteString(fmt.Sprintf("%s", printExpression(&a, sp+"    ")))
		}
	}
	if f.FilterExpr != nil {
		out.WriteString(fmt.Sprintf("%sFilter:\n", sp))
		out.WriteString(fmt.Sprintf("%s - Expression:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printExpression(f.FilterExpr, sp+"    ")))
	}
	if f.OverClause != nil {
		out.WriteString(fmt.Sprintf("%sOver:\n", sp))
		if len(f.OverClause.WindowName) != 0 {
			out.WriteString(fmt.Sprintf("%s WindowName: %s\n", sp, f.OverClause.WindowName))
		} else {
			if len(f.OverClause.BaseWindowName) != 0 {
				out.WriteString(fmt.Sprintf("%s BaseWindowName: %s\n", sp, f.OverClause.BaseWindowName))
			}
			if len(f.OverClause.PartitionExpr) != 0 {
				out.WriteString(fmt.Sprintf("%s Partition:\n", sp))
				for _, p := range f.OverClause.PartitionExpr {
					out.WriteString(fmt.Sprintf("%s  - Expression:\n", sp))
					out.WriteString(fmt.Sprintf("%s", printExpression(&p, sp+"     ")))
				}
			}
			if len(f.OverClause.OrderBy) != 0 {
				out.WriteString(fmt.Sprintf("%s OrderBy:\n", sp))
				for _, o := range f.OverClause.OrderBy {
					out.WriteString(fmt.Sprintf("%s  - Order:\n", sp))
					out.WriteString(fmt.Sprintf("%s     Expression:\n", sp))
					out.WriteString(fmt.Sprintf("%s", printExpression(o.Expr, sp+"      ")))
					if len(o.CollateName) != 0 {
						out.WriteString(fmt.Sprintf("%s     Collation: %s\n", sp, o.CollateName))
					}
					if o.Asc {
						out.WriteString(fmt.Sprintf("%s     ASC: true\n", sp))
					}
					if o.Desc {
						out.WriteString(fmt.Sprintf("%s     DESC: true\n", sp))
					}
					if o.NullsFirst {
						out.WriteString(fmt.Sprintf("%s     NULLSFirst: true\n", sp))
					}
					if o.NullsLast {
						out.WriteString(fmt.Sprintf("%s     NULLSLast: true\n", sp))
					}
				}
			}
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
