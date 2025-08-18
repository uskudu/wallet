package main

import (
	"fmt"
	"wallet/services"
)

func main() {
	i := services.CreateCard("IVANOV IVAN")
	j := services.CreateCard("JOHN DOE")
	//services.DeleteCard(i)
	i.Add(10)
	err := services.Transaction(i, j, 5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(i)
	fmt.Println("___")
	fmt.Println(j)
	fmt.Println("___")
}
