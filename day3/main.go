package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
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
	dat, err := os.ReadFile("C:/Users/nicod/Documents/GitHub/AdventOfCode/input3.txt")
	check(err)
	text := string(dat)
	var textArray = strings.Split(text, "\n")
	// part1(textArray)
	// test(textArray)
	// part1_2(textArray)
	part2(textArray)

}

func part1(textarray []string) {
	var LinesNumbers [][][]int
	var LinesSymbols [][]int
	for _, line := range textarray {
		var numberStorage [][]int
		var letterIndexStorage []int
		var symbolStorage []int
		// removing the whitespaces at the end and beginning of a line also makes it so the code doesnt find the second non contiguos numbers.
		// line = strings.TrimSpace(line)
		for index, letter := range string(line) {
			letra := string(letter)
			test := regexp.MustCompile(`[^a-z 0-9 .]+`).MatchString(letra)

			// one may think that if used this but i passed when unicode.IsNumber is true i would gett all the numbers
			// right?
			// NO, you get only the first contiguos numbers on each line. idk why and i dont care im tired of trying shit.
			if letra != "." {
				// fmt.Print(string(letter))
				// fmt.Println(" Is this a number?", unicode.IsNumber(letter))
				if test {
					fmt.Println(letra)
				} else if index == len(line)-1 {
					numberStorage = append(numberStorage, letterIndexStorage)
					break
				}
				if len(letterIndexStorage) == 0 {
					letterIndexStorage = append(letterIndexStorage, index)
				} else if index-1 == letterIndexStorage[len(letterIndexStorage)-1] {
					letterIndexStorage = append(letterIndexStorage, index)
				} else {
					numberStorage = append(numberStorage, letterIndexStorage)
					letterIndexStorage = nil
					letterIndexStorage = append(letterIndexStorage, index)
				}
				// fmt.Printf("This is the index:%v\nThis is the letter:%v\nThis is the line len:%v\n", index, string(letter), len(line))
				// fmt.Println("this is number storage:", numberStorage)
				// fmt.Println("this is indexstorage:", letterIndexStorage)

			}
		}
		LinesNumbers = append(LinesNumbers, numberStorage)
		LinesSymbols = append(LinesSymbols, symbolStorage)

	}
	// for _, linea := range LinesNumbers {
	// 	fmt.Printf("En linea: %v \n", linea[0])
	// 	fmt.Println(linea[1:])
	// }
	fmt.Println(LinesNumbers)
	fmt.Println(LinesSymbols)
}
func test(textarray []string) {
	for _, element := range textarray {
		for _, element2 := range element {
			if element2 != 13 {
				fmt.Print(element2, " ")
			}
		}
	}
}

// this code SHOULD NOT WORK i never check for multiple symbols touching the same number
// but it passes the part1 even when there are symbols touching multiple numbers.
// i dont know if im extremly lucky in that it passes or in that it works like it should.
func part1_2(textarray []string) {
	// this could be one array lower on depth.
	var LinesNumbers [][][]int
	var LinesSymbols [][]int
	var numberstoadd []int
	var result int
	// loop through each line
	for lineIndex, line := range textarray {
		var linesStorage [][]int
		var numberStorage []int
		var symbolStorage []int
		var temp []int
		temp = append(temp, lineIndex)
		// numberStorage = append(numberStorage, lineIndex)
		linesStorage = append(linesStorage, temp)
		// Loop through each character in the line
		for index, letter := range line {
			// Check if the character is a number
			if unicode.IsDigit(letter) {
				// if i try to use not len(numberstorage) in the second check it breaks and i dont even wanna know why.
				if len(numberStorage) == 0 {
					numberStorage = append(numberStorage, index)
				} else if numberStorage[len(numberStorage)-1] != index-1 {
					// If another number is found on the same line it closes the first number and THEN it adds itself to the list
					linesStorage = append(linesStorage, numberStorage)
					numberStorage = nil
					numberStorage = append(numberStorage, index)
				} else {
					numberStorage = append(numberStorage, index)
				}
			} else if letter != '.' && letter != 13 {
				// If it's not a number, it's a symbol. rune 13 is an empty whitespace so empty your
				// keyboard jumps straight trough it when moving the cursor. It was !Fun trying to find it.
				symbolStorage = append(symbolStorage, index)
			}
		}
		linesStorage = append(linesStorage, numberStorage)
		LinesNumbers = append(LinesNumbers, linesStorage)
		LinesSymbols = append(LinesSymbols, symbolStorage)
		// this stores an empty array/slice if the line is empty. i like it cause everything is stored
		// in the same line as it was on the input so i can traverse through it "easier"
	}
	// this is bad code, this amount of for loops is not the correct choice i should have first searched for the symbols
	// and then searched for the numbers that are touching.
	for index, line2 := range LinesSymbols {
		if len(line2) > 0 {
			for _, line := range line2 {
				for lineIndex, num := range LinesNumbers {
					if lineIndex == index+1 || lineIndex == index-1 || lineIndex == index {
						for _, groupofnumbers := range num[1:] {
							var numbertouching bool
							for _, indexofnumber := range groupofnumbers {
								if indexofnumber == line || indexofnumber+1 == line || indexofnumber-1 == line {
									// Now try to do this when theres more than n dimensions.
									numbertouching = true
								}
							}
							if numbertouching {
								var numberinstring string
								var err error
								for _, indexofnumber := range groupofnumbers {
									line := textarray[lineIndex]
									numberinstring += string(line[indexofnumber])
									check(err)

								}
								temp, err := strconv.Atoi(numberinstring)
								check(err)
								numberstoadd = append(numberstoadd, temp)
							}
						}
					}
				}
			}
		}
	}
	//maybe theres a sum function maybe i dont care.
	for _, numbnumber := range numberstoadd {
		result += numbnumber
	}
	println(result)
}

