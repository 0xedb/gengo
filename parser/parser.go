package parser

import (
	"go/token"

	"github.com/0xedb/gengo/lexer"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}
