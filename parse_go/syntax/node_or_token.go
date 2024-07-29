package syntax

type NT struct {
	token  *Token
	node   *Node
	isNode bool
}

func FromNode(node *Node) NT {
	return NT{
		node:   node,
		token:  nil,
		isNode: true,
	}
}

func FromToken(token *Token) NT {
	return NT{
		node:   nil,
		token:  token,
		isNode: false,
	}
}

func (nt *NT) IsNode() bool {
	return nt.isNode
}

func (nt *NT) IsToken() bool {
	return !nt.isNode
}

func (nt *NT) Token() *Token {
	return nt.token
}

func (nt *NT) Node() *Node {
	return nt.node
}

func (nt *NT) Length() int {
	if nt.IsNode() {
		return nt.node.Length()
	} else {
		return nt.token.Length()
	}
}

func (nt *NT) Kind() SyntaxKind {
	if nt.IsNode() {
		return nt.node.Kind()
	} else {
		return nt.token.Kind()
	}
}

func (nt *NT) Children() *[]NT {
	if nt.IsNode() {
		return &nt.node.children
	} else {
		return nil
	}
}

func (nt *NT) ToString() string {
	if nt.IsNode() {
		return nt.node.ToString()
	} else {
		return nt.token.ToString()
	}
}
