package parser

import (
	"errors"

	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/token"
	"github.com/yakawa/makeDatabase/logger"
)

type parser struct {
	tokens []token.Token

	currentToken token.Token
	pos          int

	unaryParseFunction  map[token.Type]unaryOpeFunction
	binaryParseFunction map[token.Type]binaryOpeFunction
}

type (
	unaryOpeFunction  func() (*ast.Expression, error)
	binaryOpeFunction func(*ast.Expression) (*ast.Expression, error)
)

const (
	_ int = iota
	LOWEST
	LOGICAL_AND
	LOGICAL_NOT
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
)

var precedences = map[token.Type]int{
	token.ASTERISK:  PRODUCT,
	token.SOLIDAS:   PRODUCT,
	token.PLUSSIGN:  SUM,
	token.MINUSSIGN: SUM,
	token.CONCAT:    SUM,
	token.EQUALS:    EQUALS,
	token.NOTEQUALS: EQUALS,
	token.K_COLLATE: EQUALS,
	token.K_NOT:     EQUALS,
	token.K_LIKE:    EQUALS,
	token.K_REGEXP:  EQUALS,
	token.K_GLOB:    EQUALS,
	token.K_MATCH:   EQUALS,
	token.K_ISNULL:  EQUALS,
	token.K_NOTNULL: EQUALS,
	token.K_IS:      EQUALS,
	token.K_BETWEEN: EQUALS,
	token.K_IN:      EQUALS,
}

// Parse
func Parse(tokens []token.Token) (*ast.SQL, error) {
	p := new(tokens)
	a, err := p.parse()
	return a, err
}

func new(tokens []token.Token) *parser {
	p := &parser{
		tokens: tokens,
	}

	p.unaryParseFunction = make(map[token.Type]unaryOpeFunction)
	p.binaryParseFunction = make(map[token.Type]binaryOpeFunction)

	p.unaryParseFunction[token.NUMBER] = p.parseNumber
	p.unaryParseFunction[token.IDENT] = p.parseIdent
	p.unaryParseFunction[token.STRING] = p.parseLiteral
	p.unaryParseFunction[token.MINUSSIGN] = p.parsePrefixExpr
	p.unaryParseFunction[token.LEFTPAREN] = p.parseGroupedExpr
	p.unaryParseFunction[token.K_CURRENT_DATE] = p.parseLiteral
	p.unaryParseFunction[token.K_CURRENT_TIME] = p.parseLiteral
	p.unaryParseFunction[token.K_CURRENT_TIMESTAMP] = p.parseLiteral
	p.unaryParseFunction[token.K_NULL] = p.parseLiteral
	p.unaryParseFunction[token.K_TRUE] = p.parseLiteral
	p.unaryParseFunction[token.K_FALSE] = p.parseLiteral
	p.unaryParseFunction[token.K_CAST] = p.parseCastExpr
	p.unaryParseFunction[token.K_NOT] = p.parseExistsExpr
	p.unaryParseFunction[token.K_EXISTS] = p.parseExistsExpr
	p.unaryParseFunction[token.K_CASE] = p.parseCaseExpr

	p.binaryParseFunction[token.PLUSSIGN] = p.parseBinaryExpr
	p.binaryParseFunction[token.MINUSSIGN] = p.parseBinaryExpr
	p.binaryParseFunction[token.ASTERISK] = p.parseBinaryExpr
	p.binaryParseFunction[token.SOLIDAS] = p.parseBinaryExpr
	p.binaryParseFunction[token.CONCAT] = p.parseBinaryExpr
	p.binaryParseFunction[token.EQUALS] = p.parseBinaryExpr
	p.binaryParseFunction[token.NOTEQUALS] = p.parseBinaryExpr
	p.binaryParseFunction[token.K_COLLATE] = p.parseCollateExpr
	p.binaryParseFunction[token.K_LIKE] = p.parseStringFunc
	p.binaryParseFunction[token.K_GLOB] = p.parseStringFunc
	p.binaryParseFunction[token.K_MATCH] = p.parseStringFunc
	p.binaryParseFunction[token.K_REGEXP] = p.parseStringFunc
	p.binaryParseFunction[token.K_ISNULL] = p.parseNullExpr
	p.binaryParseFunction[token.K_NOTNULL] = p.parseNullExpr
	p.binaryParseFunction[token.K_IS] = p.parseIsExpr
	p.binaryParseFunction[token.K_BETWEEN] = p.parseBetweenExpr
	p.binaryParseFunction[token.K_IN] = p.parseInExpr

	p.binaryParseFunction[token.K_NOT] = p.parseNotExpr

	return p
}

func (p *parser) parse() (a *ast.SQL, err error) {
	logger.Tracef("Parser Start!")
	defer logger.Tracef("Parse End!")

	p.readToken()
	a = &ast.SQL{}
	for p.currentToken.Type != token.EOS {
		switch p.currentToken.Type {
		case token.SEMICOLON, token.EOS:
			return
		case token.K_WITH, token.K_SELECT:
			ss, err := p.parseSelectStatement()
			if err != nil {
				return a, err
			}
			a.SelectStatement = ss
		default:
			err = errors.New("Parse Error: Unknown Token")
			return
		}
		p.readToken()
	}
	return
}

func (p *parser) parseSelectStatement() (ss *ast.SelectStatement, err error) {
	logger.Tracef("Parse: SELECT Statement")
	defer logger.Tracef("Parse: SELECT Statement End")

	ss = &ast.SelectStatement{}
	switch p.currentToken.Type {
	case token.K_WITH:
		wc, e := p.parseWithClause()
		if e != nil {
			return ss, e
		}
		ss.WithClause = wc
	case token.K_SELECT, token.K_VALUES:
		sc, e := p.parseSelectClause()
		if e != nil {
			return ss, e
		}
		ss.SelectClause = sc
	case token.K_ORDER:
	case token.K_LIMIT:
	default:
		err = errors.New("Parse Error Invalid Token")
		return
	}
	return
}

func (p *parser) readToken() {
	if p.pos >= len(p.tokens) {
		p.currentToken = token.Token{Type: token.EOS}
		return
	}
	p.currentToken = p.tokens[p.pos]
	p.pos++
	return
}

func (p *parser) rewindToken() {
	if p.pos-2 < 0 {
		p.currentToken = p.tokens[0]
		p.pos = 1
		return
	}
	p.currentToken = p.tokens[p.pos-2]
	p.pos--
	return
}

func (p *parser) peekToken() (t token.Token) {
	if p.pos >= len(p.tokens) {
		return token.Token{Type: token.EOS}
	}
	t = p.tokens[p.pos]
	return
}

func (p *parser) peek2Token() (t token.Token) {
	if p.pos+1 >= len(p.tokens) {
		return token.Token{Type: token.EOS}
	}
	t = p.tokens[p.pos+1]
	return
}

func (p *parser) getCurrentPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken().Type]; ok {
		return p
	}
	return LOWEST
}
