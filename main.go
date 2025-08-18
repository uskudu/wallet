package main

import (
	"fmt"
	"wallet/services"
)

func main() {
	services.CreateCard("JOHN DOE")
	q := services.CreateCard("IVANOV IVAN")
	services.CreateCard("MAGOMEDOV MAGOMED")
	fmt.Println(services.ReadCard(q.Data.CardNumber))
}
