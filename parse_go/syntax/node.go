package syntax

type Node struct {
	kind     SyntaxKind
	children []NT
	length   int
}

func NewNode(kind SyntaxKind, children []NT) *Node {
	len := 0
	for _, v := range children {
		len += v.Length()
	}
	return &Node{
		length:   len,
		children: children,
		kind:     kind,
	}
}

func (node *Node) Kind() SyntaxKind {
	return node.kind
}

func (node *Node) Children() []NT {
	return node.children
}

func (node *Node) Length() int {
	return node.length
}

func (node *Node) ToString() string {
	str := ""
	for _, v := range node.children {
		str += v.ToString()
	}
	return str
}

func (node *Node) ToNT() NT {
	return FromNode(node)
}
