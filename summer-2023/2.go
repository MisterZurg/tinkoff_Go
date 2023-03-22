package main

import "fmt"

func main() {
	var n, m, k uint16
	fmt.Scan(&n, &m, &k)

	programsPerMin := n / (m / k)

	fmt.Println(programsPerMin)
}
