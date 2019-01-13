package cards

import (
	"math/rand"
	"time"
)

// Deck this struct represents a collection of Cards so that we can play a game
type Deck struct {
	Cards []Card
}

// NewDeck return a deck of cards
func NewDeck() Deck {
	c := []Card{}
	for _, s := range validCardsuits {
		// suit := Suit{s[0], s[1]}
		for _, v := range validCardTypes {
			c = append(c, Card{Suit{s[0], s[1]}, v})
		}
	}
	d := Deck{c}

	return d
}

// Shuffle return shuffled deck
func (d Deck) Shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d.Cards {
		rand := r.Intn(len(d.Cards) - 1)
		d.Cards[i], d.Cards[rand] = d.Cards[rand], d.Cards[i]
	}
}
