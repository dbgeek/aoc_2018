package main

import (
	"bufio"
	"fmt"

	"os"
)

func nrExists(v map[string]int, nr int) bool {
	for _, v := range v {
		if v == nr {
			return true
		}
	}
	return false
}

func sameChars(s1 string, s2 string) string {
	s := ""
	for p, v := range s1 {
		if string(s2[p]) == string(v) {
			s += string(v)
		}
	}
	return s
}

func compareString(s1 string, s2 string) int {
	c := 0
	for p, v := range s1 {
		if string(s2[p]) != string(v) {
			c++
		}
	}
	return c
}

func main() {
	/*frequency := 0
	frequencyTwice := 0
	var value []int
	scan := true
	var m map[string]int

	m = make(map[string]int)

	*/
	xxx := 0
	xx := 0
	var s []string
	var m map[string]int
	m = make(map[string]int)
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		m = make(map[string]int)
		s = append(s, string(line))
		for _, char := range line {
			if _, ok := m[string(char)]; ok {
				m[string(char)]++
			} else {
				m[string(char)] = 1
			}
		}
		if nrExists(m, 2) {
			xx++
		}
		if nrExists(m, 3) {
			xxx++
		}
	}
	fmt.Printf("3: %d 2: %d \n", xxx, xx)
	f := false
	for p1, v1 := range s {
		for p2, v2 := range s {
			if p1 != p2 && compareString(v1, v2) == 1 {
				//fmt.Printf("v1: %s, v2: %s \n", v1, v2)
				fmt.Printf("Same chars: %s \n", sameChars(v1, v2))
				f = true
			}
		}
		if f {
			break
		}
	}
}
