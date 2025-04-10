package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	pre := make([]int, n)
	d := 0
	for i, c := range s {
		d += int((c%2)*2 - 1)
		pre[i] += d
	}
	hashmap := make(map[int]int)
	ans := 0
	for i, v := range pre {
		value, ok := hashmap[v]
		if ok {
			ans = max(ans, i-value)
		} else {
			hashmap[v] = i
		}
	}
	fmt.Println(ans)
}
