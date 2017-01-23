package search

import "github.com/calvn/al-go/tree"

// BFS performs breadth-first-search traversal and return the ordered values traversed tree
func BFS(t *tree.BTree) []int {
	queue := []*tree.BTree{}
	result := []int{}

	queue = append(queue, t)

	for len(queue) != 0 {
		current := queue[0]
		// Assign value to result
		result = append(result, current.Value)

		// Dequeue zeroth value
		queue = queue[1:]

		// Add left and right to queue
		if current.LTree != nil {
			queue = append(queue, current.LTree)
		}

		if current.RTree != nil {
			queue = append(queue, current.RTree)
		}
	}

	return result
}
