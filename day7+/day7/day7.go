package day7

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/icza/abcsort"
)

func Day7() int64 {
	dat, err := os.ReadFile("C:/Users/nicod/Documents/GitHub/AdventOfCode/Inputs/input7.txt")
	check(err)
	text := string(dat)
	var textArray = strings.Split(text, "\n")
	return Analyze(textArray)
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
	// I hold to all cards unordered cause i will need to check later for overall order
	cards string
	bid   int64
	// i call the hand/combo poker
	poker string
}

func Analyze(textArray []string) int64 {
	// I could create a map of letter to order but i looked at the code for this package and it was better than i could do.
	// the creators of advent may say you should not use packages but a sorter is boring to code specially if i want to make it fast in anyway.
	sorter := abcsort.New("23456789TJQKA")
	// i moved the files so i could upload them to github and now the sorter doesnt work and im so angry im not doing part2.
	// it should not be that hard to make a project i can move from folder to folder and still work.
	// but somehow every programming languaje i use has difficulties with it. I hate it and im dying on this hill.
	sorter2 := abcsort.New("J23456789TQKA")
	// This is the return value
	var result int64
	var playedHands [][]string
	// I think im making the same code a different way each day.
	// this Separates unordened cards from bid value.
	for _, line := range textArray {
		playedHands = append(playedHands, strings.Split(line, " "))
	}
	// for each handInPlay
	handStorage := getCardsInfo(playedHands)
	// at the point im writing this i have used golang for only 1 week dont expect more from me.
	// i prov should have added them to a slice the moment i got the poker hand but i dont want to refactor.
	var highCards []data
	var pair []data
	var twoPair []data
	var threeOfAKind []data
	var fullHouse []data
	var poker []data
	var fiveOfAKind []data
	for _, hand := range handStorage {
		if hand.poker == "highCard" {
			highCards = append(highCards, hand)
		} else if hand.poker == "pair" {
			pair = append(pair, hand)
		} else if hand.poker == "twoPair" {
			twoPair = append(twoPair, hand)
		} else if hand.poker == "threeOfAKind" {
			threeOfAKind = append(threeOfAKind, hand)
		} else if hand.poker == "fullHouse" {
			fullHouse = append(fullHouse, hand)
		} else if hand.poker == "poker" {
			poker = append(poker, hand)
		} else {
			fiveOfAKind = append(fiveOfAKind, hand)
		}

	}
	// i need somewhere to store them to get the position of each on the overall game.
	// another thing i think i could do at the same time but i dont want to deal with it.
	sorter.Slice(highCards, func(i int) string { return highCards[i].cards })
	sorter.Slice(pair, func(i int) string { return pair[i].cards })
	sorter.Slice(twoPair, func(i int) string { return twoPair[i].cards })
	sorter.Slice(threeOfAKind, func(i int) string { return threeOfAKind[i].cards })
	sorter.Slice(fullHouse, func(i int) string { return fullHouse[i].cards })
	sorter.Slice(poker, func(i int) string { return poker[i].cards })
	sorter.Slice(fiveOfAKind, func(i int) string { return fiveOfAKind[i].cards })
	forGetingResult := concatSlice(highCards, pair, twoPair, threeOfAKind, fullHouse, poker, fiveOfAKind)
	for index, hand := range forGetingResult {
		result += (int64(index) + 1) * int64(hand.bid)
		fmt.Println(hand)
	}
	return result
}

// I wanted to make a map but i dont want to stop to understand maps on go.
func getPoker(duplicate map[string]int) string {
	if duplicate["J"] != 0 {

	} else {
		if len(duplicate) == 5 {
			return "highCard"
		} else if len(duplicate) == 4 {
			return "pair"
		} else if len(duplicate) == 3 {
			for _, i := range duplicate {
				if i == 3 {
					// i accidentally wrote treeOfAKind here and modified the code so much i forgot how any of this works
					// it took me 2Hours and solving 2 other bugs to see this.
					return "threeOfAKind"
				}
			}
			return "twoPair"
		} else if len(duplicate) == 2 {
			for _, i := range duplicate {
				if i == 3 {
					return "fullHouse"
				}
			}
			return "poker"
		} else if len(duplicate) == 1 {
			return "fiveOfAKind"
		}
	}
	return "fail"
}
func concatSlice(slices ...[]data) []data {
	var result []data
	for _, slice := range slices {
		for _, hand := range slice {
			result = append(result, hand)
		}
	}
	return result
}

func getCardsInfo(playedHands [][]string) []data {
	sorter := abcsort.New("23456789TJQKA")
	var handStorage []data
	for _, play := range playedHands {
		// unordened cards
		hand := strings.Split(play[0], "")
		// ordered cards
		sorter.Slice(hand, func(i int) string { return hand[i] })
		bid, e := strconv.Atoi(play[1])
		check(e)
		poker := data{
			cards: play[0],
			bid:   int64(bid),
		}
		// it keeps track of all duplicated cards i use this to get what type of poker hand it is.
		duplicate := make(map[string]int)
		for _, card := range hand {
			_, exist := duplicate[card]
			if exist {
				duplicate[card] += 1
			} else {
				duplicate[card] = 1 // else start counting from 1
			}

		}
		//should i do better? No, i should do worst
		poker.poker = getPoker(duplicate)
		handStorage = append(handStorage, poker)
	}
	return handStorage
}
