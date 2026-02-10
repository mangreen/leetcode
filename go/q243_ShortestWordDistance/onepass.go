package q243_ShortestWordDistance

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	minDist := len(wordsDict)
	idx1, idx2 := -1, -1

	for i, w := range wordsDict {
		switch w {
			case word1:
				idx1 = i
			case word2:
				idx2 = i
		}

		if idx1!=-1 && idx2!=-1 {
			dist := abs(idx1 - idx2)
			minDist = min(minDist, dist) 
		}
	}

	return minDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}	

	return x
}