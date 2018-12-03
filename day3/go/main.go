package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	claim struct {
		x int
		y int
	}
	fabric map[claim][]string
)

func (c claim) String() string {
	return fmt.Sprintf("%d:%d", c.x, c.y)
}

func newClaim(x int, y int) claim {
	return claim{
		x: x,
		y: y,
	}
}

func (c fabric) allocateClaim(claim string, fromLeft int, fromTop int, wide int, tall int) {
	//fmt.Printf("id: %s,  left: %d, top: %d, wide: %d, tall: %d\n", claim, fromLeft, fromTop, wide, tall)
	withStart := 1 + fromLeft
	withEnd := withStart + wide
	topStart := 1 + fromTop
	topEnd := topStart + tall
	//fmt.Printf("claim: %s\n", claim)
	//fmt.Printf("withStart: %d, withEnd: %d, with: %d, fromLeft: %d\n", withStart, withEnd, wide, fromLeft)
	//fmt.Printf("topStart: %d, topend: %d, tall: %d, fromTop: %d\n", topStart, topEnd, tall, fromTop)

	for i := topStart; i < topEnd; i++ {

		//wCnt = 0
		for x := withStart; x < withEnd; x++ {
			c[newClaim(x, i)] = append(c[newClaim(x, i)], claim)
			//fmt.Printf("x")
		}
		//fmt.Printf("\n")
	}
	//fmt.Printf("t: %d, w: %d \n", tCnt, wCnt)
}

func (c *fabric) sumNoUniqueClaimedinches() int {
	sum := 0
	for _, v := range *c {
		if len(v) > 1 {
			//fmt.Printf("claim: %v\n", v)
			sum++
		}
	}
	return sum
}

func (c *fabric) loadInput() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		firstSplit := strings.Split(line, "@")
		id := strings.Trim(firstSplit[0], " ")
		secondSplit := strings.Split(firstSplit[1], ":")
		leftTop := strings.TrimSpace(secondSplit[0])
		left, _ := strconv.Atoi(strings.Split(leftTop, ",")[0])
		top, _ := strconv.Atoi(strings.Split(leftTop, ",")[1])

		wideTall := strings.TrimSpace(secondSplit[1])
		wide, _ := strconv.Atoi(strings.Split(wideTall, "x")[0])
		tall, _ := strconv.Atoi(strings.Split(wideTall, "x")[1])
		//fmt.Printf("id: %s,  left: %s, top: %s, wide: %s, tall: %s\n", id, left, top, wide, tall)
		c.allocateClaim(id, left, top, wide, tall)
	}
}

func main() {

	f := make(fabric)
	//f[newClaim(1, 2)] = append(f[newClaim(1, 2)], "a")
	f.loadInput()

	//fmt.Printf("i: %s\n", c)
	//fmt.Printf("v: %v\n", f)

	//fmt.Printf("v: %d\n", len(f))

	fmt.Printf("result: %d\n", f.sumNoUniqueClaimedinches())
	uniqueClaim := make(map[string]bool)
	for _, v := range f {
		if len(v) == 1 {
			if _, ok := uniqueClaim[v[0]]; !ok {
				uniqueClaim[v[0]] = true
			}
		} else {
			for _, x := range v {
				if _, ok := uniqueClaim[x]; ok {
					uniqueClaim[x] = false
				}
			}
		}
	}
	for key, value := range uniqueClaim {
		if value {
			fmt.Printf("unique id: %s \n", strings.Trim(key, "#"))
			break
		}
	}
	//fmt.Printf("v: %v\n", f)
}

/*

f[newClaim(1, 2)] = append(f[newClaim(1, 2)], 1)
	f[newClaim(1, 2)] = append(f[newClaim(1, 2)], 2)

	c := claim{
		x: 1,
		y: 2,
	}
*/
