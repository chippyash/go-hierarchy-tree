# Hierarchy tree for Go
## github.com/chippyash/go-hierarchy-tree/tree

Go: 1.18

## What
Provides a basic but flexible hierarchical tree data structure together with an implementation of a 
Visitor pattern for tree manipulation.

## How

`import "github.com/chippyash/go-hierarchy-tree/tree"`

### For Development
#### Setup

- clone repository
- cd to project directory
- `go get ./...`

#### Usage
##### Create a node
```go
import "github.com/chippyash/go-hierarchy-tree/tree"

//node with no value or children
node := tree.NewNode(nil, nil)
//node with some value
node := tree.NewNode("foo", nil)
//node with children
child1 := tree.NewNode("bar", nil)
child2 := tree.NewNode("baz", nil)
children := []tree.NodeIFace{child1, child2}
node := tree.NewNode("foo", &children)
```

##### Getting and setting a node value
A node can have any value. The value of a node is an interface{}
```go
val := tree.NewNode(nil, nil).
		SetValue("foo").
		GetValue()
```

##### Adding one or more children
You can add children when you construct a node (see above) or afterwards.
```go
child1 := tree.NewNode(1, nil)
child2 := tree.NewNode(2, nil)
children := []tree.NodeIFace{child1, child2}

childNodes := tree.NewNode(0, nil).
		SetChildren(children...).
		GetChildren()

childNodes = tree.NewNode(0, nil).
AddChild(child1).
AddChild(child2).
GetChildren()
```

##### Removing child nodes
```go
node.RemoveChild(child1)
node.RemoveAllChildren()
```

##### Testing nodes
```go
node.IsLeaf()  //true if node has no children
node.IsRoot()  //true if node is the root node, i.e. no parent
node.IsChild() //true if node has a parent
```

##### Parents, siblings and ancestors
```go
node.SetParent(someOtherNode)  //set the parent to the node
node.GetParent()            //return parent of this node - nil if node is root
node.GetSiblings()          //returns nodes with same parent
node.GetSiblingsAndSelf()   //returns nodes with same parent including self
node.GetAncestors()         //parents and grandparents
node.GetAncestorsAndSelf()  //self, parents and grandparents

```
##### Traversing a tree
The tree implements the Visitor pattern Accept method.
###### Pre-order traversal
```go
/**
 *    root
 *    /|\
 *   a b c
 *  /| |
 * d e f
 */
visitor := tree.NewPreOrderVisitor()
result := root.Accept(visitor)
//returns []tree.NodeIFace{root, a, d, e, b, f, c}
```

###### Post-order traversal
```go
/**
 *    root
 *    /|\
 *   a b c
 *  /| |
 * d e f
 */
visitor := tree.NewPostOrderVisitor()
result := root.Accept(visitor)
//returns []tree.NodeIFace{d, e, a, f, b, c, root}
```

###### Gathering leaves
```go
/**
 *    root
 *    /|\
 *   a b c
 *  /| |
 * d e f
 */
visitor := tree.NewLeafVisitor()
result := root.Accept(visitor)
//returns []tree.NodeIFace{d, e, f, c}
```

##### Filtering
```go
/**
 *    root
 *    /|\
 *   a b c
 *  /| |
 * d e f
 */
visitor := tree.NewFilterVisitor(func(n tree.NodeIFace) bool {
return n.GetValue() == "e"
})
result := root.Accept(visitor)
//returns []tree.NodeIFace{e}
```
see [Filter Test](tree/filter_test.go) for other examples

##### Node information
```go
node.GetDepth()   //returns the distance from the current node to the root
node.GetHeight()  //returns the height of the tree whose root is this node
node.GetSize()    //returns the number of nodes in the tree rooted at this node
```
#### Testing

`go test ./...`

#### Before you do a PR

- Update the readme if necessary



## References

- [Github](https://github.com/chippyash/go-hierarchy-tree)
- [Based on Nicmart/Tree](https://github.com/nicmart/Tree)
