package main

import "fmt"

func boring(m map[int]int) bool {
	myMap := make(map[int]int)
	for _, v := range m {
		if v != 0 {
			myMap[v]++ // freq - number of numbers of that freq
		}
	}

	if len(myMap) > 2 {
		return false
	}

	if myMap[1] > 0 && len(myMap) == 1 { // all nums are of freq 1
		return true
	}

	if len(myMap) == 1 {
		return false
	}

	if myMap[1] == 1 && len(myMap) == 2 {
		return true
	}

	var a [2][2]int
	i := 0
	for key, val := range myMap {
		a[i][0] = key
		a[i][1] = val
		i++
	}

	var biggus, smallus [2]int
	if a[0][0] < a[1][0] {
		smallus = a[0]
		biggus = a[1]
	} else {
		smallus = a[1]
		biggus = a[0]
	}

	if biggus[0]-smallus[0] == 1 && biggus[1] == 1 {
		return true
	}
	return false
}

func main() {
	n, cur := 0, 0
	fmt.Scan(&n)

	a := make([]int, n)
	m := make(map[int]int)

	for i := 0; i < n; i++ {
		fmt.Scan(&cur)
		a[i] = cur
		m[cur]++ // freq of each number
	}

	for i := len(a) - 1; i > 1; i-- {
		if boring(m) {
			fmt.Print(i + 1)
			return
		}

		pooped := a[i]
		a = a[:i] // pop
		m[pooped]--
	}
}
