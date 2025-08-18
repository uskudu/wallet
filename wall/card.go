package card

import "strconv"

type CardData struct {
	CardholderName string
	CardNumber     int64
	ExpirationDate string // using time.Format(DDMMYYYY), which is string
	CvcCode        int
}

// implement card other data
type Card struct {
	Data    CardData
	Balance int
}

func (c Card) Read() string {
	return "Cardholder: " + c.Data.CardholderName + "\nCardNumber: " + strconv.Itoa(int(c.Data.CardNumber)) + "\nExpirationDate: " + c.Data.ExpirationDate + "\nCvcCode: " + strconv.Itoa(c.Data.CvcCode) + "\nBalance: " + strconv.Itoa(c.Balance)
}
