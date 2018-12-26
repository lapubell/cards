package cards

import "errors"

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
	values["A"] = []int{1, 11}
	values["K"] = []int{10}
	values["Q"] = []int{10}
	values["J"] = []int{10}
	values["1"] = []int{1}
	values["2"] = []int{2}
	values["3"] = []int{3}
	values["4"] = []int{4}
	values["5"] = []int{5}
	values["6"] = []int{6}
	values["7"] = []int{7}
	values["8"] = []int{8}
	values["9"] = []int{9}
	values["10"] = []int{10}

	return values[c.Type][0], nil
}
