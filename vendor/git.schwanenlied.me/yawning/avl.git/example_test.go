// example_test.go - AVL tree example.
//
// To the extent possible under law, Yawning Angel has waived all copyright
// and related or neighboring rights to avl, using the Creative
// Commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package avl

import "fmt"

func CompareIntegers(a, b interface{}) int {
	// Returns < 0, 0, > 1 if a < b, a == b, a > b respectively.
	return a.(int) - b.(int)
}

func Example() {
	// Create a new tree that will store integers.
	tree := New(CompareIntegers)

	// Insert a handful of integers in random order.
	s := []int{5, 2, 6, 3, 1, 4}
	for _, i := range s {
		tree.Insert(i)
	}

	// Traverse the tree forward in-order.
	forwardInOrder := make([]int, 0, len(s))
	tree.ForEach(Forward, func(node *Node) bool {
		forwardInOrder = append(forwardInOrder, node.Value.(int))
		return true
	})

	fmt.Println(forwardInOrder)

	// Traverse the tree backward using an interator.
	backwardInOrder := make([]int, 0, len(s))
	iter := tree.Iterator(Backward)
	for node := iter.First(); node != nil; node = iter.Next() {
		backwardInOrder = append(backwardInOrder, node.Value.(int))

		// It is safe to remove the current node while iterating.
		tree.Remove(node)
	}

	fmt.Println(backwardInOrder)

	// The tree is empty after the Remove() calls.
	fmt.Println(tree.Len())

	// Output: [1 2 3 4 5 6]
	// [6 5 4 3 2 1]
	// 0
}
