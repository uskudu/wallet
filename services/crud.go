package services

import (
	"wallet/database"
	card "wallet/wall"
)

func CreateCard(CardholderName string) *card.Card {
	newCard := card.New(CardholderName)
	database.AddCard(newCard)
	return newCard
}

func ReadCard(c *card.Card) string {
	cardNumber := c.Data.CardNumber
	cardInDB := database.DatabaseMock[cardNumber]
	result := cardInDB.Read()
	if cardInDB.Data.CardNumber == "" {
		return "card doesnt exist"
	}
	return result
}

func DeleteCard(c *card.Card) {
	database.RemoveCard(c)
}
