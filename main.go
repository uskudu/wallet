package main

import (
	"wallet/business"
	"wallet/services"
)

func main() {
	i := services.CreateCard("IVANOV IVAN")
	//i.Add(100)
	//fmt.Println(i)
	//
	//fmt.Println("___")
	//a := business.New(i.Data.CardholderName, i)
	//fmt.Println(a)
	//fmt.Println(a.Card)
	var _ = i
	h1 := business.Horse{Number: "077", Name: "alfa", OwnerName: "IVANOV IVAN", RaceHistory: []int{1, 1, 2, 1, 3, 2}}
	h2 := business.Horse{Number: "023", Name: "beta", OwnerName: "IVANOV IVAN", RaceHistory: []int{2, 3, 1, 2, 1, 4}}
	h3 := business.Horse{Number: "007", Name: "gamma", OwnerName: "IVANOV IVAN", RaceHistory: []int{2, 3, 3, 1, 4}}

	res := business.MakeRace([]business.Horse{h1, h2, h3})
	business.ReadRace(res)
}
