package bet

import card "wallet/wall"

type Account struct {
	Name          string
	Card          card.Card
	Balance       int
	BetStatistics []string
}

func New(Name string, cardInfo card.Card) *Account {
	newAccount := Account{Name, cardInfo, 0, []string{}}
	return &newAccount
}
