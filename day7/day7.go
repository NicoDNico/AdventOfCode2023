package day7

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func day7() {
	dat, err := os.ReadFile("C:/Users/nicod/Documents/GitHub/AdventOfCode/Inputs/input4.txt")
	check(err)
	text := string(dat)
	var textArray = strings.Split(text, "\n")
	part1(textArray)
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
	if e == nil {
	}
}

type data struct {
	cards []string
	bid   int
	poker string
	e     error
}

func part1(textArray []string) int {
	var playedHands [][]string
	for _, line := range textArray {
		playedHands = append(playedHands, strings.Split(line, " "))
	}
	for _, play := range playedHands[:6] {
		hand := strings.Split(play[0], "")
		sort.Slice(hand, func(i, j int) bool {
			return hand[i] < hand[j]
		})
		slices.Compact[[]string, string](hand)
		bid, e := strconv.Atoi(play[1])
		check(e)
		poker := data{
			cards: hand,
			bid:   bid,
		}
		if len(hand) == 5 {
			poker.poker = "high"
		}
		fmt.Println(poker)
	}
	return 0
}
