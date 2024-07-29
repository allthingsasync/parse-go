package syntax

type SyntaxKind int

const (
    PlusToken SyntaxKind = iota
    MinusToken
    StarToken
    SlashToken
    PlusPlusToken
    MinusMinusToken
    PlusEqToken
    MinusEqToken
    StarEqToken
    SlashEqToken

    EqToken
    EqEqToken
    BangToken
    BangEqToken

    OpenParenToken
    CloseParenToken
    OpenBracketToken
    CloseBracketToken
    OpenBraceToken
    CloseBraceToken

    LessToken
    LessLessToken
    LessEqToken
    MoreToken
    MoreEqToken

    DotToken
    ColonToken
    ColonColonToken
    CommaToken
    QuestionToken
    CaretToken
    AtToken
    AndToken
    AndAndToken
    PipeToken
    PipePipeToken
    SemicolonToken
    TildeToken

    IdentifierToken
    StringLiteralToken
    NumericLiteralToken

    IfKeyword
    ElseKeyword
    WhileKeyword
    ExtendsKeyword
    ImplementsKeyword
    ReturnKeyword
    UsingKeyword
    PrivateKeyword
    PublicKeyword
    InternalKeyword
    ProtectedKeyword
    TrueKeyword
    FalseKeyword

    Usings
    UsingDirective
    ElseClause
    ExtendsClause
    ImplementsClause
    Modifiers

    IfStatement
    WhileStatement
    ReturnStatement
    ExpressionStatement

    IdentifierName
    QualifiedName


    EofToken
    BadToken
)
