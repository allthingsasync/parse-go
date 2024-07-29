package parser

import (
    "testing"

	"github.com/allthingsasync/parse_go/syntax"
)

func TestLexerString(t *testing.T) {
    lexer := LexerForString("--+.::")
    tok := lexer.LexToken()
    t.Log(tok.Kind())

    for tok.Kind() != syntax.EofToken {
        t.Log(tok.Kind())
        tok = lexer.LexToken()
    }
}
