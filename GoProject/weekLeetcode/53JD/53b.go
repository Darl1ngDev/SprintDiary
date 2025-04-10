package main

import "fmt"

func main() {
	var s string
	var count int = 0
	fmt.Scan(&s)
	pacman := "pacman"
	n := len(s)
	for i := 0; i < n; {
		j := 0
		for ; j < 6 && i+j < n; j++ {
			if s[i+j] != pacman[j] {
				break
			}
		}
		if j == 6 {
			count++
		}
		i = i + j

	}
	fmt.Println(count)
}
