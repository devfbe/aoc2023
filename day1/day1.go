package day1

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func Run() {
	split := strings.Split(input, "\n")
	sum := 0
	for _, s := range split {
		if len(s) > 0 {
			firstNum := ""
			lastNum := ""
			for _, c := range s {
				if unicode.IsDigit(c) {
					if firstNum == "" {
						firstNum = string(c)
					}
					lastNum = string(c)
				}
			}
			num := fmt.Sprintf("%s%s", firstNum, lastNum)
			log.Printf("input: %s - firstNum: %s, lastNum: %s (num: %s%s)", s, firstNum, lastNum, firstNum, lastNum)
			asInt, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			sum += asInt
			log.Println(s)
		}
	}
	log.Printf("Result: %d", sum)
}
