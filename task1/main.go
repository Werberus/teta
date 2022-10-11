package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(f("sasdgassd"))

}

func f(str string) (outStr string) {
	m := make(map[rune]int)
	for _, r := range str {
		m[r] += 1
	}
	for k, v := range m {
		outStr += string(k) + strconv.Itoa(v)
	}
	return
}
