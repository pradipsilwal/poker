package main

import (
	"fmt"
	"testing"
)

type straightNSameSuits struct {
	listRanks              []int64
	listSuits              []string
	wantStraightNSameSuits bool
	wantHighRank           int64
}

func TestIsStraightNSameSuit(t *testing.T) {
	testData := []straightNSameSuits{
		straightNSameSuits{listRanks: []int64{10, 11, 12, 13, 14}, listSuits: []string{"D", "H", "D", "S", "D"},
			wantStraightNSameSuits: false, wantHighRank: 14},
		straightNSameSuits{listRanks: []int64{10, 11, 12, 13, 14}, listSuits: []string{"D", "D", "D", "S", "D"},
			wantStraightNSameSuits: false, wantHighRank: 14},
		straightNSameSuits{listRanks: []int64{10, 11, 12, 13, 14}, listSuits: []string{"H", "H", "D", "H", "D"},
			wantStraightNSameSuits: false, wantHighRank: 14},
		straightNSameSuits{listRanks: []int64{10, 11, 12, 13, 14}, listSuits: []string{"D", "D", "D", "D", "D"},
			wantStraightNSameSuits: true, wantHighRank: 14},
		straightNSameSuits{listRanks: []int64{9, 10, 11, 12, 13}, listSuits: []string{"D", "D", "D", "D", "D"},
			wantStraightNSameSuits: true, wantHighRank: 13},
		straightNSameSuits{listRanks: []int64{10, 11, 11, 13, 14}, listSuits: []string{"D", "D", "H", "D", "D"},
			wantStraightNSameSuits: false, wantHighRank: 14},
		straightNSameSuits{listRanks: []int64{5, 10, 11, 12, 13}, listSuits: []string{"D", "D", "D", "D", "D"},
			wantStraightNSameSuits: false, wantHighRank: 13},
	}
	for _, test := range testData {
		gotStraightNSameSuits, gotHighRank := isStraightNSameSuit(test.listRanks, test.listSuits)
		if gotStraightNSameSuits != test.wantStraightNSameSuits || gotHighRank != test.wantHighRank {
			t.Error(fmt.Sprintf("Ranks: %#v, Suits: %#v, wantStraightNSameSuits: %t, gotStraightNSameSuits: %t | WantHighRank: %d: GotHighRank: %d",
				test.listRanks, test.listSuits, test.wantStraightNSameSuits, gotStraightNSameSuits, test.wantHighRank, gotHighRank))
		}
	}
}
