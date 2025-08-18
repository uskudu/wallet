package database

import (
	card "wallet/wall"
)

var DatabaseMock = map[int64]card.Card{}

func AddCard(c card.Card) bool {
	id := c.Data.CardNumber
	DatabaseMock[id] = c
	return true
}

func RemoveCard(cardNumber int64) bool {
	delete(DatabaseMock, cardNumber)
	return true
}
