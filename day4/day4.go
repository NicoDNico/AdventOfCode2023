package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
	if e == nil {
	}
}

func main() {
	dat, err := os.ReadFile("C:/Users/nicod/Documents/GitHub/AdventOfCode/input4.txt")
	check(err)
	text := string(dat)
	var textArray = strings.Split(text, "\n")
	part2(textArray)
}
func part1(textarray []string) {
	var winningNumbers [][]string
	var numbersHad [][]string
	var result int
	// Separating array into the parts i want
	for _, line := range textarray {
		// Separates numbers from card counter.
		temp := strings.Split(line, ":")
		// Separates winning numbers and numbers we got.
		temp2 := strings.Split(temp[1], "|")
		winningNumbers = append(winningNumbers, strings.Split(temp2[0], " "))
		numbersHad = append(numbersHad, strings.Split(temp2[1], " "))
	}
	// checking for winners
	for cardIndex, card := range winningNumbers {
		var cardpoints = 0
		for _, numberW := range card {
			if string(numberW) != "" {
				for _, numberH := range numbersHad[cardIndex] {
					if numberH == string(numberW) {
						fmt.Println(cardIndex, numberH)
						if cardpoints != 0 {
							cardpoints = cardpoints * 2
						} else {
							cardpoints = 1
						}
					}
				}
			}
		}
		result = result + cardpoints
		// fmt.Println(cardpoints)
	}

	fmt.Println("Result sum: ", result)
}

func part2(textarray []string) {
	// All i need to solve the problem + index that i guesses that i nedded but i dont.
	type datos struct {
		index          int
		winningNumbers [][]string
		numbersHad     [][]string
		timesToRun     int
		cardsHad       int
	}
	var result int
	var cardsArray []datos
	for index, line := range textarray {
		card := datos{
			timesToRun: 1,
			cardsHad:   1,
		}
		// Separates numbers from card counter.
		temp := strings.Split(line, ":")
		card.index = index
		// Separates winning numbers and numbers we got.
		temp2 := strings.Split(temp[1], "|")
		// this avoids having to deal with runes while also storing the numbers as a whole and not as single digits
		// that i would need to put together later. Yeah i figured it out too late.
		card.winningNumbers = append(card.winningNumbers, strings.Split(temp2[0], " "))
		card.numbersHad = append(card.numbersHad, strings.Split(temp2[1], " "))
		cardsArray = append(cardsArray, card)
	}
	// The moment i saw this take more than 1 sec to compute i though of a better solution.
	// im writing it for prosperity.
	// i dont need to iterate trough the same card multiple times i just need to know
	// how many cards i have, iterate once and then multiply the number i add to the other cards
	// by the amount of current cards i have.
	// why didnt i tought of this? i did i just forgot to start each card with 1 so it
	// allways resolved to the wrong answer and when i figured it out i had already changed
	// my function to this.
	for i := 0; i < len(cardsArray); i++ {
		var card = cardsArray[i]
		for j := 0; j < card.timesToRun; j++ {
			var points int
			for _, wNumber := range card.winningNumbers[0] {
				if wNumber != "" {
					for _, hNumber := range card.numbersHad[0] {
						if wNumber == hNumber {
							points = points + 1
							cardsArray[i+points].timesToRun = cardsArray[(i+points)].timesToRun + 1
							cardsArray[i+points].cardsHad = cardsArray[(i+points)].cardsHad + 1
						}
					}
				}
			}
		}
		result = result + cardsArray[i].cardsHad
		fmt.Println(cardsArray[i])
	}
	fmt.Println("result:", result)
}

// also for prosperity i had solved the test example for a while but for some reason my code didnt work on the actual input.
// i was trying to figure out why on the first 6 lines of the actual input. To the one who made it so they are powers of 2
// i hate you. i though i was doing it wrong for sooooo long cause it didnt process that the actual answer was for the first 6 lines
// to be doubled each time.
