package database

import (
	card "wallet/wall"
)

var DatabaseMock = map[int64]card.CardData{}

func AddCard(c card.CardData) bool {
	id := c.CardNumber
	DatabaseMock[id] = c
	return true
}

func RemoveCard(cardNumber int64) bool {
	delete(DatabaseMock, cardNumber)
	return true
}
