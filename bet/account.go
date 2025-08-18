package bet

import (
	"fmt"
	"strconv"
	card "wallet/wall"
)

type Account struct {
	Name       string
	Card       card.Card
	Balance    int
	BetHistory string
}

func New(Name string, cardInfo *card.Card) *Account {
	newAccount := Account{Name, *cardInfo, 0, ""}
	return &newAccount
}

func (a *Account) String() string {
	return "Name: " + a.Name +
		"\nBalance: " + strconv.Itoa(a.Balance) +
		"\nBet history: " + a.BetHistory
}

func (a *Account) add(amount int) {
	a.Balance += amount
}

func (a *Account) sub(amount int) error {
	if amount > a.Balance {
		return fmt.Errorf("insufficient funds")
	}
	a.Balance -= amount
	return nil
}
