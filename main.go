package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Card value of the hand
// Hand with the lowest number is the strongest hand
const (
	ROYALFLUSH    = 1
	STRAIGHTFLUSH = 2
	FOUROFAKIND   = 3
	FULLHOUSE     = 4
	FLUSH         = 5
	STRAIGHT      = 6
	THREEOFAKIND  = 7
	TWOPAIRS      = 8
	PAIRS         = 9
	HIGHCARD      = 10
)

type handStrength struct {
	strength        int64
	highRank        int64
	onePairHighRank int64
	twoPairHighRank int64
}

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
func splitPlayersHand(allHands []string) [][]string {
	var singleHand [][]string
	for _, hand := range allHands {
		row := []string{hand[:10], hand[10:]}
		singleHand = append(singleHand, row)
	}
	return singleHand
}

// Split ranks and suits into separate slices
func splitRanksAndSuits(singlePlayerHand string) ([]int64, []string) {
	var ranks []int64
	var suits []string
	for i := 0; i < len(singlePlayerHand); i += 2 {
		rankVal, err := changeRankToInt(string(singlePlayerHand[i]))
		checkError(err)
		ranks = append(ranks, rankVal)
		suits = append(suits, string(singlePlayerHand[i+1]))
	}
	return ranks, suits
}

// Assign integer value to the hand for easy comparision
func changeRankToInt(stringRank string) (int64, error) {
	var intRank int64
	if stringRank == "T" {
		intRank = 10
	} else if stringRank == "J" {
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

func calculateWinner(singleHand []string) {
	var varHandStrength []handStrength
	for _, hand := range singleHand {
		var singleHandStrength handStrength
		ranks, suits := splitRanksAndSuits(hand)
		fmt.Println(ranks, suits)
		if isRoyalFlush(ranks, suits) {
			singleHandStrength.strength = ROYALFLUSH
		} else if status, highRank := isStraightFlush(ranks, suits); status {
			singleHandStrength.strength = STRAIGHTFLUSH
			singleHandStrength.highRank = highRank
		} else if isSameSuit(suits) {
			singleHandStrength.strength = FLUSH
		} else if status, highRank := isStraight(ranks); status {
			singleHandStrength.strength = STRAIGHT
			singleHandStrength.highRank = highRank
		} else {
			pairMap := countingSameRanks(ranks)
			if len(pairMap) == 2 {
				k := getKey(pairMap, 4)[0]
				if k != 0 {
					singleHandStrength.strength = FOUROFAKIND
					singleHandStrength.onePairHighRank = k
				} else {
					singleHandStrength.strength = FULLHOUSE
					singleHandStrength.twoPairHighRank = getKey(pairMap, 3)[0]
				}
			} else if len(pairMap) == 3 {
				k := getKey(pairMap, 3)[0]
				if k != 0 {
					singleHandStrength.strength = THREEOFAKIND
					singleHandStrength.onePairHighRank = k
				} else {
					singleHandStrength.strength = TWOPAIRS
					key := getKey(pairMap, 2)
					singleHandStrength.onePairHighRank = key[0]
					singleHandStrength.twoPairHighRank = key[1]
					singleHandStrength.highRank = getKey(pairMap, 1)[0]
				}
			} else if len(pairMap) == 4 {
				singleHandStrength.strength = PAIRS
				singleHandStrength.onePairHighRank = getKey(pairMap, 2)[0]
			} else {
				singleHandStrength.strength = HIGHCARD
			}
		}
		varHandStrength = append(varHandStrength, singleHandStrength)
	}
	fmt.Println(varHandStrength)
}

func getKey(myMap map[int64]int64, value int64) []int64 {
	var key []int64
	for k, v := range myMap {
		if value == v {
			key = append(key, k)
		}
	}
	return key
}

func main() {
	allHands, err := getHandFromFile("hand.txt")
	checkError(err)
	singleHand := splitPlayersHand(allHands)
	for _, hand := range singleHand {
		calculateWinner(hand)
	}
}
