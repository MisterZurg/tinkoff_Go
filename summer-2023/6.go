package main

import (
	"fmt"
	"sort"
)

func main() {
	var learners, gradeLimit, med int
	var max []int

	fmt.Scan(&learners, &gradeLimit)
	medianSl := make([]int, learners)
	for i := 0; i < learners; i++ {
		var tmp1, tmp2 int
		fmt.Scan(&tmp1, &tmp2)
		med += tmp1
		medianSl[i] = tmp1
		max = append(max, tmp2)
	}
	lasting := gradeLimit - med
	sort.Ints(medianSl)

	divineConsent := lasting / (learners/2 + 1)
	for i := learners/2 + 1; i < learners; i++ {
		if medianSl[i] < max[i] && medianSl[i]+divineConsent < max[i] {
			medianSl[i] += divineConsent
		}
	}
	fmt.Println(medianSl[learners/2+1])
}
