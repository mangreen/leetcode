package q1257_SmallestCommonRegion

import ()

func findSmallestRegion(regions [][]string, region1, region2 string) string {
	m := make(map[string]string)
    for _, arr := range regions {
        for _, r := range arr[1:] {
            m[r] = arr[0]
        } 
    }
    
    used := make(map[string]bool)
    
    for region1 != "" {
        used[region1] = true
        region1 = m[region1]
    }

    for !used[region2] {
        region2 = m[region2]
    }
    
    return region2
}
/*
由上往下 235 & 236 & 1644 & 1650 
這題由下往上

    "Earth"
      / \
   "NA" "SA"
    / \    \
"USA" "CA"  "BR"

r1="USA"
r2="SA"

m={NA:Earth, SA:Earth, USA:NA, CA:NA, BR:SA}

r1="USA" ⭢ used=[USA]
r1="NA" ⭢ used=[USA, NA]
r1="Earth" ⭢ used=[USA, NA, Earth]
r1=""

r2="SA" ⭢ not in used=[USA, NA, Earth]
r2="Earth" ⭢ used["Earth"]==true ⭢ return
*/