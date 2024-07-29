package parser

import (
	"github.com/allthingsasync/parse_go/syntax"
	"github.com/allthingsasync/parse_go/text"
)

type Lexer struct {
	source text.SourceText
}

func LexerForString(source string) Lexer {
	return Lexer{
		source: text.NewSource(source),
	}
}

func LexerForSource(source text.SourceText) Lexer {
	return Lexer{
		source: source,
	}
}

func (this *Lexer) LexToken() *syntax.Token {
	if this.source.IsEnd() {
		return syntax.NewToken(syntax.EofToken, "")
	}
	switch this.source.Current() {
	case '+':
		switch this.source.LA() {
		case '+':
			this.source.Advance(2)
			return syntax.NewToken(syntax.PlusPlusToken, "++")
		case '=':
			this.source.Advance(2)
			return syntax.NewToken(syntax.PlusEqToken, "+=")
		default:
			this.source.Advance(1)
			return syntax.NewToken(syntax.PlusToken, "+")
		}
	case '-':
		switch this.source.LA() {
		case '-':
			this.source.Advance(2)
			return syntax.NewToken(syntax.MinusMinusToken, "--")
		case '=':
			this.source.Advance(2)
			return syntax.NewToken(syntax.MinusEqToken, "-=")
		default:
			this.source.Advance(1)
			return syntax.NewToken(syntax.MinusToken, "-")
		}
	case '*':
		switch this.source.LA() {
		case '=':
			this.source.Advance(2)
			return syntax.NewToken(syntax.StarEqToken, "*=")
		default:
			this.source.Advance(1)
			return syntax.NewToken(syntax.StarToken, "*")
		}
	case '/':
		switch this.source.LA() {
		case '=':
			this.source.Advance(2)
			return syntax.NewToken(syntax.SlashEqToken, "/=")
		default:
			this.source.Advance(1)
			return syntax.NewToken(syntax.SlashToken, "/")
		}
	case '!':
		switch this.source.LA() {
		case '=':
			this.source.Advance(2)
			return syntax.NewToken(syntax.BangEqToken, "!=")
		default:
			this.source.Advance(1)
			return syntax.NewToken(syntax.BangToken, "!")
		}
	case '=':
		switch this.source.LA() {
		case '=':
			this.source.Advance(2)
			return syntax.NewToken(syntax.EqEqToken, "==")
		default:
			this.source.Advance(1)
			return syntax.NewToken(syntax.EqToken, "=")
		}
	case '&':
		switch this.source.LA() {
		case '&':
			this.source.Advance(2)
			return syntax.NewToken(syntax.AndAndToken, "&&")
		default:
			this.source.Advance(1)
			return syntax.NewToken(syntax.AndToken, "&")
		}
	case '|':
		switch this.source.LA() {
		case '|':
			this.source.Advance(2)
			return syntax.NewToken(syntax.PipePipeToken, "||")
		default:
			this.source.Advance(1)
			return syntax.NewToken(syntax.PipeToken, "|")
		}
	case '<':
		switch this.source.LA() {
		case '<':
			this.source.Advance(2)
			return syntax.NewToken(syntax.LessLessToken, "<<")
		case '=':
			this.source.Advance(2)
			return syntax.NewToken(syntax.LessEqToken, "<=")
		default:
			this.source.Advance(1)
			return syntax.NewToken(syntax.LessToken, "<")
		}
	case '>':
		switch this.source.LA() {
		case '=':
			this.source.Advance(2)
			return syntax.NewToken(syntax.MoreEqToken, ">=")
		default:
			this.source.Advance(1)
			return syntax.NewToken(syntax.MoreToken, ">")
		}
	case '(':
		this.source.Advance(1)
		return syntax.NewToken(syntax.OpenParenToken, "(")
	case ')':
		this.source.Advance(1)
		return syntax.NewToken(syntax.CloseParenToken, ")")
	case '[':
		this.source.Advance(1)
		return syntax.NewToken(syntax.OpenBracketToken, "[")
	case ']':
		this.source.Advance(1)
		return syntax.NewToken(syntax.CloseBracketToken, "]")
	case '{':
		this.source.Advance(1)
		return syntax.NewToken(syntax.OpenBraceToken, "{")
	case '}':
		this.source.Advance(1)
		return syntax.NewToken(syntax.CloseBraceToken, "}")
	case ':':
		this.source.Advance(1)
		return syntax.NewToken(syntax.ColonToken, ":")
	case ';':
		this.source.Advance(1)
		return syntax.NewToken(syntax.SemicolonToken, ";")
	case '.':
		this.source.Advance(1)
		return syntax.NewToken(syntax.DotToken, ".")
	case '?':
		this.source.Advance(1)
		return syntax.NewToken(syntax.QuestionToken, "?")
	case '~':
		this.source.Advance(1)
		return syntax.NewToken(syntax.TildeToken, "~")
	case '^':
		this.source.Advance(1)
		return syntax.NewToken(syntax.CaretToken, "^")
	default:
		return syntax.NewToken(syntax.BadToken, "")
	}
}
