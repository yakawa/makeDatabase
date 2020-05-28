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
   ALL: false
   DISTINCT: false
   ResultColumns:
    - Expression:
       Number: 1
`),
		},
		{"SELECT 1 a1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ALL: false
   DISTINCT: false
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
   ALL: false
   DISTINCT: false
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
   ALL: false
   DISTINCT: false
   ResultColumns:
    - Expression:
       String: "1"
`),
		},
		{"SELECT CURRENT_TIME;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ALL: false
   DISTINCT: false
   ResultColumns:
    - Expression:
       Literal: CURRENT_TIME
`),
		},
		{"SELECT CURRENT_DATE;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ALL: false
   DISTINCT: false
   ResultColumns:
    - Expression:
       Literal: CURRENT_DATE
`),
		},
		{"SELECT CURRENT_TIMESTAMP;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ALL: false
   DISTINCT: false
   ResultColumns:
    - Expression:
       Literal: CURRENT_TIMESTAMP
`),
		},
		{"SELECT TRUE;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ALL: false
   DISTINCT: false
   ResultColumns:
    - Expression:
       Literal: TRUE
`),
		},
		{"SELECT FALSE;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ALL: false
   DISTINCT: false
   ResultColumns:
    - Expression:
       Literal: FALSE
`),
		},
		{"SELECT NULL;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ALL: false
   DISTINCT: false
   ResultColumns:
    - Expression:
       Literal: NULL
`),
		},
		{"SELECT -1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ALL: false
   DISTINCT: false
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
   ALL: false
   DISTINCT: false
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
   ALL: false
   DISTINCT: false
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
   ALL: false
   DISTINCT: false
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
   ALL: false
   DISTINCT: false
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
   ALL: false
   DISTINCT: false
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
   ALL: false
   DISTINCT: false
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
   ALL: false
   DISTINCT: false
   ResultColumns:
    - Expression:
       Column: c1
`),
		},
		{"SELECT t1.c1;",
			string(`SQL:
 SELECTStatement:
  SELECT:
   ALL: false
   DISTINCT: false
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
   ALL: false
   DISTINCT: false
   ResultColumns:
    - Expression:
       Schema: d1
       Table: t1
       Column: c1
`),
		},
		{"SELECT SUM();",
			string(`SQL:
				 SELECTStatement:
				  SELECT:
				   ALL: false
				   DISTINCT: false
				   ResultColumns:
				    - Expression:
				       Schema: d1
				       Table: t1
				       Column: c1
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
