package bet

func (a *Account) AddFunds(amount int) error {
	err := a.Card.Sub(amount)
	if err != nil {
		return err
	}
	a.add(amount)
	return nil
}

func (a *Account) Withdraw(amount int) error {
	err := a.sub(amount)
	if err != nil {
		return err
	}
	a.Card.Add(amount)
	return nil
}
