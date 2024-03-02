package day7

import (
	"os"
	"strings"
	"testing"
)

func Testpart1(t *testing.T) {
	dat, err := os.ReadFile("C:/Users/nicod/Documents/GitHub/AdventOfCode/Inputs/input4.txt")
	check(err)
	text := string(dat)
	var textArray = strings.Split(text, "\n")
	got := part1(textArray)
	want := 6440
	if got != want {
		t.Errorf("Got %v But wanted %v", got, want)
	}
}
