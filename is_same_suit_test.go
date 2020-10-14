package main

import (
	"fmt"
	"testing"
)

type sameSuitTestData struct {
	list           []string
	wantIsSameSuit bool
}

func TestIsSameSuit(t *testing.T) {
	testData := []sameSuitTestData{
		sameSuitTestData{list: []string{"D", "D", "D", "D", "D"}, wantIsSameSuit: true},
		sameSuitTestData{list: []string{"H", "H", "H", "H", "H"}, wantIsSameSuit: true},
		sameSuitTestData{list: []string{"S", "S", "S", "S", "S"}, wantIsSameSuit: true},
		sameSuitTestData{list: []string{"C", "C", "C", "C", "C"}, wantIsSameSuit: true},
		sameSuitTestData{list: []string{"D", "H", "S", "C", "D"}, wantIsSameSuit: false},
		sameSuitTestData{list: []string{"C", "C", "S", "D", "S"}, wantIsSameSuit: false},
	}
	for _, test := range testData {
		gotIsSameSuit := isSameSuit(test.list)
		if gotIsSameSuit != test.wantIsSameSuit {
			t.Error(fmt.Sprintf("Suit is: (%#v). wantIsSameSuit: %t, gotIsSameSuit: %t", test.list, test.wantIsSameSuit, gotIsSameSuit))
		}
	}
}
