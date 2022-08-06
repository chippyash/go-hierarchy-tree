package tree_test

import (
	"github.com/chippyash/go-hierarchy-tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode_NewNode(t *testing.T) {
	node := tree.NewNode(nil, nil)
	//type assertion
	assert.IsType(t, tree.Node{}, *node.(*tree.Node))
	//interface assertion
	_, ok := node.(tree.NodeIFace)
	assert.True(t, ok)
}

func TestNode_NewNodeWithValue(t *testing.T) {
	node := tree.NewNode("foo", nil)
	assert.Equal(t, "foo", node.GetValue())
}

func TestNode_NewNodeWithChildren(t *testing.T) {
	children := []tree.NodeIFace{tree.NewNode(1, nil), tree.NewNode(2, nil)}
	node := tree.NewNode("foo", &children)
	assert.Equal(t, "foo", node.GetValue())
	assert.Equal(t, 2, len(node.GetChildren()))
}

func TestNode_SetAndGetValue(t *testing.T) {
	val := tree.NewNode(nil, nil).
		SetValue("foo").
		GetValue()
	assert.Equal(t, "foo", val)
}

func TestNode_SetAndGetChildren(t *testing.T) {
	children := []tree.NodeIFace{tree.NewNode(1, nil), tree.NewNode(2, nil)}
	root := tree.NewNode(0, nil)
	val := root.
		SetChildren(children...).
		GetChildren()
	assert.Equal(t, 2, len(val))
	//check that parent nodes are set on children
	assert.Equal(t, 0, val[0].GetParent().GetValue())
	assert.Equal(t, 0, val[1].GetParent().GetValue())
}

func TestNode_AddAndGetChildren(t *testing.T) {
	children := []tree.NodeIFace{tree.NewNode(1, nil), tree.NewNode(2, nil)}
	root := tree.NewNode(0, nil)
	val := root.
		AddChild(children[0]).
		AddChild(children[1]).
		GetChildren()
	assert.Equal(t, 2, len(val))
	//check that parent nodes are set on children
	assert.Equal(t, 0, val[0].GetParent().GetValue())
	assert.Equal(t, 0, val[1].GetParent().GetValue())
}

func TestNode_SetAndGetParent(t *testing.T) {
	root := tree.NewNode(0, nil)
	child := tree.NewNode(1, nil).SetParent(root)
	assert.Equal(t, 0, child.GetParent().GetValue())
}

func TestNode_RemoveChild(t *testing.T) {
	children := []tree.NodeIFace{tree.NewNode(1, nil), tree.NewNode(2, nil)}
	node := tree.NewNode("foo", &children).RemoveChild(children[0])
	assert.Equal(t, 1, len(node.GetChildren()))
	assert.Equal(t, 2, node.GetChildren()[0].GetValue())
	//ensure first original child has no parent
	assert.Nil(t, children[0].GetParent())
}

func TestNode_RemoveAllChildren(t *testing.T) {
	children := []tree.NodeIFace{tree.NewNode(1, nil), tree.NewNode(2, nil)}
	node := tree.NewNode("foo", &children).RemoveAllChildren()
	assert.Equal(t, 0, len(node.GetChildren()))
	//ensure first original children have no parent
	assert.Nil(t, children[0].GetParent())
	assert.Nil(t, children[1].GetParent())
}

func TestNode_GetAncestors(t *testing.T) {
	child3 := tree.NewNode(3, nil)
	child2 := tree.NewNode(2, &[]tree.NodeIFace{child3})
	child1 := tree.NewNode(1, &[]tree.NodeIFace{child2})
	root := tree.NewNode(1, &[]tree.NodeIFace{child1})

	expected := []tree.NodeIFace{root, child1, child2}
	assert.Equal(t, expected, child3.GetAncestors())
}

