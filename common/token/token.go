package token

import "strings"

//go:generate stringer -type=Type
type Type int

const (
	INVALID Type = -1
	EOS     Type = iota
	IDENT
	KEYWORD
	SYMBOL
	NUMBER
	STRING
)

const (
	DOUBLEQUOTE Type = iota + 10
	PERCENT
	AMPERSAND
	QUOTE
	LEFTPAREN
	RIGHTPAREN
	ASTERISK
	PLUSSIGN
	COMMA
	MINUSSIGN
	PERIOD
	SOLIDAS
	COLON
	SEMICOLON
	LESSTHAN
	EQUALS
	GREATERTHAN
	QUESTION
	LEFTBRACKET
	RIGHTBRACKET
	CIRCUMFLEX
	UNDERSCORE
	VERTICALBAR
	LEFTBRACE
	RIGHTBRACE
	NOTEQUALS
	GREATERTHANEQUALS
	LESSTHANEQUALS
	CONCAT
)

const (
	K_ALL Type = iota + 100
	K_AND
	K_AS
	K_ASC
	K_BETWEEN
	K_BY
	K_CASE
	K_CAST
	K_COLLATE
	K_CROSS
	K_CURRENT
	K_CURRENT_DATE
	K_CURRENT_TIME
	K_CURRENT_TIMESTAMP
	K_DESC
	K_DISTINCT
	K_ELSE
	K_END
	K_EXCLUDE
	K_ESCAPE
	K_EXCEPT
	K_EXISTS
	K_FALSE
	K_FIRST
	K_FOLLOWING
	K_FROM
	K_GLOB
	K_GROUP
	K_GROUPS
	K_HAVING
	K_INTERSECT
	K_IN
	K_INDEXED
	K_INNER
	K_IS
	K_ISNULL
	K_JOIN
	K_LAST
	K_LEFT
	K_LIKE
	K_LIMIT
	K_MATCH
	K_NATURAL
	K_NO
	K_NOT
	K_NOTNULL
	K_NULL
	K_NULLS
	K_OFFSET
	K_ON
	K_ORDER
	K_OTHERS
	K_OUTER
	K_OVER
	K_PARTITION
	K_PRECEDING
	K_RANGE
	K_RECURSIVE
	K_REGEXP
	K_RIGHT
	K_ROW
	K_ROWS
	K_SELECT
	K_THEN
	K_TIES
	K_TRUE
	K_UNBOUNDED
	K_UNION
	K_USING
	K_VALUES
	K_WHEN
	K_WHERE
	K_WINDOW
	K_WITH
)

type Token struct {
	Type    Type
	Literal string
	Pos     int
	Line    int
}

func (t *Token) Equals(tk Token) bool {
	if t.Type == tk.Type && t.Literal == tk.Literal {
		return true
	}
	return false
}

func LookupKeyword(k string) Type {
	switch strings.ToUpper(k) {
	case "ALL":
		return K_ALL
	case "AND":
		return K_AND
	case "AS":
		return K_AS
	case "ASC":
		return K_ASC
	case "BETWEEN":
		return K_BETWEEN
	case "BY":
		return K_BY
	case "CASE":
		return K_CASE
	case "CAST":
		return K_CAST
	case "COLLATE":
		return K_COLLATE
	case "CROSS":
		return K_CROSS
	case "CURRENT":
		return K_CURRENT
	case "CURRENT_DATE":
		return K_CURRENT_DATE
	case "CURRENT_TIME":
		return K_CURRENT_TIME
	case "CURRENT_TIMESTAMP":
		return K_CURRENT_TIMESTAMP
	case "DESC":
		return K_DESC
	case "DISTINCT":
		return K_DISTINCT
	case "ELSE":
		return K_ELSE
	case "END":
		return K_END
	case "EXCLUDE":
		return K_EXCLUDE
	case "ESCAPE":
		return K_ESCAPE
	case "EXCEPT":
		return K_EXCEPT
	case "EXISTS":
		return K_EXISTS
	case "FALSE":
		return K_FALSE
	case "FIRST":
		return K_FIRST
	case "FOLLOWING":
		return K_FOLLOWING
	case "FROM":
		return K_FROM
	case "GLOB":
		return K_GLOB
	case "GROUP":
		return K_GROUP
	case "GROUPS":
		return K_GROUPS
	case "HAVING":
		return K_HAVING
	case "INTERSECT":
		return K_INTERSECT
	case "IN":
		return K_IN
	case "INDEXED":
		return K_INDEXED
	case "INNER":
		return K_INNER
	case "IS":
		return K_IS
	case "ISNULL":
		return K_ISNULL
	case "JOIN":
		return K_JOIN
	case "LAST":
		return K_LAST
	case "LEFT":
		return K_LEFT
	case "LIKE":
		return K_LIKE
	case "LIMIT":
		return K_LIMIT
	case "MATCH":
		return K_MATCH
	case "NATURAL":
		return K_NATURAL
	case "NO":
		return K_NO
	case "NOT":
		return K_NOT
	case "NOTNULL":
		return K_NOTNULL
	case "NULL":
		return K_NULL
	case "NULLS":
		return K_NULLS
	case "OFFSET":
		return K_OFFSET
	case "ON":
		return K_ON
	case "ORDER":
		return K_ORDER
	case "OTHERS":
		return K_OTHERS
	case "OUTER":
		return K_OUTER
	case "OVER":
		return K_OVER
	case "PARTITION":
		return K_PARTITION
	case "PRECEDING":
		return K_PRECEDING
	case "RANGE":
		return K_RANGE
	case "RECURSIVE":
		return K_RECURSIVE
	case "REGEXP":
		return K_REGEXP
	case "RIGHT":
		return K_RIGHT
	case "ROW":
		return K_ROW
	case "ROWS":
		return K_ROWS
	case "SELECT":
		return K_SELECT
	case "THEN":
		return K_THEN
	case "TIES":
		return K_TIES
	case "TRUE":
		return K_TRUE
	case "UNBOUNDED":
		return K_UNBOUNDED
	case "UNION":
		return K_UNION
	case "USING":
		return K_USING
	case "VALUES":
		return K_VALUES
	case "WHEN":
		return K_WHEN
	case "WHERE":
		return K_WHERE
	case "WINDOW":
		return K_WINDOW
	case "WITH":
		return K_WITH
	}
	return IDENT
}
