package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func findLastNumber(str string) (string, error) {
	writtenNumbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	convertedNumbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	tmpString := ""
	idx := len(str)

	for len(tmpString) != len(str) {
		newRune := str[idx-len(tmpString)-1]

		if unicode.IsDigit(rune(newRune)) {
			return string(newRune), nil
		}

		tmpString = string(newRune) + tmpString

		for i, v := range writtenNumbers {
			if idx := strings.Index(tmpString, v); idx != -1 {
				return convertedNumbers[i], nil
			}
		}

	}

	return "", errors.New("last not found")
}

func findFirstNumber(str string) (string, error) {
	writtenNumbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	convertedNumbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	writtenIndex := len(str)
	convertedIndex := -1
	numberIndex := -1

	for i, v := range writtenNumbers {
		if idx := strings.Index(str, v); idx != -1 {
			if idx < writtenIndex {
				writtenIndex = idx
				convertedIndex = i
			}
		}
	}

	for i, v := range str {
		if unicode.IsNumber(v) {
			numberIndex = i
			break
		}
	}

	if writtenIndex == len(str) {
		return string(str[numberIndex]), nil
	}

	if writtenIndex != len(str) && numberIndex != -1 {
		if writtenIndex < numberIndex {
			return convertedNumbers[convertedIndex], nil
		}
		if numberIndex < writtenIndex {
			return string(str[numberIndex]), nil
		}
	}

	return "", errors.New("first not found")
}

func getStringNumber(str string) string {
	first, err := findFirstNumber(str)
	if err != nil {
		log.Fatal(err)
	}

	last, err := findLastNumber(str)
	if err != nil {
		log.Fatal(err)
	}

	return first + last
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	numberList := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()

		stringNumber := getStringNumber(str)

		number, err := strconv.Atoi(stringNumber)
		if err != nil {
			log.Fatal(err)
		}

		numberList = append(numberList, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := 0

	for _, v := range numberList {
		result += v
	}

	fmt.Println(result)
}