func TestNode_GetAncestorsAndSelf(t *testing.T) {
	child3 := tree.NewNode(3, nil)
	child2 := tree.NewNode(2, &[]tree.NodeIFace{child3})
	child1 := tree.NewNode(1, &[]tree.NodeIFace{child2})
	root := tree.NewNode(1, &[]tree.NodeIFace{child1})

	expected := []tree.NodeIFace{root, child1, child2, child3}
	assert.Equal(t, expected, child3.GetAncestorsAndSelf())
}

func TestNode_GetSiblings(t *testing.T) {
	child1 := tree.NewNode(1, nil)
	child2 := tree.NewNode(2, nil)
	child3 := tree.NewNode(3, nil)
	_ = tree.NewNode(1, &[]tree.NodeIFace{child1, child2, child3})

	expected := []tree.NodeIFace{child1, child2}
	assert.Equal(t, expected, child3.GetSiblings())
}

func TestNode_GetSiblingsAndSelf(t *testing.T) {
	child1 := tree.NewNode(1, nil)
	child2 := tree.NewNode(2, nil)
	child3 := tree.NewNode(3, nil)
	_ = tree.NewNode(1, &[]tree.NodeIFace{child1, child2, child3})

	expected := []tree.NodeIFace{child1, child2, child3}
	assert.Equal(t, expected, child3.GetSiblingsAndSelf())
}

func TestNode_IsRoot(t *testing.T) {
	root := tree.NewNode(nil, nil)
	assert.True(t, root.IsRoot())

	child := tree.NewNode(nil, nil)
	root = root.AddChild(child)
	assert.False(t, child.IsRoot())
}

func TestNode_IsLeaf(t *testing.T) {
	root := tree.NewNode(nil, nil)
	assert.True(t, root.IsLeaf())

	child := tree.NewNode(nil, nil)
	root = root.AddChild(child)
	assert.False(t, root.IsLeaf())
}

func TestNode_IsChild(t *testing.T) {
	root := tree.NewNode(nil, nil)
	child := tree.NewNode(nil, nil)
	assert.False(t, child.IsChild())

	root = root.AddChild(child)
	assert.True(t, child.IsLeaf())
}

func TestNode_GetDepth(t *testing.T) {
	child1 := tree.NewNode(1, nil)
	child2 := tree.NewNode(2, nil)
	child3 := tree.NewNode(3, nil)
	child4 := tree.NewNode(4, nil)
	root := tree.NewNode(1, &[]tree.NodeIFace{child1, child2, child3})
	child3 = child3.AddChild(child4)

	assert.Equal(t, 1, child1.GetDepth())
	assert.Equal(t, 0, root.GetDepth())
	assert.Equal(t, 2, child4.GetDepth())
}

func TestNode_GetHeight(t *testing.T) {
	child1 := tree.NewNode(1, nil)
	child2 := tree.NewNode(2, nil)
	child3 := tree.NewNode(3, nil)
	child4 := tree.NewNode(4, nil)
	root := tree.NewNode(nil, &[]tree.NodeIFace{child1, child2, child3})
	child3 = child3.AddChild(child4)

	assert.Equal(t, 0, child1.GetHeight())
	assert.Equal(t, 2, root.GetHeight())
	assert.Equal(t, 1, child3.GetHeight())
}

func TestNode_GetSize(t *testing.T) {
	child1 := tree.NewNode(1, nil)
	child2 := tree.NewNode(2, nil)
	child3 := tree.NewNode(3, nil)
	child4 := tree.NewNode(4, nil)
	child5 := tree.NewNode(4, nil)
	root := tree.NewNode(nil, &[]tree.NodeIFace{child1, child2, child3})
	child3 = child3.AddChild(child4).AddChild(tree.NewNode(nil, nil))
	child4 = child4.AddChild(child5)
	child5 = child5.AddChild(tree.NewNode(nil, nil)).AddChild(tree.NewNode(nil, nil))

	assert.Equal(t, 9, root.GetSize())
	assert.Equal(t, 3, child5.GetSize())
	assert.Equal(t, 4, child4.GetSize())
	assert.Equal(t, 6, child3.GetSize())
	assert.Equal(t, 1, child2.GetSize())
}
