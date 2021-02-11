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
		{list: []string{"D", "D", "D", "D", "D"}, wantIsSameSuit: true},
		{list: []string{"H", "H", "H", "H", "H"}, wantIsSameSuit: true},
		{list: []string{"S", "S", "S", "S", "S"}, wantIsSameSuit: true},
		{list: []string{"C", "C", "C", "C", "C"}, wantIsSameSuit: true},
		{list: []string{"D", "H", "S", "C", "D"}, wantIsSameSuit: false},
		// sameSuitTestData{list: []string{"C", "C", "S", "D", "S"}, wantIsSameSuit: false},
	}
	for _, test := range testData {
		gotIsSameSuit := isSameSuit(test.list)
		if gotIsSameSuit != test.wantIsSameSuit {
			t.Error(fmt.Sprintf("Suit is: (%#v). wantIsSameSuit: %t, gotIsSameSuit: %t", test.list, test.wantIsSameSuit, gotIsSameSuit))
		}
	}
}
