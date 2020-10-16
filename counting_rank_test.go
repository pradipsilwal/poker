package main

import (
	"fmt"
	"reflect"
	"testing"
)

type countRankData struct {
	ranks        []int64
	wantCountMap map[int64]int64
}

func TestCountingRank(t *testing.T) {
	testData := []countRankData{
		countRankData{ranks: []int64{2, 3, 4, 5, 7}, wantCountMap: map[int64]int64{2: 1, 3: 1, 4: 1, 5: 1, 7: 1}},
		countRankData{ranks: []int64{2, 2, 2, 5, 6}, wantCountMap: map[int64]int64{2: 3, 5: 1, 6: 1}},
		countRankData{ranks: []int64{2, 2, 4, 5, 5}, wantCountMap: map[int64]int64{2: 2, 4: 1, 5: 2}},
		countRankData{ranks: []int64{3, 3, 3, 6, 6}, wantCountMap: map[int64]int64{3: 3, 6: 2}},
		countRankData{ranks: []int64{3, 3, 5, 5, 6}, wantCountMap: map[int64]int64{3: 2, 5: 2, 6: 1}},
		countRankData{ranks: []int64{5, 7, 8, 10, 12}, wantCountMap: map[int64]int64{5: 1, 7: 1, 8: 1, 10: 1, 12: 1}},
	}

	for _, test := range testData {
		gotCountMap := countingSameRanks(test.ranks)
		if !reflect.DeepEqual(gotCountMap, test.wantCountMap) {
			t.Error(fmt.Sprintf("Ranks: %#v, WantCountMap: %#v, GotCountMap: %#v", test.ranks, test.wantCountMap, gotCountMap))
		}
	}
}
