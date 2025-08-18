package services

import (
	"math/rand"
	"time"
	"wallet/database"
	card "wallet/wall"
)

const (
	MMYYYY = "01/2006"
)

func cardNumGenerator() int64 {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	min := int64(1_0000_0000_0000_000)  // 10^15
	max := int64(9_9999_9999_9999_9999) // 10^16 - 1

	return r.Int63n(max-min+1) + min
}

func cvcGenerator() int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	min := int(100)
	max := int(999)

	return r.Intn(max-min+1) + min
}

func CreateCard(CardholderName string) card.Card {
	dateInTenYears := time.Now().Add(time.Hour * 24 * 365 * 10)
	dateExp := dateInTenYears.Format(MMYYYY)

	newCardData := card.CardData{
		CardholderName,
		cardNumGenerator(),
		dateExp,
		cvcGenerator(),
	}
	newCard := card.Card{Data: newCardData, Balance: 0}
	database.AddCard(newCard)
	return newCard
}

func ReadCard(cardNumber int64) string {
	cardFromDb := database.DatabaseMock[cardNumber]
	result := cardFromDb.Read()
	if cardFromDb.Data.CardNumber == 0 {
		return "card with that number doesnt appear in database"
	}
	return result
}

func DeleteCard(cardNumber int64) string {
	deleted := database.RemoveCard(cardNumber)
	if deleted {
		return "card deleted"
	}
	return "could not delete card"
}
