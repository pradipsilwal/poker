package main

import (
	"fmt"
	"testing"
)

type testData struct {
	list           []int64
	wantHigh       int64
	wantIsStraight bool
}

func TestIsStraight(t *testing.T) {
	tests := []testData{
		testData{list: []int64{2, 2, 3, 4, 4}, wantHigh: 4, wantIsStraight: false},
		testData{list: []int64{3, 5, 6, 7, 8}, wantHigh: 8, wantIsStraight: false},
		testData{list: []int64{2, 3, 4, 5, 14}, wantHigh: 14, wantIsStraight: false},
		testData{list: []int64{5, 6, 7, 8, 9}, wantHigh: 9, wantIsStraight: true},
	}

	for _, test := range tests {
		gotIsStraight, gotHigh := isStraight(test.list)
		if gotIsStraight != test.wantIsStraight || gotHigh != test.wantHigh {
			t.Error(errorString(test, gotHigh, gotIsStraight))
		}
	}
}

func errorString(test testData, gotHigh int64, gotIsStraight bool) string {
	return fmt.Sprintf("Hand = (%#v). Wanted High: %d, Got High: %d || Wanted IsStraight: %t, Got IsStraight: %t", test.list, test.wantHigh, gotHigh, test.wantIsStraight, gotIsStraight)
}
