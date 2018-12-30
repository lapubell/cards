package cards

// Suit this struct represents a suit that a card holds
type Suit struct {
	Name   string
	Symbol string
}

var validCardsuits = [...][]string{
	[]string{"Spades", "\u2660"},
	[]string{"Hearts", "\u2665"},
	[]string{"Clubs", "\u2663"},
	[]string{"Diamond", "\u2666"},
}
