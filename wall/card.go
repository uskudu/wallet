package card

import (
	"fmt"
	"strconv"
)

type CardData struct {
	CardholderName string
	CardNumber     string
	ExpirationDate string // using time.Format(MM/YY), which is string
	CvcCode        string
}

type Card struct {
	Data    CardData
	Balance int
}

func New(CardholderName string) *Card {
	newCardData := CardData{
		CardholderName,
		cardNumGenerator(),
		expDateGenerator(),
		cvcGenerator(),
	}
	newCard := &Card{Data: newCardData, Balance: 0}
	return newCard
}

func (c *Card) String() string {
	return "Cardholder: " + c.Data.CardholderName +
		"\nCardNumber: " + c.Data.CardNumber +
		"\nExpirationDate: " + c.Data.ExpirationDate +
		"\nCvcCode: " + c.Data.CvcCode +
		"\nBalance: " + strconv.Itoa(c.Balance)
}

func (c *Card) Add(amount int) {
	c.Balance += amount
}

func (c *Card) Sub(amount int) error {
	if amount > c.Balance {
		return fmt.Errorf("insufficient funds")
	}
	c.Balance -= amount
	return nil
}
