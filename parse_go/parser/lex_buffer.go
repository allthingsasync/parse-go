package parser

import "github.com/allthingsasync/parse_go/syntax"

type LexBuffer struct {
	l Lexer
    tokens []*syntax.Token
    current int
    length int
}

func NewLexBufferForString(source string) LexBuffer {
    lexer := LexerForString(source)
    tokens := make([]*syntax.Token, 5, 20)
    tok := lexer.LexToken()
    for tok.Kind() != syntax.EofToken {
        tokens = append(tokens, tok)
        tok = lexer.LexToken()
    }
    tokens = append(tokens, tok)
	return LexBuffer{
		l: lexer,
        tokens: tokens,
        current: 0,
        length: 0,
	}
}

func (this *LexBuffer) Peek(n int) *syntax.Token {
	if this.current + n >= this.length {
        return this.tokens[this.length-1]
    }
    return this.tokens[this.current + n]
}

func (this *LexBuffer) Current() *syntax.Token {
	return this.Peek(0)
}

func (this *LexBuffer) LA() *syntax.Token {
	return this.Peek(1)
}

func (this *LexBuffer) EatToken() *syntax.Token {
    x := this.Current()
    this.current += 1
    return x
}
