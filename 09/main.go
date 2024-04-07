package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Map(vs []string, f func(string) (int, error)) []int {
	vsm := make([]int, len(vs))

	for i, v := range vs {
		vsm[i], _ = f(v)
	}
	return vsm
}

func parseLine(s string) []int {
	return Map(strings.Split(s, " "), strconv.Atoi)
}

func validateList(l []int) bool {
	result := true

	for _, v := range l {
		if v != 0 {
			result = false
		}
	}

	return result
}

func extrapolateOne(l []int) int {
	t := make([]int, 0)
	lst := make([]int, 0)

	// slices.Reverse(l) - for part two

	lst = append(lst, l[len(l)-1])
	for !validateList(l) {
		for i := 0; i < len(l)-1; i++ {
			t = append(t, l[i+1]-l[i])
		}

		if len(t) == 0 {
			t = append(t, 0)
		}

		lst = append(lst, t[len(t)-1])

		l = nil
		l = append(l, t...)
		t = nil
	}

	var r int

	for _, v := range lst {
		r += v
	}

	return r
}

func extrapolateAll(l [][]int) int {
	var lst []int

	for _, v := range l {
		lst = append(lst, extrapolateOne(v))
	}

	var r int

	for _, v := range lst {
		r += v
	}

	return r
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var nList [][]int

	for scanner.Scan() {
		s := scanner.Text()

		nList = append(nList, parseLine(s))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	r1 := extrapolateAll(nList)

	fmt.Println(r1)
}
