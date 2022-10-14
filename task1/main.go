package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(f("aZЯAaaBBЯz"))
}

func f(str string) (outStr string) {
	m := make(map[string]int)
	var charSlice []string
	for _, r := range str {
		char := string(r)
		if _, ok := m[char]; !ok {
			charSlice = append(charSlice, char)
		}
		m[char] += 1
	}
	sort.Strings(charSlice)
	for _, v := range charSlice {
		outStr += v + strconv.Itoa(m[v])
	}
	return
}
