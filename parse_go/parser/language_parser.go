package parser

import "github.com/allthingsasync/parse_go/syntax"

type Parser struct {
	l LexBuffer
}

func NewParser(source string) Parser {
	return Parser{
		l: NewLexBufferForString(source),
	}
}

func ParseText(source string) *syntax.Node {
	p := NewParser(source)
	return p.ParseCompilationUnit()
}

func (this *Parser) ParseCompilationUnit() *syntax.Node {
	return nil
}

func (this *Parser) parseUsings() *syntax.Node {
	ls := make([]syntax.NT, 0, 4)
	for this.currentKind() == syntax.UsingKeyword {
		usingClause := this.parseUsing()
		ls = append(ls, usingClause.ToNT())
	}
	return syntax.NewNode(syntax.Usings, ls)
}

func (this *Parser) parseUsing() *syntax.Node {
	usingKw := this.eatToken()
	name := this.parseQualifiedName()
	semicolonToken := this.matchToken(syntax.SemicolonToken)

	return syntax.NewNode(syntax.UsingDirective, []syntax.NT{usingKw.ToNT(), name.ToNT(), semicolonToken.ToNT()})
}

func (this *Parser) parseExpression() *syntax.Node {
	return nil
}

func (this *Parser) parseStatement() *syntax.Node {
	return nil
}

func (this *Parser) parseMethodDeclaration() *syntax.Node {
	return nil
}

func (this *Parser) parseFieldDeclaration() *syntax.Node {
	return nil
}

func (this *Parser) parseModifiers() *syntax.Node {
	ls := make([]syntax.NT, 0, 2)
	for true {
		switch this.currentKind() {
		case syntax.PrivateKeyword:
		case syntax.PublicKeyword:
		case syntax.InternalKeyword:
		case syntax.ProtectedKeyword:
			ls = append(ls, this.eatToken().ToNT())
		default:
			break
		}
	}
	return syntax.NewNode(syntax.Modifiers, ls)
}

func (this *Parser) parseTypeDeclaration() *syntax.Node {
	return nil
}

func (this *Parser) parseOptionalExtends() *syntax.Node {
	if this.currentKind() != syntax.ExtendsKeyword {
		return nil
	}

	extendsKw := this.eatToken()
	id := this.parseQualifiedName()

	return syntax.NewNode(syntax.ExtendsClause, []syntax.NT{extendsKw.ToNT(), id.ToNT()})
}

func (this *Parser) parseOptionalImplements() *syntax.Node {
	if this.currentKind() != syntax.ImplementsKeyword {
		return nil
	}

	extendsKw := this.eatToken()
	id := this.parseQualifiedName()

	return syntax.NewNode(syntax.ImplementsClause, []syntax.NT{extendsKw.ToNT(), id.ToNT()})
}

func (this *Parser) parseDoStatement() *syntax.Node {
	return nil
}

func (this *Parser) parseWhileStatement() *syntax.Node {
	whileKw := this.matchToken(syntax.WhileKeyword)
	openParen := this.matchToken(syntax.OpenParenToken)
	expr := this.parseExpression()
	closeParen := this.matchToken(syntax.CloseParenToken)
	stmt := this.parseStatement()

	return syntax.NewNode(syntax.WhileStatement, []syntax.NT{whileKw.ToNT(), openParen.ToNT(), expr.ToNT(), closeParen.ToNT(), stmt.ToNT()})
}

func (this *Parser) parseIfStatement() *syntax.Node {
	ifKw := this.matchToken(syntax.IfKeyword)
	openParen := this.matchToken(syntax.OpenParenToken)
	expr := this.parseExpression()
	closeParen := this.matchToken(syntax.CloseParenToken)
	stmt := this.parseStatement()
	optionalElse := this.parseOptionalElse()

	return syntax.NewNode(syntax.IfStatement, []syntax.NT{ifKw.ToNT(), openParen.ToNT(), expr.ToNT(), closeParen.ToNT(), stmt.ToNT(), syntax.FromNode(optionalElse)})
}

func (this *Parser) parseOptionalElse() *syntax.Node {

	if this.currentKind() != syntax.ElseKeyword {
		return nil
	}

	elseKw := this.eatToken()
	stmt := this.parseStatement()
	return syntax.NewNode(syntax.ElseClause, []syntax.NT{elseKw.ToNT(), stmt.ToNT()})
}

func (this *Parser) parseForStatement() *syntax.Node {
	return nil
}

func (this *Parser) parseReturnStatement() *syntax.Node {
	returnKw := this.matchToken(syntax.ReturnKeyword)
	var expr *syntax.Node = nil

	if this.currentKind() != syntax.SemicolonToken {
		expr = this.parseExpression()
	}
	semicolonToken := this.matchToken(syntax.SemicolonToken)

	return syntax.NewNode(syntax.ReturnStatement, []syntax.NT{returnKw.ToNT(), syntax.FromNode(expr), semicolonToken.ToNT()})
}

func (this *Parser) parseExpressionStatement() *syntax.Node {
	return nil
}

func (this *Parser) parseLocalDeclaration() *syntax.Node {
	return nil
}

func (this *Parser) parseQualifiedName() *syntax.Node {
	sn := this.parseSimpleName()
	for this.currentKind() == syntax.DotToken {
		dotToken := this.eatToken()
		r := this.parseQualifiedName()
		sn = syntax.NewNode(syntax.QualifiedName, []syntax.NT{sn.ToNT(), dotToken.ToNT(), r.ToNT()})
	}
	return sn
}

func (this *Parser) parseSimpleName() *syntax.Node {
	id := this.matchToken(syntax.IdentifierToken)
	return syntax.NewNode(syntax.IdentifierName, []syntax.NT{id.ToNT()})
}

func (this *Parser) matchToken(kind syntax.SyntaxKind) *syntax.Token {
	if this.l.Current().Kind() == kind {
		return this.l.EatToken()
	}
	return syntax.MissingToken(kind)
}

func (this *Parser) eatToken() *syntax.Token {
	return this.l.EatToken()
}

func (this *Parser) currentKind() syntax.SyntaxKind {
	return this.l.Current().Kind()
}
