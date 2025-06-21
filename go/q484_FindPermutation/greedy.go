package q484_FindPermutation

func findPermutation(s string) []int {
	N := len(s)

    ans := make([]int, N+1)
    for i := range ans {
        ans[i] = i + 1
    }

    i := 0
    for i < N {
        j := i

        for j<N && s[j]=='D' {
            j++
        }

        reverse(ans, i, j)

        i = max(i+1, j)
    }

    return ans
}

func reverse(arr []int, i, j int) {
    for i < j {
        arr[i], arr[j] = arr[j], arr[i]

        i++
        j--
    }
}

/*
ex.
s="DDI"
   i
a=[1,2,3,4]
       j

       i
a=[3,2,1,4]
       j

         i
a=[3,2,1,4]
       j
*/
