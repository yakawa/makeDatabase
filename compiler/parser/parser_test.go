package parser

import (
	"testing"

	"github.com/yakawa/makeDatabase/compiler/lexer"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		input string
	}{
		{"SELECT a FROM t;"},
	}
	for _, tc := range testCases {
		l, _ := lexer.Tokenize(tc.input)
		Parse(l)
	}
}
