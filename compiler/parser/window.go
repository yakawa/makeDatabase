package parser

import (
	"errors"

	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/token"
	"github.com/yakawa/makeDatabase/logger"
)

func (p *parser) parseWindowFunc() (w *ast.WindowExpression, err error) {
	logger.Tracef("Parse: Window")
	defer logger.Tracef("Parse: Window End")

	for {
		if p.currentToken.Type != token.IDENT {
			return w, errors.New("Parse Error Invalid Token")
		}

		w = &ast.WindowExpression{}
		wd := &ast.WindowDefinition{}
		wd.Name = p.currentToken.Literal

		p.readToken()

		if p.currentToken.Type != token.K_AS {
			return w, errors.New("Parse Error Invalid Token")
		}

		p.readToken()

		if p.currentToken.Type != token.LEFTPAREN {
			return w, errors.New("Parse Error Invalid Token")
		}

		p.readToken()

		if p.currentToken.Type == token.IDENT {
			wd.BaseWindowName = p.currentToken.Literal
			p.readToken()
		}

		if p.currentToken.Type == token.K_PARTITION {
			p.readToken()
			if p.currentToken.Type != token.K_BY {
				return w, errors.New("Parse Error Invalid Token")
			}
			p.readToken()

			for {
				e, er := p.parseExpr(LOWEST)
				if er != nil {
					return w, er
				}
				wd.PartitionExpr = append(wd.PartitionExpr, *e)
				p.readToken()
				if p.currentToken.Type != token.COMMA {
					break
				}
				p.readToken()
			}
		}

		if p.currentToken.Type == token.K_ORDER {
			p.readToken()
			if p.currentToken.Type != token.K_BY {
				return w, errors.New("Parse Error Invalid Token")
			}
			p.readToken()
			for {
				ob := &ast.OrderClause{}
				e, er := p.parseExpr(LOWEST)
				if er != nil {
					return w, er
				}
				ob.Expr = e
				p.readToken()
				if p.currentToken.Type == token.K_COLLATE {
					p.readToken()
					ob.CollateName = p.currentToken.Literal
					p.readToken()
				}
				if p.currentToken.Type == token.K_ASC || p.currentToken.Type == token.K_DESC {
					if p.currentToken.Type == token.K_ASC {
						ob.Asc = true
					}
					if p.currentToken.Type == token.K_DESC {
						ob.Desc = true
					}
					p.readToken()
				}
				if p.currentToken.Type == token.K_NULLS {
					p.readToken()
					if p.currentToken.Type == token.K_FIRST || p.currentToken.Type == token.K_LAST {
						if p.currentToken.Type == token.K_FIRST {
							ob.NullsFirst = true
						}
						if p.currentToken.Type == token.K_LAST {
							ob.NullsFirst = false
						}
						p.readToken()
					}
				}
				wd.OrderExpr = append(wd.OrderExpr, *ob)
				if p.currentToken.Type != token.COMMA {
					break
				}
				p.readToken()
			}
		}
		if p.currentToken.Type == token.K_RANGE || p.currentToken.Type == token.K_ROWS || p.currentToken.Type == token.K_GROUPS {
			fs, er := p.parseFrameSpec()
			if er != nil {
				return w, er
			}
			wd.Frame = fs
		}

		w.Defn = append(w.Defn, *wd)

		if p.peekToken().Type != token.COMMA {
			break
		}
		p.readToken()
		p.readToken()
	}
	return
}
