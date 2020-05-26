package lexer

import (
	"testing"

	"github.com/yakawa/makeDatabase/common/token"
)

func TestTokenize(t *testing.T) {
	testCases := []struct {
		input    string
		expected []token.Token
	}{
		{
			"\"",
			[]token.Token{
				{
					Type:    token.DOUBLEQUOTE,
					Literal: "\"",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"%",
			[]token.Token{
				{
					Type:    token.PERCENT,
					Literal: "%",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"&",
			[]token.Token{
				{
					Type:    token.AMPERSAND,
					Literal: "&",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"'",
			[]token.Token{
				{
					Type:    token.QUOTE,
					Literal: "'",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"(",
			[]token.Token{
				{
					Type:    token.LEFTPAREN,
					Literal: "(",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			")",
			[]token.Token{
				{
					Type:    token.RIGHTPAREN,
					Literal: ")",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"*",
			[]token.Token{
				{
					Type:    token.ASTERISK,
					Literal: "*",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"+",
			[]token.Token{
				{
					Type:    token.PLUSSIGN,
					Literal: "+",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			",",
			[]token.Token{
				{
					Type:    token.COMMA,
					Literal: ",",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"-",
			[]token.Token{
				{
					Type:    token.MINUSSIGN,
					Literal: "-",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			".",
			[]token.Token{
				{
					Type:    token.PERIOD,
					Literal: ".",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"/",
			[]token.Token{
				{
					Type:    token.SOLIDAS,
					Literal: "/",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			":",
			[]token.Token{
				{
					Type:    token.COLON,
					Literal: ":",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			";",
			[]token.Token{
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"<",
			[]token.Token{
				{
					Type:    token.LESSTHAN,
					Literal: "<",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"=",
			[]token.Token{
				{
					Type:    token.EQUALS,
					Literal: "=",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			">",
			[]token.Token{
				{
					Type:    token.GREATERTHAN,
					Literal: ">",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"?",
			[]token.Token{
				{
					Type:    token.QUESTION,
					Literal: "?",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"[",
			[]token.Token{
				{
					Type:    token.LEFTBRACKET,
					Literal: "[",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"]",
			[]token.Token{
				{
					Type:    token.RIGHTBRACKET,
					Literal: "]",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"^",
			[]token.Token{
				{
					Type:    token.CIRCUMFLEX,
					Literal: "^",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"_",
			[]token.Token{
				{
					Type:    token.UNDERSCORE,
					Literal: "_",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"|",
			[]token.Token{
				{
					Type:    token.VERTICALBAR,
					Literal: "|",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"{",
			[]token.Token{
				{
					Type:    token.LEFTBRACE,
					Literal: "{",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"}",
			[]token.Token{
				{
					Type:    token.RIGHTBRACE,
					Literal: "}",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"<>",
			[]token.Token{
				{
					Type:    token.NOTEQUALS,
					Literal: "<>",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			">=",
			[]token.Token{
				{
					Type:    token.GREATERTHANEQUALS,
					Literal: ">=",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"<=",
			[]token.Token{
				{
					Type:    token.LESSTHANEQUALS,
					Literal: "<=",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"||",
			[]token.Token{
				{
					Type:    token.CONCAT,
					Literal: "||",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"'Comment'",
			[]token.Token{
				{
					Type:    token.STRING,
					Literal: "Comment",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"\"Comment\"",
			[]token.Token{
				{
					Type:    token.STRING,
					Literal: "Comment",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"'Comment'''",
			[]token.Token{
				{
					Type:    token.STRING,
					Literal: "Comment'",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"'Comment\"\"'",
			[]token.Token{
				{
					Type:    token.STRING,
					Literal: "Comment\"\"",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"'Comment';",
			[]token.Token{
				{
					Type:    token.STRING,
					Literal: "Comment",
				},
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"\"Comment\";",
			[]token.Token{
				{
					Type:    token.STRING,
					Literal: "Comment",
				},
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"1",
			[]token.Token{
				{
					Type:    token.NUMBER,
					Literal: "1",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"12",
			[]token.Token{
				{
					Type:    token.NUMBER,
					Literal: "12",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"12.",
			[]token.Token{
				{
					Type:    token.NUMBER,
					Literal: "12.",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"12.1",
			[]token.Token{
				{
					Type:    token.NUMBER,
					Literal: "12.1",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			".1",
			[]token.Token{
				{
					Type:    token.NUMBER,
					Literal: ".1",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"12;",
			[]token.Token{
				{
					Type:    token.NUMBER,
					Literal: "12",
				},
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"4.2.1",
			[]token.Token{
				{
					Type:    token.NUMBER,
					Literal: "4.2",
				},
				{
					Type:    token.NUMBER,
					Literal: ".1",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"abc",
			[]token.Token{
				{
					Type:    token.IDENT,
					Literal: "abc",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"abc def",
			[]token.Token{
				{
					Type:    token.IDENT,
					Literal: "abc",
				},
				{
					Type:    token.IDENT,
					Literal: "def",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"abc\tdef",
			[]token.Token{
				{
					Type:    token.IDENT,
					Literal: "abc",
				},
				{
					Type:    token.IDENT,
					Literal: "def",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"abc\ndef",
			[]token.Token{
				{
					Type:    token.IDENT,
					Literal: "abc",
				},
				{
					Type:    token.IDENT,
					Literal: "def",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"'abc def'",
			[]token.Token{
				{
					Type:    token.STRING,
					Literal: "abc def",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"4e+2;",
			[]token.Token{
				{
					Type:    token.NUMBER,
					Literal: "4",
				},
				{
					Type:    token.IDENT,
					Literal: "e",
				},
				{
					Type:    token.PLUSSIGN,
					Literal: "+",
				},
				{
					Type:    token.NUMBER,
					Literal: "2",
				},
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"WITH",
			[]token.Token{
				{
					Type:    token.K_WITH,
					Literal: "WITH",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"with",
			[]token.Token{
				{
					Type:    token.K_WITH,
					Literal: "with",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"With",
			[]token.Token{
				{
					Type:    token.K_WITH,
					Literal: "With",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"ALL",
			[]token.Token{
				{
					Type:    token.K_ALL,
					Literal: "ALL",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"AND",
			[]token.Token{
				{
					Type:    token.K_AND,
					Literal: "AND",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"AS",
			[]token.Token{
				{
					Type:    token.K_AS,
					Literal: "AS",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"ASC",
			[]token.Token{
				{
					Type:    token.K_ASC,
					Literal: "ASC",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"BETWEEN",
			[]token.Token{
				{
					Type:    token.K_BETWEEN,
					Literal: "BETWEEN",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"BY",
			[]token.Token{
				{
					Type:    token.K_BY,
					Literal: "BY",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"CASE",
			[]token.Token{
				{
					Type:    token.K_CASE,
					Literal: "CASE",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"CAST",
			[]token.Token{
				{
					Type:    token.K_CAST,
					Literal: "CAST",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"COLLATE",
			[]token.Token{
				{
					Type:    token.K_COLLATE,
					Literal: "COLLATE",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"CROSS",
			[]token.Token{
				{
					Type:    token.K_CROSS,
					Literal: "CROSS",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"CURRENT",
			[]token.Token{
				{
					Type:    token.K_CURRENT,
					Literal: "CURRENT",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"CURRENT_DATE",
			[]token.Token{
				{
					Type:    token.K_CURRENT_DATE,
					Literal: "CURRENT_DATE",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"CURRENT_TIME",
			[]token.Token{
				{
					Type:    token.K_CURRENT_TIME,
					Literal: "CURRENT_TIME",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"CURRENT_TIMESTAMP",
			[]token.Token{
				{
					Type:    token.K_CURRENT_TIMESTAMP,
					Literal: "CURRENT_TIMESTAMP",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"DESC",
			[]token.Token{
				{
					Type:    token.K_DESC,
					Literal: "DESC",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"DISTINCT",
			[]token.Token{
				{
					Type:    token.K_DISTINCT,
					Literal: "DISTINCT",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"ELSE",
			[]token.Token{
				{
					Type:    token.K_ELSE,
					Literal: "ELSE",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"END",
			[]token.Token{
				{
					Type:    token.K_END,
					Literal: "END",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"EXCLUDE",
			[]token.Token{
				{
					Type:    token.K_EXCLUDE,
					Literal: "EXCLUDE",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"ESCAPE",
			[]token.Token{
				{
					Type:    token.K_ESCAPE,
					Literal: "ESCAPE",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"EXCEPT",
			[]token.Token{
				{
					Type:    token.K_EXCEPT,
					Literal: "EXCEPT",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"EXISTS",
			[]token.Token{
				{
					Type:    token.K_EXISTS,
					Literal: "EXISTS",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"FALSE",
			[]token.Token{
				{
					Type:    token.K_FALSE,
					Literal: "FALSE",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"FIRST",
			[]token.Token{
				{
					Type:    token.K_FIRST,
					Literal: "FIRST",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"FOLLOWING",
			[]token.Token{
				{
					Type:    token.K_FOLLOWING,
					Literal: "FOLLOWING",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"FROM",
			[]token.Token{
				{
					Type:    token.K_FROM,
					Literal: "FROM",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"GLOB",
			[]token.Token{
				{
					Type:    token.K_GLOB,
					Literal: "GLOB",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"GROUP",
			[]token.Token{
				{
					Type:    token.K_GROUP,
					Literal: "GROUP",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"GROUPS",
			[]token.Token{
				{
					Type:    token.K_GROUPS,
					Literal: "GROUPS",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"HAVING",
			[]token.Token{
				{
					Type:    token.K_HAVING,
					Literal: "HAVING",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"INTERSECT",
			[]token.Token{
				{
					Type:    token.K_INTERSECT,
					Literal: "INTERSECT",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"IN",
			[]token.Token{
				{
					Type:    token.K_IN,
					Literal: "IN",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"INDEXED",
			[]token.Token{
				{
					Type:    token.K_INDEXED,
					Literal: "INDEXED",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"INNER",
			[]token.Token{
				{
					Type:    token.K_INNER,
					Literal: "INNER",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"IS",
			[]token.Token{
				{
					Type:    token.K_IS,
					Literal: "IS",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"ISNULL",
			[]token.Token{
				{
					Type:    token.K_ISNULL,
					Literal: "ISNULL",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"JOIN",
			[]token.Token{
				{
					Type:    token.K_JOIN,
					Literal: "JOIN",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"LAST",
			[]token.Token{
				{
					Type:    token.K_LAST,
					Literal: "LAST",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"LEFT",
			[]token.Token{
				{
					Type:    token.K_LEFT,
					Literal: "LEFT",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"LIKE",
			[]token.Token{
				{
					Type:    token.K_LIKE,
					Literal: "LIKE",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"LIMIT",
			[]token.Token{
				{
					Type:    token.K_LIMIT,
					Literal: "LIMIT",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"MATCH",
			[]token.Token{
				{
					Type:    token.K_MATCH,
					Literal: "MATCH",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"NATURAL",
			[]token.Token{
				{
					Type:    token.K_NATURAL,
					Literal: "NATURAL",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"NO",
			[]token.Token{
				{
					Type:    token.K_NO,
					Literal: "NO",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"NOT",
			[]token.Token{
				{
					Type:    token.K_NOT,
					Literal: "NOT",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"NOTNULL",
			[]token.Token{
				{
					Type:    token.K_NOTNULL,
					Literal: "NOTNULL",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"NULL",
			[]token.Token{
				{
					Type:    token.K_NULL,
					Literal: "NULL",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"NULLS",
			[]token.Token{
				{
					Type:    token.K_NULLS,
					Literal: "NULLS",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"OFFSET",
			[]token.Token{
				{
					Type:    token.K_OFFSET,
					Literal: "OFFSET",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"ON",
			[]token.Token{
				{
					Type:    token.K_ON,
					Literal: "ON",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"ORDER",
			[]token.Token{
				{
					Type:    token.K_ORDER,
					Literal: "ORDER",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"OTHERS",
			[]token.Token{
				{
					Type:    token.K_OTHERS,
					Literal: "OTHERS",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"OUTER",
			[]token.Token{
				{
					Type:    token.K_OUTER,
					Literal: "OUTER",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"OVER",
			[]token.Token{
				{
					Type:    token.K_OVER,
					Literal: "OVER",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"PARTITION",
			[]token.Token{
				{
					Type:    token.K_PARTITION,
					Literal: "PARTITION",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"PRECEDING",
			[]token.Token{
				{
					Type:    token.K_PRECEDING,
					Literal: "PRECEDING",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"RANGE",
			[]token.Token{
				{
					Type:    token.K_RANGE,
					Literal: "RANGE",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"RECURSIVE",
			[]token.Token{
				{
					Type:    token.K_RECURSIVE,
					Literal: "RECURSIVE",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"REGEXP",
			[]token.Token{
				{
					Type:    token.K_REGEXP,
					Literal: "REGEXP",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"RIGHT",
			[]token.Token{
				{
					Type:    token.K_RIGHT,
					Literal: "RIGHT",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"ROW",
			[]token.Token{
				{
					Type:    token.K_ROW,
					Literal: "ROW",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"ROWS",
			[]token.Token{
				{
					Type:    token.K_ROWS,
					Literal: "ROWS",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"SELECT",
			[]token.Token{
				{
					Type:    token.K_SELECT,
					Literal: "SELECT",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"THEN",
			[]token.Token{
				{
					Type:    token.K_THEN,
					Literal: "THEN",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"TIES",
			[]token.Token{
				{
					Type:    token.K_TIES,
					Literal: "TIES",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"TRUE",
			[]token.Token{
				{
					Type:    token.K_TRUE,
					Literal: "TRUE",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"UNBOUNDED",
			[]token.Token{
				{
					Type:    token.K_UNBOUNDED,
					Literal: "UNBOUNDED",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"UNION",
			[]token.Token{
				{
					Type:    token.K_UNION,
					Literal: "UNION",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"USING",
			[]token.Token{
				{
					Type:    token.K_USING,
					Literal: "USING",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"VALUES",
			[]token.Token{
				{
					Type:    token.K_VALUES,
					Literal: "VALUES",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"WHEN",
			[]token.Token{
				{
					Type:    token.K_WHEN,
					Literal: "WHEN",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"WHERE",
			[]token.Token{
				{
					Type:    token.K_WHERE,
					Literal: "WHERE",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"WINDOW",
			[]token.Token{
				{
					Type:    token.K_WINDOW,
					Literal: "WINDOW",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			"WITH",
			[]token.Token{
				{
					Type:    token.K_WITH,
					Literal: "WITH",
				},
				{
					Type: token.EOS,
				},
			},
		},
	}

	for i, tc := range testCases {
		r, err := Tokenize(tc.input)
		if err != nil {
			t.Fatalf("[%d] returned err not nil %v input(%s)", i, err, tc.input)
		}
		if len(tc.expected) != len(r) {
			t.Errorf("[%d] Mistmatch length expected %d, but got %d\n", i, len(tc.expected), len(r))
			return
		}
		for n, tk := range tc.expected {
			if !tk.Equals(r[n]) {
				t.Errorf("[%d] Token Mistmatch expected %s(%s), but got %s(%s)\n", i, tk.Type.String(), tk.Literal, r[n].Type.String(), r[n].Literal)
			}
		}
	}
}

func TestErrorTokenize(t *testing.T) {
	testCases := []struct {
		input    string
		expected []token.Token
	}{
		{
			"\"Comm\nent\"",
			[]token.Token{
				{
					Type: token.INVALID,
				},
			},
		},
	}
	for i, tc := range testCases {
		r, err := Tokenize(tc.input)
		if err == nil {
			t.Fatalf("[%d] returned err is nil expected not nil", i)
		}
		if len(tc.expected) != len(r) {
			t.Errorf("[%d] Mistmatch length expected %d, but got %d\n", i, len(tc.expected), len(r))
			return
		}
		for n, tk := range tc.expected {
			if !tk.Equals(r[n]) {
				t.Errorf("[%d] Token Mistmatch expected %s(%s), but got %s(%s)\n", i, tk.Type.String(), tk.Literal, r[n].Type.String(), r[n].Literal)
			}
		}
	}
}
