package main

import (
	"bufio"
	"fmt"
	"strconv"

	"os"
)

func main() {
	frequency := 0
	frequencyTwice := 0
	var value []int
	scan := true
	var m map[int]int
	m = make(map[int]int)
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)
		frequency += i
		if _, ok := m[frequency]; ok {
			frequencyTwice = frequency
		}
		m[frequency] = i
		value = append(value, i)
	}
	fmt.Printf("result: %d\n", frequency)
	for scan {
		for _, v := range value {
			frequency += v
			if _, ok := m[frequency]; ok {
				scan = false
				frequencyTwice = frequency
				break
			}
		}
	}

	fmt.Printf("frequencyTwice: %d\n", frequencyTwice)
}
