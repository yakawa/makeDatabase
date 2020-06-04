package parser

import (
	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/errors"
	"github.com/yakawa/makeDatabase/common/token"
	"github.com/yakawa/makeDatabase/logger"
)

func (p *parser) parseWindowClause() (w *ast.WindowExpression, err error) {
	logger.Tracef("Parse: Window")
	defer logger.Tracef("Parse: Window End")

	for {
		if p.currentToken.Type != token.IDENT {
			return w, errors.NewErrParseInvalidToken(p.currentToken)
		}

		w = &ast.WindowExpression{}
		name := p.currentToken.Literal

		p.readToken()

		if p.currentToken.Type != token.K_AS {
			return w, errors.NewErrParseInvalidToken(p.currentToken)
		}

		p.readToken()

		if p.currentToken.Type != token.LEFTPAREN {
			return w, errors.NewErrParseInvalidToken(p.currentToken)
		}

		p.readToken()

		d, er := p.parseWindowDefinition()
		if er != nil {
			return w, er
		}

		d.Name = name

		w.Defn = append(w.Defn, *d)

		if p.currentToken.Type != token.COMMA {
			break
		}
		p.readToken()
	}

	if p.currentToken.Type != token.RIGHTPAREN {
		return w, errors.NewErrParseInvalidToken(p.currentToken)
	}
	p.readToken()

	return
}

func (p *parser) parseWindowDefinition() (d *ast.WindowDefinition, err error) {
	d = &ast.WindowDefinition{}
	if p.currentToken.Type == token.IDENT {
		d.BaseWindowName = p.currentToken.Literal
		p.readToken()
	}

	if p.currentToken.Type == token.K_PARTITION {
		p.readToken()
		if p.currentToken.Type != token.K_BY {
			return d, errors.NewErrParseInvalidToken(p.currentToken)
		}
		p.readToken()

		for {
			e, er := p.parseExpr(LOWEST)
			if er != nil {
				return d, er
			}
			d.PartitionExpr = append(d.PartitionExpr, *e)
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
			return d, errors.NewErrParseInvalidToken(p.currentToken)
		}
		p.readToken()
		for {
			ob := &ast.OrderClause{}
			e, er := p.parseExpr(LOWEST)
			if er != nil {
				return d, er
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
			d.OrderExpr = append(d.OrderExpr, *ob)
			if p.currentToken.Type != token.COMMA {
				break
			}
			p.readToken()
		}
	}
	if p.currentToken.Type == token.K_RANGE || p.currentToken.Type == token.K_ROWS || p.currentToken.Type == token.K_GROUPS {
		fs, er := p.parseFrameSpec()
		if er != nil {
			return d, er
		}
		d.Frame = fs
	}
	p.readToken()
	if p.currentToken.Type != token.RIGHTPAREN {
		return d, errors.NewErrParseInvalidToken(p.currentToken)
	}
	p.readToken()
	return
}

func (p *parser) parseFrameSpec() (fs *ast.FrameSpecification, err error) {
	fs = &ast.FrameSpecification{}
	if p.currentToken.Type == token.K_RANGE || p.currentToken.Type == token.K_ROWS || p.currentToken.Type == token.K_GROUPS {
		if p.currentToken.Type == token.K_RANGE {
			fs.Range = true
		} else if p.currentToken.Type == token.K_ROWS {
			fs.Rows = true
		} else {
			fs.Groups = true
		}
		p.readToken()
	}
	if p.currentToken.Type == token.K_BETWEEN {
		p.readToken()
		if p.currentToken.Type == token.K_UNBOUNDED {
			p.readToken()
			if p.currentToken.Type != token.K_PRECEDING {
				return fs, errors.NewErrParseInvalidToken(p.currentToken)
			}
			fs.UnboundedPreceding1 = true
		} else if p.currentToken.Type == token.K_CURRENT {
			p.readToken()
			if p.currentToken.Type != token.K_ROW {
				return fs, errors.NewErrParseInvalidToken(p.currentToken)
			}
			fs.CurrentRow1 = true
		} else {
			ex, er := p.parseExpr(LOWEST)
			if er != nil {
				return fs, er
			}
			p.readToken()
			if p.currentToken.Type == token.K_PRECEDING {
				fs.ExprPreceding1 = ex
			} else if p.currentToken.Type == token.K_FOLLOWING {
				fs.ExprFollowing1 = ex
			} else {
				return fs, errors.NewErrParseInvalidToken(p.currentToken)
			}
		}
		p.readToken()
		if p.currentToken.Type != token.K_AND {
			return fs, errors.NewErrParseInvalidToken(p.currentToken)
		}
		p.readToken()
		if p.currentToken.Type == token.K_UNBOUNDED {
			p.readToken()
			if p.currentToken.Type != token.K_FOLLOWING {
				return fs, errors.NewErrParseInvalidToken(p.currentToken)
			}
			fs.UnboundedFollowing2 = true
		} else if p.currentToken.Type == token.K_CURRENT {
			p.readToken()
			if p.currentToken.Type != token.K_ROW {
				return fs, errors.NewErrParseInvalidToken(p.currentToken)
			}
			fs.CurrentRow2 = true
		} else {
			ex, er := p.parseExpr(LOWEST)
			if er != nil {
				return fs, er
			}
			p.readToken()
			if p.currentToken.Type == token.K_PRECEDING {
				fs.ExprPreceding2 = ex
			} else if p.currentToken.Type == token.K_FOLLOWING {
				fs.ExprFollowing2 = ex
			} else {
				return fs, errors.NewErrParseInvalidToken(p.currentToken)
			}
		}

	} else if p.currentToken.Type == token.K_UNBOUNDED {
		p.readToken()
		if p.currentToken.Type != token.K_PRECEDING {
			return fs, errors.NewErrParseInvalidToken(p.currentToken)
		}
		fs.UnboundedPreceding = true
	} else if p.currentToken.Type == token.K_CURRENT {
		p.readToken()
		if p.currentToken.Type != token.K_ROW {
			return fs, errors.NewErrParseInvalidToken(p.currentToken)
		}
		fs.CurrentRow = true
	} else {
		ex, er := p.parseExpr(LOWEST)
		if er != nil {
			return fs, er
		}
		fs.ExprPreceding = ex
		p.readToken()
		if p.currentToken.Type != token.K_PRECEDING {
			return fs, errors.NewErrParseInvalidToken(p.currentToken)
		}
	}

	if p.peekToken().Type == token.K_EXCLUDE {
		p.readToken()
		p.readToken()
		if p.currentToken.Type == token.K_NO {
			p.readToken()
			if p.currentToken.Type != token.K_OTHERS {
				return fs, errors.NewErrParseInvalidToken(p.currentToken)
			}
			fs.ExcludeNoOthers = true
		} else if p.currentToken.Type == token.K_CURRENT {
			p.readToken()
			if p.currentToken.Type != token.K_ROW {
				return fs, errors.NewErrParseInvalidToken(p.currentToken)
			}
			fs.ExcludeCurrentRow = true
		} else if p.currentToken.Type == token.K_GROUP {
			fs.ExcludeGroup = true
		} else if p.currentToken.Type == token.K_TIES {
			fs.ExcludeTies = true
		}
	}
	return
}
