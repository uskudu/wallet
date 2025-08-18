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

// implement card other data
type Card struct {
	Data    CardData
	Balance int
}

func (c *Card) Read() string {
	return "Cardholder: " + c.Data.CardholderName + "\nCardNumber: " + c.Data.CardNumber + "\nExpirationDate: " + c.Data.ExpirationDate + "\nCvcCode: " + c.Data.CvcCode + "\nBalance: " + strconv.Itoa(c.Balance)
}

func (c *Card) Sub(amount int) error {
	if amount > c.Balance {
		return fmt.Errorf("insufficient funds")
	}
	c.Balance -= amount
	return nil
}

func (c *Card) Add(amount int) {
	c.Balance += amount
}
