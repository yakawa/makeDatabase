package parser

import (
	"testing"

	"github.com/yakawa/makeDatabase/common/errors"
	"github.com/yakawa/makeDatabase/compiler/lexer"
	"github.com/yakawa/makeDatabase/tools/printer"
)

func TestExpression(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"SELECT 1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Number: 1
`),
		},
		{"SELECT 1 a1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Number: 1
       Alias: a1
`),
		},
		{"SELECT 1 AS a1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Number: 1
       Alias: a1
`),
		},
		{"SELECT \"1\";",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       String: "1"
`),
		},
		{"SELECT CURRENT_TIME;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Literal: CURRENT_TIME
`),
		},
		{"SELECT CURRENT_DATE;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Literal: CURRENT_DATE
`),
		},
		{"SELECT CURRENT_TIMESTAMP;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Literal: CURRENT_TIMESTAMP
`),
		},
		{"SELECT TRUE;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Literal: TRUE
`),
		},
		{"SELECT FALSE;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Literal: FALSE
`),
		},
		{"SELECT NULL;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Literal: NULL
`),
		},
		{"SELECT -1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       PrefixOpe:
        Operator: -
        Ope1:
         Number: 1
`),
		},
		{"SELECT 1 + 1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       BinaryOpe:
        Operator: +
        Ope1:
         Number: 1
        Ope2:
         Number: 1
`),
		},
		{"SELECT 1 - 1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       BinaryOpe:
        Operator: -
        Ope1:
         Number: 1
        Ope2:
         Number: 1
`),
		},
		{"SELECT 1 * 1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       BinaryOpe:
        Operator: "*"
        Ope1:
         Number: 1
        Ope2:
         Number: 1
`),
		},
		{"SELECT 1 / 1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       BinaryOpe:
        Operator: /
        Ope1:
         Number: 1
        Ope2:
         Number: 1
`),
		},
		{"SELECT 1 * (2 + 3);",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       BinaryOpe:
        Operator: "*"
        Ope1:
         Number: 1
        Ope2:
         BinaryOpe:
          Operator: +
          Ope1:
           Number: 2
          Ope2:
           Number: 3
`),
		},
		{"SELECT 1 * ((2 + 3) / 4);",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       BinaryOpe:
        Operator: "*"
        Ope1:
         Number: 1
        Ope2:
         BinaryOpe:
          Operator: /
          Ope1:
           BinaryOpe:
            Operator: +
            Ope1:
             Number: 2
            Ope2:
             Number: 3
          Ope2:
           Number: 4
`),
		},
		{"SELECT c1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Column: c1
`),
		},
		{"SELECT t1.c1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Table: t1
       Column: c1
`),
		},
		{"SELECT d1.t1.c1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Schema: d1
       Table: t1
       Column: c1
`),
		},
		{"SELECT COUNT(1);",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Function:
        Name: COUNT
        ARGS:
         - arg:
            Number: 1
`),
		},
		{"SELECT COUNT(*);",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Function:
        Name: COUNT
        Asterisk: true
`),
		},
		{"SELECT SUM(c1, c2 * 2);",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Function:
        Name: SUM
        ARGS:
         - arg:
            Column: c1
         - arg:
            BinaryOpe:
             Operator: "*"
             Ope1:
              Column: c2
             Ope2:
              Number: 2
`),
		},
		{"SELECT SUM(DISTINCT c1);",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Function:
        Name: SUM
        DISTINCT: true
        ARGS:
         - arg:
            Column: c1
`),
		},
		{"SELECT SUM(DISTINCT c1);",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Function:
        Name: SUM
        DISTINCT: true
        ARGS:
         - arg:
            Column: c1
`),
		},
		{"SELECT SUM(c1) FILTER (WHERE c2 = 3) OVER(test PARTITION BY c4, c5 ORDER BY c6 DESC NULLS LAST ROWS BETWEEN UNBOUNDED PRECEDING AND c4 PRECEDING EXCLUDE CURRENT ROW);",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Function:
        Name: SUM
        ARGS:
         - arg:
            Column: c1
        Filter:
         - Expression:
            BinaryOpe:
             Operator: "="
             Ope1:
              Column: c2
             Ope2:
              Number: 3
        Over:
         BaseWindowName: test
         Partition:
          - Expression:
             Column: c4
          - Expression:
             Column: c5
         OrderBy:
          - Order:
             Expression:
              Column: c6
             DESC: true
         Frame:
          Rows: true
          UnboundedPreceding1: true
          ExprPreceding2:
           Column: c4
          ExcludeCurrentRow: true
`),
		},
		{"SELECT (c1, c2, c3);",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       ExpressionGroup:
        - Expression:
           Column: c1
        - Expression:
           Column: c2
        - Expression:
           Column: c3
`),
		},
		{"SELECT CAST (1 AS INTEGER);",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Cast:
        Type: INTEGER
        Expression:
         Number: 1
`),
		},
		{"SELECT c1 COLLATE utf8mb4;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Collate:
        Name: utf8mb4
        Expression:
         Column: c1
`),
		},
		{"SELECT c1 FROM t1 AS a1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Column: c1
  FROM:
   - ToS:
      Table: t1
      Alias: a1
`),
		},
		{"SELECT c1 FROM t1 AS a1 INNER JOIN t2 AS a2 ON a1.c1 = a2.c2;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Column: c1
  FROM:
   - ToS:
      Table: t1
      Alias: a1
   - ToS:
      Table: t2
      Alias: a2
      JOIN: Inner
      OnExpr:
       BinaryOpe:
        Operator: "="
        Ope1:
         Table: a1
         Column: c1
        Ope2:
         Table: a2
         Column: c2
`),
		},
		{"SELECT c1 FROM t1 INNER JOIN t2 ON a1.c1 = a2.c2;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Column: c1
  FROM:
   - ToS:
      Table: t1
   - ToS:
      Table: t2
      JOIN: Inner
      OnExpr:
       BinaryOpe:
        Operator: "="
        Ope1:
         Table: a1
         Column: c1
        Ope2:
         Table: a2
         Column: c2
`),
		},
		{"SELECT c1 FROM t1 LEFT OUTER JOIN t2 ON a1.c1 = a2.c2;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Column: c1
  FROM:
   - ToS:
      Table: t1
   - ToS:
      Table: t2
      JOIN: Left
      OnExpr:
       BinaryOpe:
        Operator: "="
        Ope1:
         Table: a1
         Column: c1
        Ope2:
         Table: a2
         Column: c2
`),
		},
		{"SELECT c1 FROM t1 LEFT JOIN t2 ON a1.c1 = a2.c2;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Column: c1
  FROM:
   - ToS:
      Table: t1
   - ToS:
      Table: t2
      JOIN: Left
      OnExpr:
       BinaryOpe:
        Operator: "="
        Ope1:
         Table: a1
         Column: c1
        Ope2:
         Table: a2
         Column: c2
`),
		},
		{"SELECT c1 FROM t1 RIGHT JOIN t2 ON a1.c1 = a2.c2;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Column: c1
  FROM:
   - ToS:
      Table: t1
   - ToS:
      Table: t2
      JOIN: Right
      OnExpr:
       BinaryOpe:
        Operator: "="
        Ope1:
         Table: a1
         Column: c1
        Ope2:
         Table: a2
         Column: c2
`),
		},
		{"SELECT c1 FROM t1 NATURAL JOIN t2 ON a1.c1 = a2.c2;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Column: c1
  FROM:
   - ToS:
      Table: t1
   - ToS:
      Table: t2
      JOIN: Natural
      OnExpr:
       BinaryOpe:
        Operator: "="
        Ope1:
         Table: a1
         Column: c1
        Ope2:
         Table: a2
         Column: c2
`),
		},
		{"SELECT c1 FROM t1, t2 USING(c1);",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Column: c1
  FROM:
   - ToS:
      Table: t1
   - ToS:
      Table: t2
      JOIN: Cross
      Using:
       - Column: c1
`),
		},
		{"SELECT c1 FROM t1, t2 USING(c1, c2);",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Column: c1
  FROM:
   - ToS:
      Table: t1
   - ToS:
      Table: t2
      JOIN: Cross
      Using:
       - Column: c1
       - Column: c2
`),
		},
		{"SELECT c1 FROM (SELECT c3, c4 FROM t1);",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Column: c1
  FROM:
   - ToS:
      SubQuery:
       SELECT:
        ResultColumns:
         - Expression:
            Column: c3
         - Expression:
            Column: c4
       FROM:
        - ToS:
           Table: t1
`),
		},

		{"SELECT c1 FROM t1 WHERE c1 > 3;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Column: c1
  FROM:
   - ToS:
      Table: t1
  WHERE:
   BinaryOpe:
    Operator: ">"
    Ope1:
     Column: c1
    Ope2:
     Number: 3
`),
		},
	}

	for i, tc := range testCases {
		l, _ := lexer.Tokenize(tc.input)
		sql, err := Parse(l)
		if err != nil {
			t.Errorf("[%d] Returned not nil %s", i, err)
			t.Log(tc.input)
			t.Log(err.(*errors.ErrParseInvalid).PrintStack(5))
		}
		y := printer.PrintSQL(sql)
		if y != tc.expected {
			t.Errorf("[%d] YAML Mistmatch", i)
			t.Log(tc.expected)
			t.Log(y)
		}
	}
}
