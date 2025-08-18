package services

import (
	"fmt"
	"wallet/database"
	card "wallet/wall"
)

func Transaction(senderCard *card.Card, receiverCard *card.Card, amount int) error {
	senderInDB, senderExists := database.DatabaseMock[senderCard.Data.CardNumber]
	receiverInDB, receiverExists := database.DatabaseMock[receiverCard.Data.CardNumber]

	if !senderExists || !receiverExists {
		return fmt.Errorf("transaction failed: one or both cards not found")
	}

	if err := senderInDB.Sub(amount); err != nil {
		return err
	}
	receiverInDB.Add(amount)
	return nil
}
