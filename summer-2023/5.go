package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readPrefixIntSlice(in *bufio.Reader) []int {
	var end = true
	byteStr := make([]byte, 0)
	for end {
		temp, err, _ := in.ReadLine()
		byteStr = append(byteStr, temp...)
		end = err
	}
	finalStr := strings.Fields(string(byteStr))
	res := make([]int, 0, len(finalStr)/2+1)
	res = append(res, 0)
	for i := 1; i <= len(finalStr); i++ {
		val, _ := strconv.Atoi(finalStr[i-1])
		res = append(res, res[i-1]+val)
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	temp, _, _ := in.ReadLine()
	total, _ := strconv.Atoi(string(temp))
	sl := readPrefixIntSlice(in)
	var res = 0
	for i := 0; i < total+1; i++ {
		for j := i + 1; j < total+1; j++ {
			if sl[j] == sl[i] {
				res += total - j + 1
			}
		}
	}
	out.WriteString(strconv.Itoa(res))
}
