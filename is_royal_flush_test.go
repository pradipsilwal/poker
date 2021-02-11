package main

import (
	"fmt"
	"testing"
)

type royalFlushData struct {
	listRanks   []int64
	listSuits   []string
	wantIsFlush bool
}

func TestIsRoyalFlush(t *testing.T) {
	testData := []royalFlushData{
		{listRanks: []int64{10, 11, 12, 13, 14}, listSuits: []string{"D", "H", "D", "S", "D"},
			wantIsFlush: false},
		{listRanks: []int64{10, 11, 12, 13, 14}, listSuits: []string{"D", "D", "D", "S", "D"},
			wantIsFlush: false},
		{listRanks: []int64{10, 11, 12, 13, 14}, listSuits: []string{"H", "H", "D", "H", "D"},
			wantIsFlush: false},
		{listRanks: []int64{10, 11, 12, 13, 14}, listSuits: []string{"D", "D", "D", "D", "D"},
			wantIsFlush: true},
		{listRanks: []int64{9, 10, 11, 12, 13}, listSuits: []string{"D", "D", "D", "D", "D"},
			wantIsFlush: false},
		{listRanks: []int64{10, 11, 11, 13, 14}, listSuits: []string{"D", "D", "H", "D", "D"},
			wantIsFlush: false},
		{listRanks: []int64{5, 10, 11, 12, 13}, listSuits: []string{"D", "D", "D", "D", "D"},
			wantIsFlush: false},
		// royalFlushData{listRanks: []int64{5, 10, 11, 12, 13}, listSuits: []string{"D", "D", "D", "D", "D"},
		// wantIsFlush: false},
	}
	for _, test := range testData {
		gotIsFlush := isRoyalFlush(test.listRanks, test.listSuits)
		if gotIsFlush != test.wantIsFlush {
			t.Error(fmt.Sprintf("Ranks: %#v, Suits: %#v, WantIsFlush: %t, GotIsFlush: %t",
				test.listRanks, test.listSuits, test.wantIsFlush, gotIsFlush))
		}
	}
}
