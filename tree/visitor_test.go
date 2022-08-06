package tree_test

import (
	"github.com/chippyash/go-hierarchy-tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPreOrderVisitor_ImplementsVisitorInterface(t *testing.T) {
	sut := tree.NewPreOrderVisitor()
	_, ok := sut.(tree.VisitorIFace)
	assert.True(t, ok)
}

func TestPreOrderVisitor_WalkTreeWithOneNode(t *testing.T) {
	root := tree.NewNode("A", nil)
	sut := tree.NewPreOrderVisitor()
	expected := []tree.NodeIFace{root}
	assert.Equal(t, expected, root.Accept(sut))
}

func TestPreOrderVisitor_WalkTreeWithTwoNodes(t *testing.T) {
	root := tree.NewNode("A", nil)
	child := tree.NewNode("b", nil)
	root = root.AddChild(child)
	sut := tree.NewPreOrderVisitor()
	expected := []tree.NodeIFace{root, child}
	assert.Equal(t, expected, root.Accept(sut))
}

/**
 *    root
 *    /|\
 *   a b c
 *  /| |
 * d e f
 */
func TestPreOrderVisitor_WalkTreeWithMoreNodes(t *testing.T) {
	root := tree.NewNode("root", nil)
	a := tree.NewNode("a", nil)
	b := tree.NewNode("b", nil)
	c := tree.NewNode("c", nil)
	d := tree.NewNode("d", nil)
	e := tree.NewNode("e", nil)
	f := tree.NewNode("f", nil)
	root = root.AddChild(a).AddChild(b).AddChild(c)
	a = a.AddChild(d).AddChild(e)
	b = b.AddChild(f)
	sut := tree.NewPreOrderVisitor()
	expected := []tree.NodeIFace{root, a, d, e, b, f, c}
	actual := root.Accept(sut)
	assert.Equal(t, expected, actual)
}

/**
 *    root
 *    /|\
 *   a b c
 *  /| |
 * d e f
 */
func TestPreOrderVisitor_WalkTreeWithSubTree(t *testing.T) {
	root := tree.NewNode("root", nil)
	a := tree.NewNode("a", nil)
	b := tree.NewNode("b", nil)
	c := tree.NewNode("c", nil)
	d := tree.NewNode("d", nil)
	e := tree.NewNode("e", nil)
	f := tree.NewNode("f", nil)
	root = root.AddChild(a).AddChild(b).AddChild(c)
	a = a.AddChild(d).AddChild(e)
	b = b.AddChild(f)
	sut := tree.NewPreOrderVisitor()
	expected := []tree.NodeIFace{a, d, e}
	actual := a.Accept(sut)
	assert.Equal(t, expected, actual)
}

func TestPostOrderVisitor_ImplementsVisitorInterface(t *testing.T) {
	sut := tree.NewPostOrderVisitor()
	_, ok := sut.(tree.VisitorIFace)
	assert.True(t, ok)
}

func TestPostOrderVisitor_WalkTreeWithOneNode(t *testing.T) {
	root := tree.NewNode("A", nil)
	sut := tree.NewPostOrderVisitor()
	expected := []tree.NodeIFace{root}
	assert.Equal(t, expected, root.Accept(sut))
}

func TestPostOrderVisitor_WalkTreeWithTwoNodes(t *testing.T) {
	root := tree.NewNode("A", nil)
	child := tree.NewNode("b", nil)
	root = root.AddChild(child)
	sut := tree.NewPostOrderVisitor()
	expected := []tree.NodeIFace{child, root}
	assert.Equal(t, expected, root.Accept(sut))
}

/**
 *    root
 *    /|\
 *   a b c
 *  /| |
 * d e f
 */
func TestPostOrderVisitor_WalkTreeWithMoreNodes(t *testing.T) {
	root := tree.NewNode("root", nil)
	a := tree.NewNode("a", nil)
	b := tree.NewNode("b", nil)
	c := tree.NewNode("c", nil)
	d := tree.NewNode("d", nil)
	e := tree.NewNode("e", nil)
	f := tree.NewNode("f", nil)
	root = root.AddChild(a).AddChild(b).AddChild(c)
	a = a.AddChild(d).AddChild(e)
	b = b.AddChild(f)
	sut := tree.NewPostOrderVisitor()
	expected := []tree.NodeIFace{d, e, a, f, b, c, root}
	actual := root.Accept(sut)
	assert.Equal(t, expected, actual)
}

/**
 *    root
 *    /|\
 *   a b c
 *  /| |
 * d e f
 */
func TestPostOrderVisitor_WalkTreeWithSubTree(t *testing.T) {
	root := tree.NewNode("root", nil)
	a := tree.NewNode("a", nil)
	b := tree.NewNode("b", nil)
	c := tree.NewNode("c", nil)
	d := tree.NewNode("d", nil)
	e := tree.NewNode("e", nil)
	f := tree.NewNode("f", nil)
	root = root.AddChild(a).AddChild(b).AddChild(c)
	a = a.AddChild(d).AddChild(e)
	b = b.AddChild(f)
	sut := tree.NewPostOrderVisitor()
	expected := []tree.NodeIFace{d, e, a}
	actual := a.Accept(sut)
	assert.Equal(t, expected, actual)
}

/**
 *    root
 *    /|\
 *   a b c
 *  /| |
 * d e f
 */
func TestLeafVisitor_GetLeaves(t *testing.T) {
	root := tree.NewNode("root", nil)
	a := tree.NewNode("a", nil)
	b := tree.NewNode("b", nil)
	c := tree.NewNode("c", nil)
	d := tree.NewNode("d", nil)
	e := tree.NewNode("e", nil)
	f := tree.NewNode("f", nil)
	root = root.AddChild(a).AddChild(b).AddChild(c)
	a = a.AddChild(d).AddChild(e)
	b = b.AddChild(f)
	sut := tree.NewLeafVisitor()
	expected := []tree.NodeIFace{d, e, f, c}
	actual := root.Accept(sut)
	assert.Equal(t, expected, actual)
}

func TestLeafVisitor_TheYieldOfALeafNodeIsTheNodeItself(t *testing.T) {
	root := tree.NewNode("root", nil)
	sut := tree.NewLeafVisitor()
	expected := []tree.NodeIFace{root}
	actual := root.Accept(sut)
	assert.Equal(t, expected, actual)
}
