package database

import (
	card "wallet/wall"
)

var DatabaseMock = map[string]*card.Card{}

func AddCard(c *card.Card) bool {
	id := c.Data.CardNumber
	DatabaseMock[id] = c
	return true
}

func RemoveCard(c *card.Card) bool {
	delete(DatabaseMock, c.Data.CardNumber)
	return true
}
