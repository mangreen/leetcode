package q269_AlienDictionary

import (

)

func alienOrder(words []string) string {
    // 初始化圖和入度
    g := make(map[byte][]byte)
    inDegree := make(map[byte]int)
    m := make(map[byte]bool)
    
    // 收集所有字符
    for _, w := range words {
        for i := range w {
            m[w[i]] = true
        }
    }
    
    // 構建圖
    for i := range words[:len(words)-1] {
        w1, w2 := words[i], words[i+1]
		// 非法前綴（長字在前）
        if len(w1)>len(w2) && w1[:len(w2)]==w2 {
            return ""
        }
        
        for j:=0; j<len(w1) && j<len(w2); j++ {
            if w1[j] != w2[j] {
                g[w1[j]] = append(g[w1[j]], w2[j])
                inDegree[w2[j]]++

                break
            }
        }
    }
    
    // Kahn 算法：拓撲排序
    q := []byte{}
    for b := range m {
        if inDegree[b] == 0 {
            q = append(q, b)
        }
    }
    
    ans := []byte{}

    for len(q) > 0 {
        for _, cur := range q {
            q = q[1:]
            
            ans = append(ans, cur)
        
            for _, nxt := range g[cur] {
                inDegree[nxt]--
                if inDegree[nxt] == 0 {
                    q = append(q, nxt)
                }
            }
        }
    }
    
    // 檢查是否所有字符都包含在結果中（檢測環）
    if len(ans) != len(m) {
        return ""
    }
    
    return string(ans)
}
/*
ex1.
["wrt", "wrf", "er", "ett", "rftt"]

"wrt" vs "wrf"	
w=w, r=r, t≠f	
t, f	
t → f

"wrf" vs "er"	
w≠e
w, e	
w → e

"er" vs "ett"	
e=e, r≠t	
r, t	
r → t

"ett" vs "rftt"	
e≠r	
e, r	
e → r

w → e → r → t → f
ans="wertf"

ex2.
[z, x, z]
z ↴  
⬑ x
inD=[z:1, x:1]
有環 
找不到inDegree為0的byte
ans=""

ex3.
[abc, ab]
非法前綴（長字在前）
ans=""
*/
