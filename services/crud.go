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

func DeleteCard(c *card.Card) {
	database.RemoveCard(c)
}
