package day7

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	dat, err := os.ReadFile("C:/Users/nicod/Documents/GitHub/AdventOfCode/Inputs/input7.txt")
	check(err)
	text := string(dat)
	// why did it change from simple \n to \r\n? i want to believe the creator of adventofcode was testing us.
	// for my personal metrics i dont consider this a read trough of the input as neither did i consider the original
	// simple split a read trough.
	var textArray = strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	got := part1(textArray)
	want := int64(251121738)
	if got != want {
		t.Errorf("Wanted %v but Got %v", want, got)
	} else {
		fmt.Println(got)
	}
}

//{2638T 17 highCard} 0{2Q7K6 806 highCard} 1{2AJT7 631 highCard} 2{38A27 303 highCard} 3{3QT25 470 highCard} 4{576J4 826 highCard} 5{5893T 841 highCard} 6{5Q397 381 highCard} 7{65732 67 highCard} 8{827K6 552 highCard} 9{834JQ 962 highCard} 10{8T954 579 highCard} 11{8TQA2 151 highCard} 12{94673 646 highCard} 13{9T5A4 738 highCard} 14{9QA53 597 highCard} 15{J37QA 670 highCard} 16{Q684T 820 highCard} 17{QKJA4 276 highCard} 18{K3795 875 highCard} 19{A3T76 71 highCard} 20{A924K 555 highCard} 21{A9TK7 710 highCard} 22{2QQ5J 32 pair} 23{3T3K9 718 pair} 24{56674 627 pair} 25{5T598 858 pair} 26{6K6J2 351 pair} 27{72889 264 pair} 28{79TK9 142 pair} 29{7Q922 53 pair} 30{7K663 291 pair} 31{884JQ 386 pair} 32{95567 364 pair} 33{9A2T2 334 pair} 34{T7627 879 pair} 35{TJ223 828 pair} 36{J79A7 865 pair} 37{JTJ63 457 pair} 38{QJK6K 964 pair} 39{K7KT6 677 pair} 40{A3AQ9 796 pair} 41{AT8JA 609 pair} 42{27233 359 twoPair} 43{323QQ 568 twoPair} 44{3553K 171 twoPair} 45{5335J 644 twoPair} 46{55TT4 895 twoPair} 47{74979 127 twoPair} 48{77883 103 twoPair} 49{7KAK7 63 twoPair} 50{8KT8K 463 twoPair} 51{92T2T 893 twoPair} 52{933TT 851 twoPair} 53{9889J 846 twoPair} 54{T2T2J 5 twoPair} 55{T3KT3 254 twoPair} 56{TT3J3 693 twoPair} 57{JTATA 186 twoPair} 58{JQ33Q 213 twoPair} 59{K2K66 909 twoPair} 60{AJ99A 312 twoPair} 61{22552 201 fullHouse} 62{53553 542 fullHouse} 63{55656 482 fullHouse} 64{78787 764 fullHouse} 65{7QQ77 426 fullHouse} 66{82228 641 fullHouse} 67{88899 258 fullHouse} 68{J888J 808 fullHouse} 69{K3K33 367 fullHouse} 70{333A3 410 poker} 71{3QQQ3 207 poker} 72{47747 538 poker} 73{66565 890 poker} 74{6666J 914 poker} 75{78887 999 poker} 76{85888 241 poker} 77{95999 475 poker} 78{TAAAA 152 poker} 79{JQJQQ 873 poker} 80{33365 495 treeOfAKind} 81{44764 115 treeOfAKind} 82{45A55 906 treeOfAKind} 83{46A66 407 treeOfAKind} 84{4KA44 244 treeOfAKind} 85{52355 782 treeOfAKind} 86{66678 582 treeOfAKind} 87{666Q3 118 treeOfAKind} 88{8JKKK 78 treeOfAKind} 89{8A8J8 611 treeOfAKind} 90{T666J 983 treeOfAKind} 91{TT3T9 509 treeOfAKind} 92{J3339 607 treeOfAKind} 93{J4434 784 treeOfAKind} 94{J99K9 427 treeOfAKind} 95{J9A99 986 treeOfAKind} 96{K3335 221 treeOfAKind} 97{KTT5T 116 treeOfAKind} 98{AA4AJ 345 treeOfAKind} 99
//{2638T 17 highCard} 0{2Q7K6 806 highCard} 1{2AJT7 631 highCard} 2{38A27 303 highCard} 3{3QT25 470 highCard} 4{576J4 826 highCard} 5{5893T 841 highCard} 6{5Q397 381 highCard} 7{65732 67 highCard} 8{827K6 552 highCard} 9{834JQ 962 highCard} 10{8T954 579 highCard} 11{8TQA2 151 highCard} 12{94673 646 highCard} 13{9T5A4 738 highCard} 14{9QA53 597 highCard} 15{J37QA 670 highCard} 16{Q684T 820 highCard} 17{QKJA4 276 highCard} 18{K3795 875 highCard} 19{A3T76 71 highCard} 20{A924K 555 highCard} 21{A9TK7 710 highCard} 22{2QQ5J 32 pair} 23{3T3K9 718 pair} 24{56674 627 pair} 25{5T598 858 pair} 26{6K6J2 351 pair} 27{72889 264 pair} 28{79TK9 142 pair} 29{7Q922 53 pair} 30{7K663 291 pair} 31{884JQ 386 pair} 32{95567 364 pair} 33{9A2T2 334 pair} 34{T7627 879 pair} 35{TJ223 828 pair} 36{J79A7 865 pair} 37{JTJ63 457 pair} 38{QJK6K 964 pair} 39{K7KT6 677 pair} 40{A3AQ9 796 pair} 41{AT8JA 609 pair} 42{27233 359 twoPair} 43{323QQ 568 twoPair} 44{3553K 171 twoPair} 45{5335J 644 twoPair} 46{55TT4 895 twoPair} 47{74979 127 twoPair} 48{77883 103 twoPair} 49{7KAK7 63 twoPair} 50{8KT8K 463 twoPair} 51{92T2T 893 twoPair} 52{933TT 851 twoPair} 53{9889J 846 twoPair} 54{T2T2J 5 twoPair} 55{T3KT3 254 twoPair} 56{TT3J3 693 twoPair} 57{JTATA 186 twoPair} 58{JQ33Q 213 twoPair} 59{K2K66 909 twoPair} 60{AJ99A 312 twoPair} 61{22552 201 fullHouse} 62{55656 482 fullHouse} 63{78787 764 fullHouse} 64{7QQ77 426 fullHouse} 65{82228 641 fullHouse} 66{J888J 808 fullHouse} 67{K3K33 367 fullHouse} 68{333A3 410 poker} 69{3QQQ3 207 poker} 70{47747 538 poker} 71{53553 542 poker} 72{66565 890 poker} 73{6666J 914 poker} 74{78887 999 poker} 75{85888 241 poker} 76{88899 258 poker} 77{95999 475 poker} 78{TAAAA 152 poker} 79{JQJQQ 873 poker} 80{33365 495 treeOfAKind} 81{44764 115 treeOfAKind} 82{45A55 906 treeOfAKind} 83{46A66 407 treeOfAKind} 84{4KA44 244 treeOfAKind} 85{52355 782 treeOfAKind} 86{66678 582 treeOfAKind} 87{666Q3 118 treeOfAKind} 88{8JKKK 78 treeOfAKind} 89{8A8J8 611 treeOfAKind} 90{T666J 983 treeOfAKind} 91{TT3T9 509 treeOfAKind} 92{J3339 607 treeOfAKind} 93{J4434 784 treeOfAKind} 94{J99K9 427 treeOfAKind} 95{J9A99 986 treeOfAKind} 96{K3335 221 treeOfAKind} 97{KTT5T 116 treeOfAKind} 98{AA4AJ 345 treeOfAKind} 9
