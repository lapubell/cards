package cards

import (
	"errors"
	"strconv"
)

// Card this struct represents a card in the standard American 52 card deck
type Card struct {
	Suit string
	Type string
}

var validTypes = [...]string{"A", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

// IsValid return a bool based on the Card.Type
func (c Card) IsValid() bool {
	for _, val := range validTypes {
		if c.Type == val {
			return true
		}
	}
	return false
}

// ValueHigh return the highest value this card could be worth
func (c Card) ValueHigh() (int, error) {
	if !c.IsValid() {
		return 0, errors.New("Invalid Card")
	}
	values := make(map[string][]int, 52)

	for _, v := range validTypes {
		switch v {
		case "K", "Q", "J", "10":
			values[v] = []int{10}
		case "A":
			values["A"] = []int{1, 11}
		default:
			value, _ := strconv.Atoi(v)
			values[v] = []int{value}
		}
	}

	return values[c.Type][0], nil
}
