package tree

/**
 * Simple Double Entry Accounting V3 for Go
 * Tree implementation
 * @author Ashley Kitson
 * @copyright Ashley Kitson, 2022, UK
 * @license BSD-3-Clause See LICENSE.md
 */

//NodeIFace an interface for a Tree Node
type NodeIFace interface {
	//SetValue sets the value of this node and returns this node
	SetValue(v interface{}) NodeIFace
	//GetValue returns the value of this node
	GetValue() interface{}
	//AddChild adds a child to the set of children for this node and returns this node
	AddChild(NodeIFace) NodeIFace
	//RemoveChild removes the child node matching the input parameter and returns this node
	RemoveChild(NodeIFace) NodeIFace
	//RemoveAllChildren removes all child nodes of this node (and their children) and returns this node
	RemoveAllChildren() NodeIFace
	//GetChildren returns all child nodes of this node
	GetChildren() []NodeIFace
	//SetChildren replaces all child nodes of this node with new ones and returns this node
	SetChildren(...NodeIFace) NodeIFace
	//SetParent sets the parent of this node and returns this node
	SetParent(NodeIFace) NodeIFace
	//GetParent returns the parent node of this node or nil if none
	GetParent() NodeIFace
	//GetAncestors retrieves all ancestors of node excluding current node
	GetAncestors() []NodeIFace
	//GetAncestorsAndSelf retrieves all ancestors of node as well as the node itself
	GetAncestorsAndSelf() []NodeIFace
	//GetSiblings retrieves all neighboring nodes (children of same parent), excluding the current node
	GetSiblings() []NodeIFace
	//GetSiblingsAndSelf returns all neighboring nodes (children of same parent), including the current node
	GetSiblingsAndSelf() []NodeIFace
	//IsRoot returns true if the node is the root, false otherwise
	IsRoot() bool
	//IsChild returns true if the node is a child, false otherwise
	IsChild() bool
	//IsLeaf  returns true if the node is a leaf node, false otherwise
	IsLeaf() bool
	//GetDepth returns the distance from the current node to the root
	GetDepth() int
	//GetHeight returns the height of the tree whose root is this node
	GetHeight() int
	//GetSize returns the number of nodes in the tree rooted at this node
	GetSize() int
	//Accept Accept method for the visitor pattern (see http://en.wikipedia.org/wiki/Visitor_pattern)
	Accept(v VisitorIFace) interface{}
}

//Node is a Tree Node
type Node struct {
	NodeIFace
	value    interface{}
	children []NodeIFace
	parent   NodeIFace
}

//NewNode returns a new Node
func NewNode(v interface{}, children *[]NodeIFace) NodeIFace {
	n := &Node{
		children: make([]NodeIFace, 0),
	}
	if children != nil {
		return n.SetValue(v).SetChildren(*children...)
	}

	return n.SetValue(v)
}

func (n *Node) SetValue(v interface{}) NodeIFace {
	n.value = v
	return n
}

func (n *Node) GetValue() interface{} {
	return n.value
}

func (n *Node) AddChild(c NodeIFace) NodeIFace {
	c = c.SetParent(n)
	n.children = append(n.children, c)
	return n
}

func (n *Node) RemoveChild(c NodeIFace) NodeIFace {
	for i, ch := range n.children {
		if c == ch {
			n.children[i] = n.children[len(n.children)-1] // Copy last element to index i.
			n.children[len(n.children)-1] = nil           // Erase last element (write zero value).
			n.children = n.children[:len(n.children)-1]   // Truncate slice.
		}
	}
	c.SetParent(nil)
	return n
}

func (n *Node) RemoveAllChildren() NodeIFace {
	for _, c := range n.children {
		c.SetParent(nil)
	}
	n.children = make([]NodeIFace, 0)
	return n
}

func (n *Node) GetChildren() []NodeIFace {
	return n.children
}

func (n *Node) SetChildren(c ...NodeIFace) NodeIFace {
	for _, child := range n.GetChildren() {
		child.SetParent(nil)
	}
	for _, cc := range c {
		n = n.AddChild(cc).(*Node)
	}
	return n
}

func (n *Node) SetParent(p NodeIFace) NodeIFace {
	n.parent = p
	return n
}

func (n *Node) GetParent() NodeIFace {
	return n.parent
}

func (n *Node) GetAncestors() []NodeIFace {
	parents := make([]NodeIFace, 0)
	node := n
	for {
		if p := node.GetParent(); p != nil {
			parents = append([]NodeIFace{p}, parents...)
			node = p.(*Node)
			continue
		}
		break
	}
	return parents
}

func (n *Node) GetAncestorsAndSelf() []NodeIFace {
	parents := n.GetAncestors()
	parents = append(parents, n)
	return parents
}

func (n *Node) GetSiblings() []NodeIFace {
	siblings := n.GetSiblingsAndSelf()
	for i, v := range siblings {
		if v == n {
			siblings[i] = siblings[len(siblings)-1] // Copy last element to index i.
			siblings[len(siblings)-1] = nil         // Erase last element (write zero value).
			siblings = siblings[:len(siblings)-1]   // Truncate slice.
		}
	}
	return siblings
}

func (n *Node) GetSiblingsAndSelf() []NodeIFace {
	return n.GetParent().GetChildren()
}

func (n *Node) IsRoot() bool {
	return n.parent == nil
}

func (n *Node) IsChild() bool {
	return n.parent != nil
}

func (n *Node) IsLeaf() bool {
	return len(n.children) == 0
}

func (n *Node) GetDepth() int {
	if n.IsRoot() {
		return 0
	}
	return n.GetParent().GetDepth() + 1
}

func (n *Node) GetHeight() int {
	if n.IsLeaf() {
		return 0
	}
	heights := make([]int, 0)
	for _, c := range n.GetChildren() {
		heights = append(heights, c.GetHeight())
	}
	return maxIntSlice(heights) + 1
}

func maxIntSlice(v []int) (m int) {
	if len(v) > 0 {
		m = v[0]
	}
	for i := 1; i < len(v); i++ {
		if v[i] > m {
			m = v[i]
		}
	}
	return
}

func (n *Node) GetSize() int {
	size := 1

	for _, child := range n.GetChildren() {
		size += child.GetSize()
	}
	return size
}

func (n *Node) Accept(v VisitorIFace) interface{} {
	return v.Visit(n)
}
