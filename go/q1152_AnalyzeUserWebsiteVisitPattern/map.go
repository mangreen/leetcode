package q1151_MinimumSwapsToGroupAll1sTogether

import (
	"sort"
	"strings"
)

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	dic := map[string][]pair{}
	for i := range username {
		dic[username[i]] = append(dic[username[i]], pair{timestamp[i], website[i]})
	}

	count := map[string]int{}
	var res string
	maxCount := 0
	
	for _, visits := range dic {
		m := len(visits)
		if m < 3 {
			continue
		}
		
		sort.Slice(visits, func(i, j int) bool {
			return visits[i].time < visits[j].time
		})
		
		seen := map[string]bool{}

		for i := 0; i < m-2; i++ {
			for j := i + 1; j < m-1; j++ {
				for k := j + 1; k < m; k++ {
					pattern := visits[i].website + "," + visits[j].website + "," + visits[k].website
					if !seen[pattern] {
						count[pattern]++
						seen[pattern] = true
					}
				}
			}
		}
		
		for pattern, c := range count {
			if c > maxCount || (c == maxCount && (res == "" || pattern < res)) {
				maxCount = c
				res = pattern
			}
		}
	}

	return strings.Split(res, ",")
}

type pair struct {
	time    int
	website string
}
