package main

import (
	"bufio"
	"os"
	"strings"
)

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

func removeSpaces(line string) string {
	line = strings.ReplaceAll(line, " ", "")
	return line
}
