package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func removeSpaces(line string) string {
	line = strings.ReplaceAll(line, " ", "")
	return line
}

// Scans the file for any hands and returns array of hands
// One line contains hand for player 1 and player 2 for one game
func getHandFromFile(filename string) ([]string, error) {
	var hands []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hands = append(hands, removeSpaces(scanner.Text()))
	}

	if scanner.Err() != nil {
		return nil, err
	}
	return hands, nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// add each player's hand in a array for each game
func splitPlayersHand(hands []string) [][]string {
	var allHands [][]string
	for _, hand := range hands {
		row := []string{hand[:10], hand[10:]}
		allHands = append(allHands, row)
	}
	return allHands
}

// Split ranks and suits into separate slices
func splitRanksAndSuits(singleHand string) ([]int64, []string) {
	var ranks []int64
	var suits []string
	for i := 0; i < len(singleHand); i += 2 {
		rankVal, err := changeRankToInt(string(singleHand[i]))
		checkError(err)
		ranks = append(ranks, rankVal)
		suits = append(suits, string(singleHand[i+1]))
	}
	return ranks, suits
}

// Assign integer value to the hand for easy comparision
func changeRankToInt(stringRank string) (int64, error) {
	var intRank int64
	if stringRank == "J" {
		intRank = 11
	} else if stringRank == "Q" {
		intRank = 12
	} else if stringRank == "K" {
		intRank = 13
	} else if stringRank == "A" {
		intRank = 14
	} else {
		intRank, err := strconv.ParseInt(stringRank, 10, 64)
		return intRank, err
	}
	return intRank, nil
}

// Checks if the hand is straight
func isStraight(ranks []int64) (bool, int64) {
	for i := 0; i < len(ranks)-1; i++ {
		if (ranks[i+1] - ranks[i]) != 1 {
			return false, ranks[len(ranks)-1]
		}
	}
	return true, ranks[len(ranks)-1]
}

//Checks if all the suits in a hand are same
func isSameSuit(suits []string) bool {
	for i := 0; i < len(suits)-1; i++ {
		if suits[i] != suits[i+1] {
			return false
		}
	}
	return true
}

// Check Royal Flush
func isRoyalFlush(ranks []int64, suits []string) bool {
	isStraightNSameSuit, highRank := isStraightNSameSuit(ranks, suits)
	if isStraightNSameSuit && highRank == 14 {
		return true
	}
	return false
}

// Check Straight Flush
func isStraightFlush(ranks []int64, suits []string) (bool, int64) {
	isStraightNSameSuit, highRank := isStraightNSameSuit(ranks, suits)
	if isStraightNSameSuit && highRank != 14 {
		return true, highRank
	}
	return false, highRank
}

// Check if it is straight and flush
func isStraightNSameSuit(ranks []int64, suits []string) (bool, int64) {
	isSameSuit := isSameSuit(suits)
	isStraight, highRank := isStraight(ranks)
	if isSameSuit && isStraight {
		return true, highRank
	}
	return false, highRank
}

// Counting the number of pairs
func countingSameRanks(ranks []int64) map[int64]int64 {
	pairMap := make(map[int64]int64)
	for _, rank := range ranks {
		pairMap[rank]++
	}
	return pairMap
}

func main() {
	hands, err := getHandFromFile("hand.txt")
	checkError(err)
	allHands := splitPlayersHand(hands)
	ranks, suits := splitRanksAndSuits(allHands[0][0])
	fmt.Println(ranks)
	fmt.Println(suits)
}
