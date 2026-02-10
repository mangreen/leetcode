package q277_FindTheCelebrity

/**
 * The Knows API is already defined for you.
 *     Knows := func(a int, b int) bool
 */

func findCelebrity(n int) int {
	candidate := 0

	// 1st pass: find a candidate
	for i := 1; i < n; i++ {
		if knows(candidate, i) {
			candidate = i
		}
	}

	// 2nd pass: verify the candidate
	for i := range n {
		if i != candidate && (knows(candidate, i) || !knows(i, candidate)) {
			return -1
		}
	}
	return candidate
}

var graph [][]bool

func knows (a, b int) bool {
	return graph[a][b]
}
