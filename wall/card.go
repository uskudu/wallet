package card

import (
	"strconv"
)

type CardData struct {
	CardholderName string
	CardNumber     int64
	ExpirationDate string // using time.Format(DDMMYYYY), which is string
	CvcCode        int
}

func (c CardData) Read() string {
	return "Cardholder: " + c.CardholderName + "\nCardNumber: " + strconv.Itoa(int(c.CardNumber)) + "\nExpirationDate: " + c.ExpirationDate + "\nCvcCode: " + strconv.Itoa(c.CvcCode)
}

// implement card other data
type Card struct {
	data    CardData
	balance int
}
