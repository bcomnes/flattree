package flattree_test

import (
	"fmt"

	"github.com/bcomnes/flattree"
)

func Example() {
	var list = make([]uint64, 16, 50)

	i := flattree.Index(1, 0) // get array index for depth: 0, offset: 0
	j := flattree.Index(3, 0) // get array index for depth: 1, offset: 0

	// use these indexes to store some data

	list[i] = i
	list[j] = j
	parent := flattree.Parent(j, 0)
	list[parent] = parent

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(parent)
	fmt.Println(list)
	// Output:
	// 1
	// 7
	// 15
	// [0 1 0 0 0 0 0 7 0 0 0 0 0 0 0 15]
}
