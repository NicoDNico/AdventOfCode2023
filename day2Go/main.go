package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	dat, err := os.ReadFile("C:/Users/nicod/Documents/GitHub/AdventOfCode/input2.txt")
	check(err)
	text := string(dat)
	var textArray = strings.Split(text, "\n")
	part2(textArray)
}

func part1(texto []string) {
	var diceInGame [3]int
	diceInGame[0] = 12 //red
	diceInGame[1] = 13 //green
	diceInGame[2] = 14 //blue
	var suma int64
	for _, element := range texto {
		temp := (strings.Split(element, ":"))
		gameId, err := strconv.Atoi(temp[0][5:])
		check(err)
		gamePlays := temp[1]
		gameRounds := strings.Split(gamePlays, ";")
		fmt.Println(gameId)
		var possible bool = true
		for _, element := range gameRounds {
			fmt.Println(strings.Split(element, ","))
			dices := strings.Split(element, ",")
			for _, dice := range dices {
				if strings.Contains(string(dice), "red") {
					re := regexp.MustCompile("[0-9]+")
					temp, err := strconv.Atoi(re.FindAllString(dice, -1)[0])
					check(err)
					fmt.Printf("Red dice is: %v \n", temp)
					if temp > diceInGame[0] {
						possible = false
					}
				} else if strings.Contains(string(dice), "green") {
					re := regexp.MustCompile("[0-9]+")
					temp, err := strconv.Atoi(re.FindAllString(dice, -1)[0])
					check(err)
					fmt.Printf("green dice is: %v \n", temp)
					if temp > diceInGame[1] {
						possible = false
					}
				} else if strings.Contains(string(dice), "blue") {
					re := regexp.MustCompile("[0-9]+")
					temp, err := strconv.Atoi(re.FindAllString(dice, -1)[0])
					check(err)
					fmt.Printf("blue dice is: %v \n", temp)
					if temp > diceInGame[2] {
						possible = false
					}
				}
			}
		}
		if possible {
			suma += int64(gameId)
		}
	}
	fmt.Println(suma)
}
func part2(texto []string) {
	var suma int
	for _, element := range texto {
		temp := (strings.Split(element, ":"))
		gamePlays := temp[1]
		gameRounds := strings.Split(gamePlays, ";")
		var diceInRound [3]int
		for _, element := range gameRounds {
			dices := strings.Split(element, ",")
			for _, dice := range dices {
				if strings.Contains(string(dice), "red") {
					re := regexp.MustCompile("[0-9]+")
					temp, err := strconv.Atoi(re.FindAllString(dice, -1)[0])
					check(err)
					if temp > diceInRound[0] {
						diceInRound[0] = temp
					}
				} else if strings.Contains(string(dice), "green") {
					re := regexp.MustCompile("[0-9]+")
					temp, err := strconv.Atoi(re.FindAllString(dice, -1)[0])
					check(err)
					if temp > diceInRound[1] {
						diceInRound[1] = temp
					}
				} else if strings.Contains(string(dice), "blue") {
					re := regexp.MustCompile("[0-9]+")
					temp, err := strconv.Atoi(re.FindAllString(dice, -1)[0])
					check(err)
					if temp > diceInRound[2] {
						diceInRound[2] = temp
					}
				}
			}
		}
		// fmt.Println(diceInRound)
		power := diceInRound[0] * diceInRound[1] * diceInRound[2]
		fmt.Printf("the power is: %v \n", power)
		suma = suma + power
	}
	fmt.Println(suma)
}
