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
	h1 := business.Horse{Number: "077", Name: "alfa", OwnerName: "IVANOV IVAN", RaceHistory: []string{"w", "l"}}
	h2 := business.Horse{Number: "023", Name: "beta", OwnerName: "IVANOV IVAN", RaceHistory: []string{}}
	h3 := business.Horse{Number: "007", Name: "gamma", OwnerName: "IVANOV IVAN", RaceHistory: []string{"w", "w"}}

	res := business.MakeRace([]business.Horse{h1, h2, h3})
	business.ReadRace(res)
}
