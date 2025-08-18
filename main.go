package main

import (
	"fmt"
	"wallet/services"
)

func main() {
	my_card := services.CreateCard("JOHN DOE")
	fmt.Println(services.ReadCard(my_card.CardNumber))
	fmt.Println("___")
	fmt.Println(services.DeleteCard(my_card.CardNumber))
	fmt.Println("___")
	fmt.Println(services.ReadCard(my_card.CardNumber))
}
