package q244_ShortestWordDistanceII

type WordDistance struct {
	dict map[string][]int
}

func Constructor(wordsDict []string) WordDistance {
	m := make(map[string][]int)
	for i, w := range wordsDict {
		m[w] = append(m[w], i)
	}	

	return WordDistance{dict: m}
}

func (wd *WordDistance) Shortest(word1 string, word2 string) int {
	idxs1 := wd.dict[word1]
	idxs2 := wd.dict[word2]
	i, j := 0, 0
	minDist := len(wd.dict)

	for i < len(idxs1) && j < len(idxs2) {
		dist := abs(idxs1[i] - idxs2[j])
		minDist = min(minDist, dist)
		
		// idx小的往后移，以期望获得更小的distance
		if idxs1[i] < idxs2[j] {
			i++
		} else {
			j++
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
