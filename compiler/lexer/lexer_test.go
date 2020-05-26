package lexer

import (
	"testing"

	"github.com/yakawa/makeDatabase/common/token"
)

func TestTokenizer(t *testing.T) {
	testCases := []struct {
		input    string
		expected []token.Token
	}{
		{
			"\"",
			[]token.Token{
				token.Token{
					Type:    token.DOUBLEQUOTE,
					Literal: "\"",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"%",
			[]token.Token{
				token.Token{
					Type:    token.PERCENT,
					Literal: "%",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"&",
			[]token.Token{
				token.Token{
					Type:    token.AMPERSAND,
					Literal: "&",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"'",
			[]token.Token{
				token.Token{
					Type:    token.QUOTE,
					Literal: "'",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"(",
			[]token.Token{
				token.Token{
					Type:    token.LEFTPAREN,
					Literal: "(",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			")",
			[]token.Token{
				token.Token{
					Type:    token.RIGHTPAREN,
					Literal: ")",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"*",
			[]token.Token{
				token.Token{
					Type:    token.ASTERISK,
					Literal: "*",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"+",
			[]token.Token{
				token.Token{
					Type:    token.PLUSSIGN,
					Literal: "+",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			",",
			[]token.Token{
				token.Token{
					Type:    token.COMMA,
					Literal: ",",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"-",
			[]token.Token{
				token.Token{
					Type:    token.MINUSSIGN,
					Literal: "-",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			".",
			[]token.Token{
				token.Token{
					Type:    token.PERIOD,
					Literal: ".",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"/",
			[]token.Token{
				token.Token{
					Type:    token.SOLIDAS,
					Literal: "/",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			":",
			[]token.Token{
				token.Token{
					Type:    token.COLON,
					Literal: ":",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			";",
			[]token.Token{
				token.Token{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"<",
			[]token.Token{
				token.Token{
					Type:    token.LESSTHAN,
					Literal: "<",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"=",
			[]token.Token{
				token.Token{
					Type:    token.EQUALS,
					Literal: "=",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			">",
			[]token.Token{
				token.Token{
					Type:    token.GREATERTHAN,
					Literal: ">",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"?",
			[]token.Token{
				token.Token{
					Type:    token.QUESTION,
					Literal: "?",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"[",
			[]token.Token{
				token.Token{
					Type:    token.LEFTBRACKET,
					Literal: "[",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"]",
			[]token.Token{
				token.Token{
					Type:    token.RIGHTBRACKET,
					Literal: "]",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"^",
			[]token.Token{
				token.Token{
					Type:    token.CIRCUMFLEX,
					Literal: "^",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"_",
			[]token.Token{
				token.Token{
					Type:    token.UNDERSCORE,
					Literal: "_",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"|",
			[]token.Token{
				token.Token{
					Type:    token.VERTICALBAR,
					Literal: "|",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"{",
			[]token.Token{
				token.Token{
					Type:    token.LEFTBRACE,
					Literal: "{",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"}",
			[]token.Token{
				token.Token{
					Type:    token.RIGHTBRACE,
					Literal: "}",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"<>",
			[]token.Token{
				token.Token{
					Type:    token.NOTEQUALS,
					Literal: "<>",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			">=",
			[]token.Token{
				token.Token{
					Type:    token.GREATERTHANEQUALS,
					Literal: ">=",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"<=",
			[]token.Token{
				token.Token{
					Type:    token.LESSTHANEQUALS,
					Literal: "<=",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"||",
			[]token.Token{
				token.Token{
					Type:    token.CONCAT,
					Literal: "||",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"'Comment'",
			[]token.Token{
				token.Token{
					Type:    token.STRING,
					Literal: "Comment",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"\"Comment\"",
			[]token.Token{
				token.Token{
					Type:    token.STRING,
					Literal: "Comment",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"\"Comm\nent\"",
			[]token.Token{
				token.Token{
					Type: token.INVALID,
				},
			},
		},
		{
			"'Comment'''",
			[]token.Token{
				token.Token{
					Type:    token.STRING,
					Literal: "Comment'",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"'Comment\"\"'",
			[]token.Token{
				token.Token{
					Type:    token.STRING,
					Literal: "Comment\"\"",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"'Comment';",
			[]token.Token{
				token.Token{
					Type:    token.STRING,
					Literal: "Comment",
				},
				token.Token{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"\"Comment\";",
			[]token.Token{
				token.Token{
					Type:    token.STRING,
					Literal: "Comment",
				},
				token.Token{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"1",
			[]token.Token{
				token.Token{
					Type:    token.NUMBER,
					Literal: "1",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"12",
			[]token.Token{
				token.Token{
					Type:    token.NUMBER,
					Literal: "12",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"12.",
			[]token.Token{
				token.Token{
					Type:    token.NUMBER,
					Literal: "12.",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"12.1",
			[]token.Token{
				token.Token{
					Type:    token.NUMBER,
					Literal: "12.1",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			".1",
			[]token.Token{
				token.Token{
					Type:    token.NUMBER,
					Literal: ".1",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"12;",
			[]token.Token{
				token.Token{
					Type:    token.NUMBER,
					Literal: "12",
				},
				token.Token{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"4.2.1",
			[]token.Token{
				token.Token{
					Type:    token.NUMBER,
					Literal: "4.2",
				},
				token.Token{
					Type:    token.NUMBER,
					Literal: ".1",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"abc",
			[]token.Token{
				token.Token{
					Type:    token.IDENT,
					Literal: "abc",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"abc def",
			[]token.Token{
				token.Token{
					Type:    token.IDENT,
					Literal: "abc",
				},
				token.Token{
					Type:    token.IDENT,
					Literal: "def",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"abc\tdef",
			[]token.Token{
				token.Token{
					Type:    token.IDENT,
					Literal: "abc",
				},
				token.Token{
					Type:    token.IDENT,
					Literal: "def",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"abc\ndef",
			[]token.Token{
				token.Token{
					Type:    token.IDENT,
					Literal: "abc",
				},
				token.Token{
					Type:    token.IDENT,
					Literal: "def",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"'abc def'",
			[]token.Token{
				token.Token{
					Type:    token.STRING,
					Literal: "abc def",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"4e+2;",
			[]token.Token{
				token.Token{
					Type:    token.NUMBER,
					Literal: "4",
				},
				token.Token{
					Type:    token.IDENT,
					Literal: "e",
				},
				token.Token{
					Type:    token.PLUSSIGN,
					Literal: "+",
				},
				token.Token{
					Type:    token.NUMBER,
					Literal: "2",
				},
				token.Token{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"WITH",
			[]token.Token{
				token.Token{
					Type:    token.K_WITH,
					Literal: "WITH",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"with",
			[]token.Token{
				token.Token{
					Type:    token.K_WITH,
					Literal: "with",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"With",
			[]token.Token{
				token.Token{
					Type:    token.K_WITH,
					Literal: "With",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"ALL",
			[]token.Token{
				token.Token{
					Type:    token.K_ALL,
					Literal: "ALL",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"AND",
			[]token.Token{
				token.Token{
					Type:    token.K_AND,
					Literal: "AND",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"AS",
			[]token.Token{
				token.Token{
					Type:    token.K_AS,
					Literal: "AS",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"ASC",
			[]token.Token{
				token.Token{
					Type:    token.K_ASC,
					Literal: "ASC",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"BETWEEN",
			[]token.Token{
				token.Token{
					Type:    token.K_BETWEEN,
					Literal: "BETWEEN",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"BY",
			[]token.Token{
				token.Token{
					Type:    token.K_BY,
					Literal: "BY",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"CASE",
			[]token.Token{
				token.Token{
					Type:    token.K_CASE,
					Literal: "CASE",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"CAST",
			[]token.Token{
				token.Token{
					Type:    token.K_CAST,
					Literal: "CAST",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"COLLATE",
			[]token.Token{
				token.Token{
					Type:    token.K_COLLATE,
					Literal: "COLLATE",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"CROSS",
			[]token.Token{
				token.Token{
					Type:    token.K_CROSS,
					Literal: "CROSS",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"CURRENT",
			[]token.Token{
				token.Token{
					Type:    token.K_CURRENT,
					Literal: "CURRENT",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"CURRENT_DATE",
			[]token.Token{
				token.Token{
					Type:    token.K_CURRENT_DATE,
					Literal: "CURRENT_DATE",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"CURRENT_TIME",
			[]token.Token{
				token.Token{
					Type:    token.K_CURRENT_TIME,
					Literal: "CURRENT_TIME",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"CURRENT_TIMESTAMP",
			[]token.Token{
				token.Token{
					Type:    token.K_CURRENT_TIMESTAMP,
					Literal: "CURRENT_TIMESTAMP",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"DESC",
			[]token.Token{
				token.Token{
					Type:    token.K_DESC,
					Literal: "DESC",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"DISTINCT",
			[]token.Token{
				token.Token{
					Type:    token.K_DISTINCT,
					Literal: "DISTINCT",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"ELSE",
			[]token.Token{
				token.Token{
					Type:    token.K_ELSE,
					Literal: "ELSE",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"END",
			[]token.Token{
				token.Token{
					Type:    token.K_END,
					Literal: "END",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"EXCLUDE",
			[]token.Token{
				token.Token{
					Type:    token.K_EXCLUDE,
					Literal: "EXCLUDE",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"ESCAPE",
			[]token.Token{
				token.Token{
					Type:    token.K_ESCAPE,
					Literal: "ESCAPE",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"EXCEPT",
			[]token.Token{
				token.Token{
					Type:    token.K_EXCEPT,
					Literal: "EXCEPT",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"EXISTS",
			[]token.Token{
				token.Token{
					Type:    token.K_EXISTS,
					Literal: "EXISTS",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"FALSE",
			[]token.Token{
				token.Token{
					Type:    token.K_FALSE,
					Literal: "FALSE",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"FIRST",
			[]token.Token{
				token.Token{
					Type:    token.K_FIRST,
					Literal: "FIRST",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"FOLLOWING",
			[]token.Token{
				token.Token{
					Type:    token.K_FOLLOWING,
					Literal: "FOLLOWING",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"FROM",
			[]token.Token{
				token.Token{
					Type:    token.K_FROM,
					Literal: "FROM",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"GLOB",
			[]token.Token{
				token.Token{
					Type:    token.K_GLOB,
					Literal: "GLOB",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"GROUP",
			[]token.Token{
				token.Token{
					Type:    token.K_GROUP,
					Literal: "GROUP",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"GROUPS",
			[]token.Token{
				token.Token{
					Type:    token.K_GROUPS,
					Literal: "GROUPS",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"HAVING",
			[]token.Token{
				token.Token{
					Type:    token.K_HAVING,
					Literal: "HAVING",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"INTERSECT",
			[]token.Token{
				token.Token{
					Type:    token.K_INTERSECT,
					Literal: "INTERSECT",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"IN",
			[]token.Token{
				token.Token{
					Type:    token.K_IN,
					Literal: "IN",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"INDEXED",
			[]token.Token{
				token.Token{
					Type:    token.K_INDEXED,
					Literal: "INDEXED",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"INNER",
			[]token.Token{
				token.Token{
					Type:    token.K_INNER,
					Literal: "INNER",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"IS",
			[]token.Token{
				token.Token{
					Type:    token.K_IS,
					Literal: "IS",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"ISNULL",
			[]token.Token{
				token.Token{
					Type:    token.K_ISNULL,
					Literal: "ISNULL",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"JOIN",
			[]token.Token{
				token.Token{
					Type:    token.K_JOIN,
					Literal: "JOIN",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"LAST",
			[]token.Token{
				token.Token{
					Type:    token.K_LAST,
					Literal: "LAST",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"LEFT",
			[]token.Token{
				token.Token{
					Type:    token.K_LEFT,
					Literal: "LEFT",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"LIKE",
			[]token.Token{
				token.Token{
					Type:    token.K_LIKE,
					Literal: "LIKE",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"LIMIT",
			[]token.Token{
				token.Token{
					Type:    token.K_LIMIT,
					Literal: "LIMIT",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"MATCH",
			[]token.Token{
				token.Token{
					Type:    token.K_MATCH,
					Literal: "MATCH",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"NATURAL",
			[]token.Token{
				token.Token{
					Type:    token.K_NATURAL,
					Literal: "NATURAL",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"NO",
			[]token.Token{
				token.Token{
					Type:    token.K_NO,
					Literal: "NO",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"NOT",
			[]token.Token{
				token.Token{
					Type:    token.K_NOT,
					Literal: "NOT",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"NOTNULL",
			[]token.Token{
				token.Token{
					Type:    token.K_NOTNULL,
					Literal: "NOTNULL",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"NULL",
			[]token.Token{
				token.Token{
					Type:    token.K_NULL,
					Literal: "NULL",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"NULLS",
			[]token.Token{
				token.Token{
					Type:    token.K_NULLS,
					Literal: "NULLS",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"OFFSET",
			[]token.Token{
				token.Token{
					Type:    token.K_OFFSET,
					Literal: "OFFSET",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"ON",
			[]token.Token{
				token.Token{
					Type:    token.K_ON,
					Literal: "ON",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"ORDER",
			[]token.Token{
				token.Token{
					Type:    token.K_ORDER,
					Literal: "ORDER",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"OTHERS",
			[]token.Token{
				token.Token{
					Type:    token.K_OTHERS,
					Literal: "OTHERS",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"OUTER",
			[]token.Token{
				token.Token{
					Type:    token.K_OUTER,
					Literal: "OUTER",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"OVER",
			[]token.Token{
				token.Token{
					Type:    token.K_OVER,
					Literal: "OVER",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"PARTITION",
			[]token.Token{
				token.Token{
					Type:    token.K_PARTITION,
					Literal: "PARTITION",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"PRECEDING",
			[]token.Token{
				token.Token{
					Type:    token.K_PRECEDING,
					Literal: "PRECEDING",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"RANGE",
			[]token.Token{
				token.Token{
					Type:    token.K_RANGE,
					Literal: "RANGE",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"RECURSIVE",
			[]token.Token{
				token.Token{
					Type:    token.K_RECURSIVE,
					Literal: "RECURSIVE",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"REGEXP",
			[]token.Token{
				token.Token{
					Type:    token.K_REGEXP,
					Literal: "REGEXP",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"ROW",
			[]token.Token{
				token.Token{
					Type:    token.K_ROW,
					Literal: "ROW",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"ROWS",
			[]token.Token{
				token.Token{
					Type:    token.K_ROWS,
					Literal: "ROWS",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"SELECT",
			[]token.Token{
				token.Token{
					Type:    token.K_SELECT,
					Literal: "SELECT",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"THEN",
			[]token.Token{
				token.Token{
					Type:    token.K_THEN,
					Literal: "THEN",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"TIES",
			[]token.Token{
				token.Token{
					Type:    token.K_TIES,
					Literal: "TIES",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"TRUE",
			[]token.Token{
				token.Token{
					Type:    token.K_TRUE,
					Literal: "TRUE",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"UNBOUNDED",
			[]token.Token{
				token.Token{
					Type:    token.K_UNBOUNDED,
					Literal: "UNBOUNDED",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"UNION",
			[]token.Token{
				token.Token{
					Type:    token.K_UNION,
					Literal: "UNION",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"USING",
			[]token.Token{
				token.Token{
					Type:    token.K_USING,
					Literal: "USING",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"VALUES",
			[]token.Token{
				token.Token{
					Type:    token.K_VALUES,
					Literal: "VALUES",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"WHEN",
			[]token.Token{
				token.Token{
					Type:    token.K_WHEN,
					Literal: "WHEN",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"WHERE",
			[]token.Token{
				token.Token{
					Type:    token.K_WHERE,
					Literal: "WHERE",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"WINDOW",
			[]token.Token{
				token.Token{
					Type:    token.K_WINDOW,
					Literal: "WINDOW",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
		{
			"WITH",
			[]token.Token{
				token.Token{
					Type:    token.K_WITH,
					Literal: "WITH",
				},
				token.Token{
					Type: token.EOS,
				},
			},
		},
	}

	for i, tc := range testCases {
		l := NewLexer(tc.input)
		r := l.Tokenize()
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
