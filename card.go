package cards

import (
	"errors"
	"strconv"
)

// Card this struct represents a card in the standard American 52 card deck
type Card struct {
	Suit Suit
	Type string
}

var validCardTypes = [...]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

// IsValid return a bool based on the Card.Type
func (c Card) IsValid() bool {
	for _, val := range validCardTypes {
		if c.Type == val {
			return true
		}
	}
	return false
}

// DrawSmall show just the type and suit symbol
func (c Card) DrawSmall() string {
	return c.Type + c.Suit.Symbol
}

// DrawFull an ASCII representation of this card
func (c Card) DrawFull() string {
	return `
/-----\
|  ` + c.Type + `  |
|  ` + c.Suit.Symbol + `  |
\-----/`
}

// ValueHigh return the highest value this card could be worth
// TODO: this should be moved out of the cards package, as these values are
// game specific
func (c Card) ValueHigh() (int, error) {
	if !c.IsValid() {
		return 0, errors.New("Invalid Card")
	}
	values := make(map[string][]int, 52)

	for _, v := range validCardTypes {
		switch v {
		case "K", "Q", "J":
			values[v] = []int{10}
		case "A":
			values["A"] = []int{11, 1}
		default:
			value, _ := strconv.Atoi(v)
			values[v] = []int{value}
		}
	}

	return values[c.Type][0], nil
}
