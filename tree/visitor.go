package tree

/**
 * Simple Double Entry Accounting V3 for Go
 * Tree implementation
 * @author Ashley Kitson
 * @copyright Ashley Kitson, 2022, UK
 * @license BSD-3-Clause See LICENSE.md
 */

//VisitorIFace Visitor interface for Nodes
type VisitorIFace interface {
	//Visit visits each node starting at given node
	Visit(NodeIFace) interface{}
}

//PreOrderVisitor walks a tree in pre-order
type PreOrderVisitor struct {
	VisitorIFace
}

func NewPreOrderVisitor() VisitorIFace {
	return new(PreOrderVisitor)
}

//Visit returns []NodeIFace in pre order
func (v *PreOrderVisitor) Visit(n NodeIFace) interface{} {
	nodes := make([]NodeIFace, 1)
	nodes[0] = n
	for _, child := range n.GetChildren() {
		vNodes := child.Accept(v).([]NodeIFace)
		nodes = append(nodes, vNodes...)
	}
	return nodes
}

//PostOrderVisitor walks a tree in post-order
type PostOrderVisitor struct {
	VisitorIFace
}

func NewPostOrderVisitor() VisitorIFace {
	return new(PostOrderVisitor)
}

//Visit returns []NodeIFace in post order
func (v *PostOrderVisitor) Visit(n NodeIFace) interface{} {
	nodes := make([]NodeIFace, 0)
	for _, child := range n.GetChildren() {
		vNodes := child.Accept(v).([]NodeIFace)
		nodes = append(nodes, vNodes...)
	}
	nodes = append(nodes, n)
	return nodes
}

//LeafVisitor returns the leaves of a tree
type LeafVisitor struct {
	VisitorIFace
}

func NewLeafVisitor() VisitorIFace {
	return new(LeafVisitor)
}

//Visit returns []NodeIFace of leaves on the tree
func (v *LeafVisitor) Visit(n NodeIFace) interface{} {
	if n.IsLeaf() {
		return []NodeIFace{n}
	}
	nodes := make([]NodeIFace, 0)
	for _, child := range n.GetChildren() {
		vNodes := child.Accept(v).([]NodeIFace)
		nodes = append(nodes, vNodes...)
	}
	return nodes
}
