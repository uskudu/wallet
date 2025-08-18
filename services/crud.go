package services

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"wallet/database"
	card "wallet/wall"
)

const (
	MMYY = "01/06"
)

func cardNumGenerator() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder
	for i := 0; i < 16; i++ {
		sb.WriteByte(byte('0' + r.Intn(10)))
	}
	return sb.String()
}

func cvcGenerator() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%03d", r.Intn(1000)) // 000 to 999
}

func expDateGenerator() string {
	dateInTenYears := time.Now().Add(time.Hour * 24 * 365 * 10)
	dateExp := dateInTenYears.Format(MMYY)
	return dateExp
}

func CreateCard(CardholderName string) *card.Card {

	newCardData := card.CardData{
		CardholderName,
		cardNumGenerator(),
		expDateGenerator(),
		cvcGenerator(),
	}
	newCard := &card.Card{Data: newCardData, Balance: 0}
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
