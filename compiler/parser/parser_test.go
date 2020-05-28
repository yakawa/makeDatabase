package parser

import (
	"testing"

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
             Operator: =
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
	}

	for i, tc := range testCases {
		l, _ := lexer.Tokenize(tc.input)
		sql, err := Parse(l)
		if err != nil {
			t.Errorf("[%d] Returned not nil %s", i, err)
		}
		y := printer.PrintSQL(sql)
		if y != tc.expected {
			t.Errorf("[%d] YAML Mistmatch", i)
			t.Log(tc.expected)
			t.Log(y)
		}
	}
}