func part2(textarray []string) {
	// this could be one array lower on depth.
	var LinesNumbers [][][]int
	var LinesSymbols [][]int
	var numberstoadd []int
	var result int
	// loop through each line
	for lineIndex, line := range textarray {
		// this stores an empty array/slice if the line is empty. i like it cause everything is stored
		// in the same line as it was on the input so i can traverse through it "easier"
		var linesStorage [][]int
		var numberStorage []int
		var symbolStorage []int
		var temp []int
		temp = append(temp, lineIndex)
		// numberStorage = append(numberStorage, lineIndex)
		linesStorage = append(linesStorage, temp)
		// Loop through each character in the line
		for index, letter := range line {
			// Check if the character is a number
			if unicode.IsDigit(letter) {
				// if i try to use not len(numberstorage) in the second check it breaks and i dont even wanna know why.
				if len(numberStorage) == 0 {
					numberStorage = append(numberStorage, index)
				} else if numberStorage[len(numberStorage)-1] != index-1 {
					// If another number is found on the same line it closes the first number and THEN it adds itself to the list
					linesStorage = append(linesStorage, numberStorage)
					numberStorage = nil
					numberStorage = append(numberStorage, index)
				} else {
					numberStorage = append(numberStorage, index)
				}
			} else if letter != '.' && letter != 13 {
				// If it's not a number, it's a symbol. rune 13 is an empty whitespace so empty your
				// keyboard jumps straight trough it when moving the cursor. It was !Fun trying to find it.
				symbolStorage = append(symbolStorage, index)
			}
		}
		linesStorage = append(linesStorage, numberStorage)
		LinesNumbers = append(LinesNumbers, linesStorage)
		LinesSymbols = append(LinesSymbols, symbolStorage)
	}
	// this is bad code, this amount of for loops is not the correct choice i should have first searched for the symbols
	// and then searched for the numbers that are touching.
	for index, line2 := range LinesSymbols {
		if len(line2) > 0 {
			for _, line := range line2 {
				var numberstouchingsymbols []int
				for lineIndex, num := range LinesNumbers {
					if lineIndex == index+1 || lineIndex == index-1 || lineIndex == index {
						for _, groupofnumbers := range num[1:] {
							var numbertouching bool
							for _, indexofnumber := range groupofnumbers {
								if indexofnumber == line || indexofnumber+1 == line || indexofnumber-1 == line {
									// Now try to do this when theres more than n dimensions.
									numbertouching = true
								}
							}
							if numbertouching {
								var numberinstring string
								var err error
								for _, indexofnumber := range groupofnumbers {
									line := textarray[lineIndex]
									numberinstring += string(line[indexofnumber])
									check(err)

								}
								temp, err := strconv.Atoi(numberinstring)
								check(err)
								numberstouchingsymbols = append(numberstouchingsymbols, temp)
							}
						}
					}
				}
				if len(numberstouchingsymbols) > 1 {
					numberstoadd = append(numberstoadd, (numberstouchingsymbols[0] * numberstouchingsymbols[1]))
				}
			}
		}
		//maybe theres a sum function maybe i dont care.
	}
	for _, numbnumber := range numberstoadd {
		fmt.Println(numbnumber)
		result = result + numbnumber
	}
	fmt.Printf("Result: %v", result)
}
