package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	windowSize := 4
	for i := 0; i < len(s)-windowSize; i++ {
		chars := make(map[string]int)

		for j := 0; j < windowSize; j++ {
			chars[string(s[j])]++
		}
		if chars["a"] > 0 && chars["b"] > 0 && chars["c"] > 0 && chars["d"] > 0 {
			fmt.Println(windowSize)
			return
		}

		for j := windowSize; j < len(s); j++ {
			chars[string(s[j-windowSize])]--
			chars[string(s[j])]++
			if chars["a"] > 0 && chars["b"] > 0 && chars["c"] > 0 && chars["d"] > 0 {
				fmt.Println(windowSize)
				return
			}
		}
		windowSize++
	}
	fmt.Println("-1")
}

/*
7
abcabac

*/
