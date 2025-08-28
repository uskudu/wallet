package main

import (
	"fmt"
	"wallet/business"
	"wallet/services"
)

func main() {
	i := services.CreateCard("IVANOV IVAN")
	i.Add(100)
	fmt.Println(i)

	fmt.Println("___")
	a := business.New(i.Data.CardholderName, i)
	fmt.Println(a)
	fmt.Println(a.Card)

}
