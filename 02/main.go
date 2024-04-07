package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const BLUE = 14
const GREEN = 13
const RED = 12

type Game struct {
	number int
	blue   int
	green  int
	red    int
}

func parseHandLine(line string, game *Game) {
	colors := strings.Split(line, ", ")

	for _, v := range colors {
		colorSplit := strings.Split(v, " ")
		colorName := colorSplit[len(colorSplit)-1]
		colorNumber, _ := strconv.Atoi(colorSplit[0])

		switch colorName {
		case "green":
			if colorNumber > game.green {
				game.green = colorNumber
			}
		case "red":
			if colorNumber > game.red {
				game.red = colorNumber
			}
		case "blue":
			if colorNumber > game.blue {
				game.blue = colorNumber
			}
		}
	}
}

func parseGameLine(str string) *Game {
	firstSplit := strings.Split(str, ": ")
	number, _ := strconv.Atoi(strings.Split(firstSplit[0], " ")[1])

	game := &Game{
		number: number,
		blue:   0,
		green:  0,
		red:    0,
	}

	listOfHands := strings.Split(firstSplit[1], "; ")
	for _, v := range listOfHands {
		parseHandLine(v, game)
	}

	return game
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	gameList := make([]Game, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()

		game := parseGameLine(str)
		gameList = append(gameList, *game)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	total := 0
	for _, v := range gameList {

		if v.blue > BLUE {
			continue
		} else if v.green > GREEN {
			continue
		} else if v.red > RED {
			continue
		} else {
			total += v.number
		}
	}

	fmt.Println(total)

	power := 0
	for _, v := range gameList {
		sum := v.blue * v.green * v.red
		power += sum
	}

	fmt.Println(power)
}
