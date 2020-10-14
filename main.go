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

func splitPlayersHand(hands []string) [][]string {
	var allHands [][]string
	for _, hand := range hands {
		row := []string{hand[:10], hand[10:]}
		allHands = append(allHands, row)
	}
	return allHands
}

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

func isStraight(ranks []int64) (bool, int64) {
	for i := 0; i < len(ranks)-1; i++ {
		if (ranks[i+1] - ranks[i]) != 1 {
			return false, ranks[len(ranks)-1]
		}
	}
	return true, ranks[len(ranks)-1]
}

func isSameSuit(suits []string) bool {
	for i := 0; i < len(suits)-1; i++ {
		if suits[i] != suits[i+1] {
			return false
		}
	}
	return true
}

func main() {
	hands, err := getHandFromFile("hand.txt")
	checkError(err)
	allHands := splitPlayersHand(hands)
	ranks, suits := splitRanksAndSuits(allHands[0][0])
	fmt.Println(ranks)
	fmt.Println(suits)
}
