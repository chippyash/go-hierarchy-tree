package tree

/**
 * Simple Double Entry Accounting V3 for Go
 * Tree implementation
 * @author Ashley Kitson
 * @copyright Ashley Kitson, 2022, UK
 * @license BSD-3-Clause See LICENSE.md
 */

//FilterFunc signature for filter functions
type FilterFunc func(NodeIFace) bool

//FilterVisitor is a visitor that filters nodes
type FilterVisitor struct {
	VisitorIFace
	filter FilterFunc
}

//NewFilterVisitor returns a FilterVisitor
func NewFilterVisitor(filter FilterFunc) VisitorIFace {
	v := &FilterVisitor{filter: filter}
	return v
}

//Visit filters nodes (pre-order traversal) that match the filter and returns []NodeIFace
func (v *FilterVisitor) Visit(n NodeIFace) interface{} {
	nodes := make([]NodeIFace, 0)
	if v.filter(n) {
		nodes = append(nodes, n)
	}
	for _, child := range n.GetChildren() {
		vNodes := child.Accept(v).([]NodeIFace)
		nodes = append(nodes, vNodes...)
	}
	return nodes
}
