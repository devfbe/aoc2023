package day2

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var limits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type game struct {
	index   int
	isValid bool
}

// Game 15: 15 blue, 10 red, 3 green; 9 green, 6 red, 11 blue; 3 green, 8 red, 5 blue; 12 green, 6 red, 16 blue; 11 red, 9 green, 15 blue
func parseGame(rawLine string) game {
	gamePartAndRemaining := strings.Split(rawLine, ":")
	gameIdx, err := strconv.Atoi(strings.Split(gamePartAndRemaining[0], " ")[1])
	if err != nil {
		panic(err)
	}
	isValid := true
	cubeSets := strings.Split(gamePartAndRemaining[1], ";")
game:
	for _, cubeSet := range cubeSets {
		log.Println(cubeSet)
		numAndColors := strings.Split(cubeSet, ",")
		for _, numAndColor := range numAndColors {
			numAndColorSplit := strings.Split(strings.TrimSpace(numAndColor), " ")
			num, err := strconv.Atoi(numAndColorSplit[0])
			if err != nil {
				panic(err)
			}
			color := numAndColorSplit[1]
			if num > limits[color] {
				isValid = false
				break game
			}
		}

	}

	g := game{
		index:   gameIdx,
		isValid: isValid,
	}

	log.Printf("%v\n", g)
	return g
}

func Run() {
	split := strings.Split(input, "\n")
	sum := 0
	for _, line := range split {
		if len(line) > 0 {
			game := parseGame(line)
			if game.isValid {
				sum += game.index
			}
		}
	}
	log.Printf("Day 2 result: %d", sum)
}
