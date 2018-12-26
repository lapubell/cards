package cards

// Deck this struct represents a collection of Cards so that we can play a game
type Deck struct {
	Cards []Card
}

// GetDeck return a deck of cards
func GetDeck() []Card {
	d := []Card{}
	for _, s := range suits {
		for _, v := range validTypes {
			d = append(d, Card{s, v})
		}
	}
	return d
}
