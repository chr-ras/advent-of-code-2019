package cards

import (
	"reflect"
	"testing"
)

func TestShuffleDeck(t *testing.T) {
	cases := []struct {
		shuffleProcessInput []string
		deckSizeInput       int
		expectedDeck        []int
	}{
		{
			[]string{"deal into new stack"},
			10,
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			[]string{"cut 3"},
			10,
			[]int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
		},
		{
			[]string{"cut -4"},
			10,
			[]int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5},
		},
		{
			[]string{"deal with increment 3"},
			10,
			[]int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3},
		},
		{
			[]string{"deal with increment 7", "deal into new stack", "deal into new stack"},
			10,
			[]int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		},
		{
			[]string{"cut 6", "deal with increment 7", "deal into new stack"},
			10,
			[]int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6},
		},
		{
			[]string{"deal with increment 7", "deal with increment 9", "cut -2"},
			10,
			[]int{6, 3, 0, 7, 4, 1, 8, 5, 2, 9},
		},
		{
			[]string{"deal into new stack", "cut -2", "deal with increment 7", "cut 8", "cut -4", "deal with increment 7", "cut 3", "deal with increment 9", "deal with increment 3", "cut -1"},
			10,
			[]int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6},
		},
	}

	for _, c := range cases {
		actualDeck := ShuffleDeck(c.shuffleProcessInput, c.deckSizeInput)

		if !reflect.DeepEqual(actualDeck, c.expectedDeck) {
			t.Errorf("ShuffleDeck(%v, %d) == %v, expected %v", c.shuffleProcessInput, c.deckSizeInput, actualDeck, c.expectedDeck)
		}
	}
}
