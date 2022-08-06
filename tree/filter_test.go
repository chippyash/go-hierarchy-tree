package tree_test

import (
	"github.com/chippyash/go-hierarchy-tree/tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 *    root
 *    /|\
 *   a b c
 *  /| |
 * d e f
 */
func TestFilterVisitor_FilterTreeWithSimpleValueNodes(t *testing.T) {
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
	sut := tree.NewFilterVisitor(func(n tree.NodeIFace) bool {
		return n.GetValue() == "e"
	})
	expected := []tree.NodeIFace{e}
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
func TestFilterVisitor_FilterTreeWithStructValueNodes(t *testing.T) {
	type testStruct struct {
		name  string
		value int
	}
	root := tree.NewNode(&testStruct{"root", 0}, nil)
	a := tree.NewNode(&testStruct{"a", 2}, nil)
	b := tree.NewNode(&testStruct{"b", 1}, nil)
	c := tree.NewNode(&testStruct{"c", 4}, nil)
	d := tree.NewNode(&testStruct{"d", 3}, nil)
	e := tree.NewNode(&testStruct{"e", 1}, nil)
	f := tree.NewNode(&testStruct{"f", 0}, nil)
	root = root.AddChild(a).AddChild(b).AddChild(c)
	a = a.AddChild(d).AddChild(e)
	b = b.AddChild(f)
	sut := tree.NewFilterVisitor(func(n tree.NodeIFace) bool {
		return n.GetValue().(*testStruct).value > 2
	})
	expected := []tree.NodeIFace{d, c}
	actual := root.Accept(sut)
	assert.Equal(t, expected, actual)
}
