// avl_test.go - AVL tree tests.
//
// To the extent possible under law, Yawning Angel has waived all copyright
// and related or neighboring rights to avl, using the Creative
// Commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package avl

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAVLTree(t *testing.T) {
	require := require.New(t)

	tree := New(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	require.Equal(0, tree.Len(), "Len(): empty")
	require.Nil(tree.First(), "First(): empty")
	require.Nil(tree.Last(), "Last(): empty")

	iter := tree.Iterator(Forward)
	require.Nil(iter.First(), "Iterator: First(), empty")
	require.Nil(iter.Next(), "Iterator: Next(), empty")

	// Test insertion.
	const nrEntries = 1024
	insertedMap := make(map[int]*Node)
	for len(insertedMap) != nrEntries {
		v := rand.Int()
		if insertedMap[v] != nil {
			continue
		}
		insertedMap[v] = tree.Insert(v)
		tree.validate(require)
	}
	require.Equal(nrEntries, tree.Len(), "Len(): After insertion")
	tree.validate(require)

	// Ensure that all entries can be found.
	for k, v := range insertedMap {
		require.Equal(v, tree.Find(k), "Find(): %v", k)
		require.Equal(k, v.Value, "Find(): %v Value", k)
	}

	// Test the forward/backward iterators.
	fwdInOrder := make([]int, 0, nrEntries)
	for k := range insertedMap {
		fwdInOrder = append(fwdInOrder, k)
	}
	sort.Ints(fwdInOrder)
	require.Equal(fwdInOrder[0], tree.First().Value, "First(), full")
	require.Equal(fwdInOrder[nrEntries-1], tree.Last().Value, "Last(), full")

	revInOrder := make([]int, 0, nrEntries)
	for i := len(fwdInOrder) - 1; i >= 0; i-- {
		revInOrder = append(revInOrder, fwdInOrder[i])
	}

	iter = tree.Iterator(Forward)
	visited := 0
	for node := iter.First(); node != nil; node = iter.Next() {
		v, idx := node.Value.(int), visited
		require.Equal(fwdInOrder[visited], v, "Iterator: Forward[%v]", idx)
		require.Equal(node, iter.Get(), "Iterator: Forward[%v]: Get()", idx)
		visited++
	}
	require.Equal(nrEntries, visited, "Iterator: Forward: Visited")

	iter = tree.Iterator(Backward)
	visited = 0
	for node := iter.First(); node != nil; node = iter.Next() {
		v, idx := node.Value.(int), visited
		require.Equal(revInOrder[idx], v, "Iterator: Backward[%v]", idx)
		require.Equal(node, iter.Get(), "Iterator: Backward[%v]: Get()", idx)
		visited++
	}
	require.Equal(nrEntries, visited, "Iterator: Backward: Visited")

	// Test the forward/backward ForEach.
	forEachValues := make([]int, 0, nrEntries)
	forEachFn := func(n *Node) bool {
		forEachValues = append(forEachValues, n.Value.(int))
		return true
	}
	tree.ForEach(Forward, forEachFn)
	require.Equal(fwdInOrder, forEachValues, "ForEach: Forward")

	forEachValues = make([]int, 0, nrEntries)
	tree.ForEach(Backward, forEachFn)
	require.Equal(revInOrder, forEachValues, "ForEach: Backward")

	// Test removal.
	for i, idx := range rand.Perm(nrEntries) { // In random order.
		v := fwdInOrder[idx]
		node := tree.Find(v)
		require.Equal(v, node.Value, "Find(): %v (Pre-remove)", v)

		tree.Remove(node)
		require.Equal(nrEntries-(i+1), tree.Len(), "Len(): %v (Post-remove)", v)
		tree.validate(require)

		node = tree.Find(v)
		require.Nil(node, "Find(): %v (Post-remove)", v)
	}
	require.Equal(0, tree.Len(), "Len(): After removal")
	require.Nil(tree.First(), "First(): After removal")
	require.Nil(tree.Last(), "Last(): After removal")

	// Refill the tree.
	for _, v := range fwdInOrder {
		tree.Insert(v)
	}

	// Test that removing the node doesn't break the iterator.
	iter = tree.Iterator(Forward)
	visited = 0
	for node := iter.Get(); node != nil; node = iter.Next() { // Omit calling First().
		v, idx := node.Value.(int), visited
		require.Equal(fwdInOrder[idx], v, "Iterator: Forward[%v] (Pre-Remove)", idx)
		require.Equal(fwdInOrder[idx], tree.First().Value, "First() (Iterator, remove)")
		visited++

		tree.Remove(node)
		tree.validate(require)
	}
	require.Equal(0, tree.Len(), "Len(): After iterating removal")
}

func (t *Tree) validate(require *require.Assertions) {
	checkInvariants(require, t.root, nil)
}

func checkInvariants(require *require.Assertions, node, parent *Node) int {
	if node == nil {
		return 0
	}

	// Validate the parent pointer.
	require.Equal(parent, node.parent)

	// Validate that the balance factor is -1, 0, 1.
	require.Condition(func() bool {
		switch node.balance {
		case -1, 0, 1:
			return true
		}
		return false
	})

	// Recursively derive the height of the left and right sub-trees.
	lHeight := checkInvariants(require, node.left, node)
	rHeight := checkInvariants(require, node.right, node)

	// Validate the AVL invariant and the balance factor.
	require.Equal(int(node.balance), rHeight-lHeight)
	if lHeight > rHeight {
		return lHeight + 1
	}
	return rHeight + 1
}
