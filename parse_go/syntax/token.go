package syntax

type Token struct {
	text    string
	kind    SyntaxKind
	missing bool
}

func MissingToken(kind SyntaxKind) *Token {
	return &Token{
		kind:    kind,
		text:    "",
		missing: true,
	}
}

func NewToken(kind SyntaxKind, text string) *Token {
	return &Token{
		kind:    kind,
		text:    text,
		missing: false,
	}
}

func (token *Token) Kind() SyntaxKind {
	return token.kind
}

func (token *Token) Length() int {
	return len(token.text)
}

func (token *Token) Text() string {
	return token.text
}

func (Token *Token) ToString() string {
	return Token.text
}

func (token *Token) ToNT() NT {
	return FromToken(token)
}
