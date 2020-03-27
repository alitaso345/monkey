package parser

import (
	"github.com/alitaso345/monkey/ast"
	"github.com/alitaso345/monkey/lexer"
	"github.com/alitaso345/monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 2つトークンを読み込む。curTokenとpeerTokenの両方がセットされる
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
