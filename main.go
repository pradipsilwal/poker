package main

import (
	"fmt"
	"log"
)

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

func main() {
	hands, err := getHandFromFile("hand.txt")
	checkError(err)
	allHands := splitPlayersHand(hands)
	for _, hand := range allHands {
		fmt.Println(hand)
	}
}
