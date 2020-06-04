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
		if ss.WithClause.IsRecursive {
			out.WriteString(fmt.Sprintf("%s Recursive: true\n", sp))
		}
		out.WriteString(fmt.Sprintf("%s CTEs:\n", sp))
		for _, cte := range ss.WithClause.CTE {
			out.WriteString(fmt.Sprintf("%s  - CTE:\n", sp))
			out.WriteString(fmt.Sprintf("%s     Name: %s\n", sp, cte.TableName))
			if len(cte.ColumnNames) != 0 {
				out.WriteString(fmt.Sprintf("%s     Columns:\n", sp))
				for _, c := range cte.ColumnNames {
					out.WriteString(fmt.Sprintf("%s      - %s\n", sp, c))
				}
			}
			out.WriteString(fmt.Sprintf("%s", printSelectStatement(cte.SelectStatement, sp+"     ")))
		}
	}
	if ss.SelectClause != nil {
		out.WriteString(fmt.Sprintf("%sSELECT:\n", sp))
		if ss.SelectClause != nil {
			out.WriteString(fmt.Sprintf("%s", printSelectClause(ss.SelectClause, sp+" ")))
			if ss.SelectClause.FromClause != nil {
				out.WriteString(fmt.Sprintf("%sFROM:\n", sp))
				out.WriteString(fmt.Sprintf("%s", printFromClause(ss.SelectClause.FromClause, sp+" ")))
			}
			if ss.SelectClause.WhereClause != nil {
				out.WriteString(fmt.Sprintf("%sWHERE:\n", sp))
				out.WriteString(fmt.Sprintf("%s", printExpression(ss.SelectClause.WhereClause.Expr, sp+" ")))
			}
			if ss.SelectClause.GroupByExpression != nil {
				out.WriteString(fmt.Sprintf("%sGROUPBy:\n", sp))
				out.WriteString(fmt.Sprintf("%s Grouping:\n", sp))
				for _, e := range ss.SelectClause.GroupByExpression.GroupingExpr {
					out.WriteString(fmt.Sprintf("%s  - Expression:\n", sp))
					out.WriteString(fmt.Sprintf("%s", printExpression(&e, sp+"     ")))
				}
				if ss.SelectClause.GroupByExpression.HavingExpr != nil {
					out.WriteString(fmt.Sprintf("%s Condition:\n", sp))
					out.WriteString(fmt.Sprintf("%s  - Expression:\n", sp))
					out.WriteString(fmt.Sprintf("%s", printExpression(ss.SelectClause.GroupByExpression.HavingExpr, sp+"     ")))
				}
			}
			if ss.SelectClause.WindowExpression != nil {
				out.WriteString(fmt.Sprintf("%s", printWindowClause(ss.SelectClause.WindowExpression, sp)))
			}
		}
	}
	if ss.ValuesClause != nil {
		out.WriteString(fmt.Sprintf("%sVALUES:\n", sp))
		for _, v := range ss.ValuesClause.Expr {
			out.WriteString(fmt.Sprintf("%s - Expression:\n", sp))
			out.WriteString(fmt.Sprintf("%s", printExpression(&v, sp+"    ")))
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
				out.WriteString(fmt.Sprintf("%s   Schema: %s\n", sp, c.SchemaName))
			}
			if len(c.TableName) != 0 {
				out.WriteString(fmt.Sprintf("%s   Table: %s\n", sp, c.TableName))
			}
			if c.Asterisk {
				out.WriteString(fmt.Sprintf("%s   Asterisk: true\n", sp))
			}
		}
		if len(c.SchemaName) == 0 && len(c.TableName) == 0 && c.Asterisk {
			out.WriteString(fmt.Sprintf("%s- Column:\n", sp))
			out.WriteString(fmt.Sprintf("%s   Asterisk: true\n", sp))
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
	case token.EQUALS:
		return "\"=\""
	case token.NOTEQUALS:
		return "<>"
	case token.K_AND:
		return "AND"
	case token.K_OR:
		return "OR"
	case token.K_NOT:
		return "NOT"
	case token.CONCAT:
		return "||"
	case token.LESSTHAN:
		return "\"<\""
	case token.LESSTHANEQUALS:
		return "\"<=\""
	case token.GREATERTHAN:
		return "\">\""
	case token.GREATERTHANEQUALS:
		return "\">=\""
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
	if len(ex.Expr) != 0 {
		out.WriteString(fmt.Sprintf("%sExpressionGroup:\n", sp))
		for _, exp := range ex.Expr {
			out.WriteString(fmt.Sprintf("%s - Expression:\n", sp))
			out.WriteString(fmt.Sprintf("%s", printExpression(&exp, sp+"    ")))
		}
	}
	if ex.Cast != nil {
		out.WriteString(fmt.Sprintf("%sCast:\n", sp))
		out.WriteString(fmt.Sprintf("%s Type: %s\n", sp, ex.Cast.TypeName))
		out.WriteString(fmt.Sprintf("%s Expression:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printExpression(ex.Cast.Expr, sp+"  ")))
		if ex.Cast.IsN1 {
			out.WriteString(fmt.Sprintf("%s N1: %d\n", sp, ex.Cast.N1))
		}
		if ex.Cast.IsN2 {
			out.WriteString(fmt.Sprintf("%s N2: %d\n", sp, ex.Cast.N2))
		}
	}
	if ex.Collate != nil {
		out.WriteString(fmt.Sprintf("%sCollate:\n", sp))
		out.WriteString(fmt.Sprintf("%s Name: %s\n", sp, ex.Collate.Name))
		out.WriteString(fmt.Sprintf("%s Expression:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printExpression(ex.Collate.Expr, sp+"  ")))
	}
	if ex.String != nil {
		out.WriteString(fmt.Sprintf("%sStringOpe:\n", sp))
		if ex.String.Glob {
			out.WriteString(fmt.Sprintf("%s GLOB: true\n", sp))
		}
		if ex.String.Like {
			out.WriteString(fmt.Sprintf("%s LIKE: true\n", sp))
		}
		if ex.String.Match {
			out.WriteString(fmt.Sprintf("%s MATCH: true\n", sp))
		}
		if ex.String.Regexp {
			out.WriteString(fmt.Sprintf("%s REGEXP: true\n", sp))
		}
		if ex.String.Not {
			out.WriteString(fmt.Sprintf("%s Not: true\n", sp))
		}
		out.WriteString(fmt.Sprintf("%s Expression1:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printExpression(ex.String.Expr1, sp+"  ")))
		out.WriteString(fmt.Sprintf("%s Expression2:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printExpression(ex.String.Expr2, sp+"  ")))
		if ex.String.EscapeExpr != nil {
			out.WriteString(fmt.Sprintf("%s EscapeExpression:\n", sp))
			out.WriteString(fmt.Sprintf("%s", printExpression(ex.String.EscapeExpr, sp+"  ")))
		}
	}
	if ex.Null != nil {
		out.WriteString(fmt.Sprintf("%sNullOpe:\n", sp))
		if ex.Null.NotNull {
			out.WriteString(fmt.Sprintf("%s Not: true\n", sp))
		}
		out.WriteString(fmt.Sprintf("%s Expression:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printExpression(ex.Null.Expr, sp+"  ")))
	}
	if ex.Between != nil {
		out.WriteString(fmt.Sprintf("%sBetweenOpe:\n", sp))
		if ex.Between.Not {
			out.WriteString(fmt.Sprintf("%s Not: true\n", sp))
		}
		out.WriteString(fmt.Sprintf("%s Expression1:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printExpression(ex.Between.Expr1, sp+"  ")))
		out.WriteString(fmt.Sprintf("%s Expression2:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printExpression(ex.Between.Expr2, sp+"  ")))
		out.WriteString(fmt.Sprintf("%s Expression3:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printExpression(ex.Between.Expr3, sp+"  ")))
	}
	if ex.In != nil {
		out.WriteString(fmt.Sprintf("%sInOpe:\n", sp))
		if ex.In.Not {
			out.WriteString(fmt.Sprintf("%s Not: true\n", sp))
		}
		out.WriteString(fmt.Sprintf("%s Expression:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printExpression(ex.In.Expr, sp+"  ")))
		if len(ex.In.Schema) != 0 {
		}
		if len(ex.In.Schema) != 0 {
			out.WriteString(fmt.Sprintf("%s Schema: %s\n", sp, ex.In.Schema))
		}
		if len(ex.In.Table) != 0 {
			out.WriteString(fmt.Sprintf("%s Table: %s\n", sp, ex.In.Table))
		}
		if ex.In.SelectStatement != nil {
			out.WriteString(fmt.Sprintf("%s SelectStatement:\n", sp))
			out.WriteString(fmt.Sprintf("%s", printSelectStatement(ex.In.SelectStatement, sp+"   ")))
		}
		if len(ex.In.InExpr) != 0 {
			out.WriteString(fmt.Sprintf("%s Expressions:\n", sp))
			for _, exp := range ex.In.InExpr {
				out.WriteString(fmt.Sprintf("%s  - Expression:\n", sp))
				out.WriteString(fmt.Sprintf("%s", printExpression(&exp, sp+"     ")))
			}
		}
	}
	if ex.Exists != nil {
		out.WriteString(fmt.Sprintf("%sExistsOpe:\n", sp))
		if ex.Exists.Not {
			out.WriteString(fmt.Sprintf("%s Not: true\n", sp))
		}
		out.WriteString(fmt.Sprintf("%s SelectStatement:\n", sp))
		out.WriteString(fmt.Sprintf("%s", printSelectStatement(ex.Exists.SelectStatement, sp+"   ")))
	}
	if ex.Case != nil {
		out.WriteString(fmt.Sprintf("%sCaseOpe:\n", sp))
		if ex.Case.CaseExpr != nil {
			out.WriteString(fmt.Sprintf("%s CaseExpression:\n", sp))
			out.WriteString(fmt.Sprintf("%s", printExpression(ex.Case.CaseExpr, sp+"   ")))
		}
		if len(ex.Case.WhenThen) != 0 {
			out.WriteString(fmt.Sprintf("%s WhenExpression:\n", sp))
			for _, wh := range ex.Case.WhenThen {
				out.WriteString(fmt.Sprintf("%s  - When:\n", sp))
				out.WriteString(fmt.Sprintf("%s", printExpression(wh.WhenExpr, sp+"     ")))
				out.WriteString(fmt.Sprintf("%s    Then:\n", sp))
				out.WriteString(fmt.Sprintf("%s", printExpression(wh.ThenExpr, sp+"     ")))
			}
		}
		if ex.Case.ElseExpr != nil {
			out.WriteString(fmt.Sprintf("%s ElseExpression:\n", sp))
			out.WriteString(fmt.Sprintf("%s", printExpression(ex.Case.ElseExpr, sp+"   ")))
		}
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
			if f.OverClause.FrameSpec != nil {
				out.WriteString(fmt.Sprintf("%s FrameSpec:\n", sp))
				if f.OverClause.FrameSpec.Range {
					out.WriteString(fmt.Sprintf("%s  Range: true\n", sp))
				}
				if f.OverClause.FrameSpec.Rows {
					out.WriteString(fmt.Sprintf("%s  Rows: true\n", sp))
				}
				if f.OverClause.FrameSpec.Groups {
					out.WriteString(fmt.Sprintf("%s  Groups: true\n", sp))
				}
				if f.OverClause.FrameSpec.UnboundedPreceding {
					out.WriteString(fmt.Sprintf("%s  UnboundedPreceding: true\n", sp))
				}
				if f.OverClause.FrameSpec.CurrentRow {
					out.WriteString(fmt.Sprintf("%s  CurrentRow: true\n", sp))
				}
				if f.OverClause.FrameSpec.ExprPreceding != nil {
					out.WriteString(fmt.Sprintf("%s  ExprPreceding:\n", sp))
					out.WriteString(fmt.Sprintf("%s", printExpression(f.OverClause.FrameSpec.ExprPreceding, sp+"   ")))
				}
				if f.OverClause.FrameSpec.UnboundedPreceding1 || f.OverClause.FrameSpec.CurrentRow1 || f.OverClause.FrameSpec.ExprPreceding1 != nil || f.OverClause.FrameSpec.ExprFollowing1 != nil {
					out.WriteString(fmt.Sprintf("%s  BetweenBefore:\n", sp))

					if f.OverClause.FrameSpec.UnboundedPreceding1 {
						out.WriteString(fmt.Sprintf("%s   UnboundedPreceding: true\n", sp))
					}
					if f.OverClause.FrameSpec.CurrentRow1 {
						out.WriteString(fmt.Sprintf("%s   CurrentRow: true\n", sp))
					}
					if f.OverClause.FrameSpec.ExprPreceding1 != nil {
						out.WriteString(fmt.Sprintf("%s   ExprPreceding:\n", sp))
						out.WriteString(fmt.Sprintf("%s", printExpression(f.OverClause.FrameSpec.ExprPreceding1, sp+"    ")))
					}
					if f.OverClause.FrameSpec.ExprFollowing1 != nil {
						out.WriteString(fmt.Sprintf("%s   ExprFollowing:\n", sp))
						out.WriteString(fmt.Sprintf("%s", printExpression(f.OverClause.FrameSpec.ExprFollowing1, sp+"    ")))
					}
				}
				if f.OverClause.FrameSpec.UnboundedFollowing2 || f.OverClause.FrameSpec.CurrentRow2 || f.OverClause.FrameSpec.ExprPreceding2 != nil || f.OverClause.FrameSpec.ExprFollowing2 != nil {
					out.WriteString(fmt.Sprintf("%s  BetweenAfter:\n", sp))

					if f.OverClause.FrameSpec.UnboundedFollowing2 {
						out.WriteString(fmt.Sprintf("%s   UnboundedFollowing: true\n", sp))
					}
					if f.OverClause.FrameSpec.CurrentRow2 {
						out.WriteString(fmt.Sprintf("%s   CurrentRow: true\n", sp))
					}
					if f.OverClause.FrameSpec.ExprPreceding2 != nil {
						out.WriteString(fmt.Sprintf("%s   ExprPreceding:\n", sp))
						out.WriteString(fmt.Sprintf("%s", printExpression(f.OverClause.FrameSpec.ExprPreceding2, sp+"    ")))
					}
					if f.OverClause.FrameSpec.ExprFollowing2 != nil {
						out.WriteString(fmt.Sprintf("%s   ExprFollowing:\n", sp))
						out.WriteString(fmt.Sprintf("%s", printExpression(f.OverClause.FrameSpec.ExprFollowing2, sp+"    ")))
					}
				}

				if f.OverClause.FrameSpec.ExcludeNoOthers {
					out.WriteString(fmt.Sprintf("%s  ExcludeNoOthers: true\n", sp))
				}
				if f.OverClause.FrameSpec.ExcludeCurrentRow {
					out.WriteString(fmt.Sprintf("%s  ExcludeCurrentRow: true\n", sp))
				}
				if f.OverClause.FrameSpec.ExcludeGroup {
					out.WriteString(fmt.Sprintf("%s  ExcludeGroup: true\n", sp))
				}
				if f.OverClause.FrameSpec.ExcludeTies {
					out.WriteString(fmt.Sprintf("%s  ExcludedTies: true\n", sp))
				}
			}
		}
	}
	return out.String()
}

func printFromClause(fr *ast.FromClause, sp string) string {
	var out bytes.Buffer
	if len(fr.ToS) != 0 {
		for _, ts := range fr.ToS {
			out.WriteString(fmt.Sprintf("%s- ToS:\n", sp))
			if len(ts.Schema) != 0 {
				out.WriteString(fmt.Sprintf("%s   Schema: %s\n", sp, ts.Schema))
			}
			if len(ts.TableName) != 0 {
				out.WriteString(fmt.Sprintf("%s   Table: %s\n", sp, ts.TableName))
			}

			if ts.Subquery != nil {
				out.WriteString(fmt.Sprintf("%s   SubQuery:\n", sp))
				out.WriteString(fmt.Sprintf("%s", printSelectStatement(ts.Subquery, sp+"    ")))
			}

			if len(ts.Alias) != 0 {
				out.WriteString(fmt.Sprintf("%s   Alias: %s\n", sp, ts.Alias))
			}

			if ts.Natural {
				out.WriteString(fmt.Sprintf("%s   JOIN: Natural\n", sp))
			}
			if ts.Left {
				out.WriteString(fmt.Sprintf("%s   JOIN: Left\n", sp))
			}
			if ts.Right {
				out.WriteString(fmt.Sprintf("%s   JOIN: Right\n", sp))
			}
			if ts.Inner {
				out.WriteString(fmt.Sprintf("%s   JOIN: Inner\n", sp))
			}
			if ts.Cross {
				out.WriteString(fmt.Sprintf("%s   JOIN: Cross\n", sp))
			}

			if ts.On != nil {
				out.WriteString(fmt.Sprintf("%s   OnExpr:\n", sp))
				out.WriteString(fmt.Sprintf("%s", printExpression(ts.On, sp+"    ")))
			}

			if len(ts.ColumnNames) != 0 {
				out.WriteString(fmt.Sprintf("%s   Using:\n", sp))
				for _, c := range ts.ColumnNames {
					out.WriteString(fmt.Sprintf("%s    - Column: %s\n", sp, c))
				}
			}
		}
	}
	return out.String()
}

func printWindowClause(w *ast.WindowExpression, sp string) string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%sWindow:\n", sp))
	out.WriteString(fmt.Sprintf("%s Definition:\n", sp))
	for _, d := range w.Defn {
		out.WriteString(fmt.Sprintf("%s  - Function:\n", sp))
		out.WriteString(fmt.Sprintf("%s     Name: %s\n", sp, d.Name))
		if len(d.BaseWindowName) != 0 {
			out.WriteString(fmt.Sprintf("%s     BaseWindowName: %s\n", sp, d.BaseWindowName))
		}
		if len(d.PartitionExpr) != 0 {
			out.WriteString(fmt.Sprintf("%s     Partition:\n", sp))
			for _, p := range d.PartitionExpr {
				out.WriteString(fmt.Sprintf("%s      - Expression:\n", sp))
				out.WriteString(fmt.Sprintf("%s", printExpression(&p, sp+"         ")))
			}
		}

		if len(d.OrderExpr) != 0 {
			out.WriteString(fmt.Sprintf("%s     OrderBy:\n", sp))
			for _, o := range d.OrderExpr {
				out.WriteString(fmt.Sprintf("%s      - Order:\n", sp))
				out.WriteString(fmt.Sprintf("%s         Expression:\n", sp))
				out.WriteString(fmt.Sprintf("%s", printExpression(o.Expr, sp+"          ")))
				if len(o.CollateName) != 0 {
					out.WriteString(fmt.Sprintf("%s        Collation: %s\n", sp, o.CollateName))
				}
				if o.Asc {
					out.WriteString(fmt.Sprintf("%s       ASC: true\n", sp))
				}
				if o.Desc {
					out.WriteString(fmt.Sprintf("%s       DESC: true\n", sp))
				}
				if o.NullsFirst {
					out.WriteString(fmt.Sprintf("%s       NULLSFirst: true\n", sp))
				}
				if o.NullsLast {
					out.WriteString(fmt.Sprintf("%s       NULLSLast: true\n", sp))
				}
			}
		}
		if d.Frame != nil {
			out.WriteString(fmt.Sprintf("%s     FrameSpec:\n", sp))
			if d.Frame.Range {
				out.WriteString(fmt.Sprintf("%s      Range: true\n", sp))
			}
			if d.Frame.Rows {
				out.WriteString(fmt.Sprintf("%s      Rows: true\n", sp))
			}
			if d.Frame.Rows {
				out.WriteString(fmt.Sprintf("%s      Groups: true\n", sp))
			}
			if d.Frame.UnboundedPreceding {
				out.WriteString(fmt.Sprintf("%s      UnboundedPreceding: true\n", sp))
			}
			if d.Frame.CurrentRow {
				out.WriteString(fmt.Sprintf("%s      CurrentRow: true\n", sp))
			}
			if d.Frame.ExprPreceding != nil {
				out.WriteString(fmt.Sprintf("%s      Preceding:\n", sp))
				out.WriteString(fmt.Sprintf("%s", printExpression(d.Frame.ExprPreceding, sp+"       ")))
			}
			if d.Frame.UnboundedPreceding1 || d.Frame.CurrentRow1 || d.Frame.ExprPreceding1 != nil || d.Frame.ExprFollowing1 != nil {
				out.WriteString(fmt.Sprintf("%s      BetweenBefore:\n", sp))
				if d.Frame.UnboundedPreceding1 {
					out.WriteString(fmt.Sprintf("%s       UnboundedPreceding: true\n", sp))
				}
				if d.Frame.CurrentRow1 {
					out.WriteString(fmt.Sprintf("%s       CurrentRow: true\n", sp))
				}
				if d.Frame.ExprPreceding1 != nil {
					out.WriteString(fmt.Sprintf("%s       ExprPreceding:\n", sp))
					out.WriteString(fmt.Sprintf("%s", printExpression(d.Frame.ExprPreceding1, sp+"        ")))
				}
				if d.Frame.ExprFollowing1 != nil {
					out.WriteString(fmt.Sprintf("%s       ExprFollowing:\n", sp))
					out.WriteString(fmt.Sprintf("%s", printExpression(d.Frame.ExprFollowing1, sp+"        ")))
				}
			}
			if d.Frame.UnboundedFollowing2 || d.Frame.CurrentRow2 || d.Frame.ExprPreceding2 != nil || d.Frame.ExprFollowing2 != nil {
				out.WriteString(fmt.Sprintf("%s      BetweenAfter:\n", sp))
				if d.Frame.UnboundedFollowing2 {
					out.WriteString(fmt.Sprintf("%s       UnboundedFollowing: true\n", sp))
				}
				if d.Frame.CurrentRow2 {
					out.WriteString(fmt.Sprintf("%s       CurrentRow: true\n", sp))
				}
				if d.Frame.ExprPreceding2 != nil {
					out.WriteString(fmt.Sprintf("%s       ExprPreceding:\n", sp))
					out.WriteString(fmt.Sprintf("%s", printExpression(d.Frame.ExprPreceding2, sp+"        ")))
				}
				if d.Frame.ExprFollowing2 != nil {
					out.WriteString(fmt.Sprintf("%s       ExprFollowing:\n", sp))
					out.WriteString(fmt.Sprintf("%s", printExpression(d.Frame.ExprFollowing2, sp+"        ")))
				}
			}
			if d.Frame.ExcludeNoOthers {
				out.WriteString(fmt.Sprintf("%s      ExcludeNoOthers: true\n", sp))
			}
			if d.Frame.ExcludeCurrentRow {
				out.WriteString(fmt.Sprintf("%s      ExcludeCurrentRow: true\n", sp))
			}
			if d.Frame.ExcludeGroup {
				out.WriteString(fmt.Sprintf("%s      ExcludeGroup: true\n", sp))
			}
			if d.Frame.ExcludeTies {
				out.WriteString(fmt.Sprintf("%s      ExcludeTies: true\n", sp))
			}

		}
	}
	return out.String()
}
