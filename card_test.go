package cards

import "testing"

func TestInvalidCardsThrowAnError(t *testing.T) {
	c := Card{
		Type: "asdf",
	}
	if c.IsValid() {
		t.Errorf("This card should not be considered valid. Card type: %s", c.Type)
	}

	c.Type = "A"
	if !c.IsValid() {
		t.Errorf("This card should be considered valid. Card type: %s", c.Type)
	}
}
func TestFaceCardsHaveAFixedValueOfTen(t *testing.T) {
	type faceCardsTest struct {
		CardType  string
		CardValue int
	}
	faceCardsTable := []faceCardsTest{
		{"K", 10}, {"Q", 10}, {"J", 10},
	}
	for _, card := range faceCardsTable {
		c := Card{
			Type: card.CardType,
		}
		highValue, _ := c.ValueHigh()
		if highValue != 10 {
			t.Errorf("A king card should have the value of 10. It has the value of %d", highValue)
		}
	}
}
