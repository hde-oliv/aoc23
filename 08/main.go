package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseLine(line string, m map[string][]string) map[string][]string {
	if line == "" {
		return m
	}

	s := strings.Split(line, " = ")
	k := s[0]

	d := strings.Split(s[1], ", ")
	left, _ := strings.CutPrefix(d[0], "(")
	right, _ := strings.CutSuffix(d[1], ")")

	d = nil

	d = append(d, left)
	d = append(d, right)

	m[k] = d

	return m
}

func parseFirstLine(line string) []string {
	return strings.Split(line, "")
}

func solveDirections(m map[string][]string, d []string) {
	var steps int

	k := "AAA"
	v := m[k]

	for {
		if d[steps%len(d)] == "L" {
			k = v[0]
		} else {
			k = v[1]
		}
		v = m[k]
		steps++

		if k == "ZZZ" {
			fmt.Println(steps)
			return
		}
	}
}

type Coordinate struct {
	key   string
	value []string
}

func createCoordinates(m map[string][]string) []*Coordinate {
	ghosts := make([]*Coordinate, 0)

	for k, v := range m {
		if k[len(k)-1] == 'A' {
			ghost := Coordinate{key: k, value: v}
			ghosts = append(ghosts, &ghost)
		}
	}

	return ghosts
}

func updateCoordinate(c *Coordinate, m map[string][]string, d string) {
	if d == "L" {
		c.key = c.value[0]
	} else if d == "R" {
		c.key = c.value[1]
	}
	c.value = m[c.key]
}

func solveCoordinate(c *Coordinate, m map[string][]string, d []string) int {
	var steps int

	for {
		updateCoordinate(c, m, d[steps%len(d)])
		steps++

		if c.key[len(c.key)-1] == 'Z' {
			return steps
		}
	}
}

func solveDirections2(m map[string][]string, d []string) {
	var allSteps []int

	cs := createCoordinates(m)

	for _, v := range cs {
		r := solveCoordinate(v, m, d)

		allSteps = append(allSteps, r)
	}

	fmt.Println(calculateLCM(allSteps))
}

func calculateLCM(s []int) int {
	ans := s[0]

	for i := 1; i < len(s); i++ {
		ans = (s[i] * ans) / GCD(s[i], ans)
	}

	return ans
}

func GCD(a int, b int) int {
	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make(map[string][]string)
	d := make([]string, 0)

	first := 0
	for scanner.Scan() {
		str := scanner.Text()

		if first == 0 {
			d = parseFirstLine(str)
			first = 1

			continue
		}

		m = parseLine(str, m)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part One
	solveDirections(m, d)

	// Part Two
	solveDirections2(m, d)
}
