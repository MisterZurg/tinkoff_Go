package main

import (
	"fmt"
)

func main() {
	var h1, h2, h3, h4 int
	fmt.Scan(&h1, &h2, &h3, &h4)
	biggus := h1 >= h2 && h2 >= h3 && h3 >= h4
	smallus := h1 <= h2 && h2 <= h3 && h3 <= h4

	if biggus || smallus {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
