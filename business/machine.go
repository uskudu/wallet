package business

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func calcChance(history []int) float64 {
	if len(history) == 0 {
		return 1 // rookie horse
	}

	N := len(history)
	P := 0   // first places from horse history
	sum := 0 // sum of all races from horse history
	for _, place := range history {
		if place == 1 {
			P++
		}
		sum += place
	}
	Pavg := float64(sum) / float64(N) // average place from horse history

	// coefficients weights
	alpha := 50.0 // 1 places importance
	beta := 30.0  // avg place importance
	gamma := 10.0 // exp importance
	delta := 10.0 // random importance

	chance := alpha*(float64(P)/float64(N+1)) +
		beta*(1.0/Pavg) +
		gamma*math.Log(float64(N)+1) +
		delta*randFloat(0, 1)

	if chance < 1 {
		chance = 1
	}
	return chance
}

func MakeRace(horses []Horse) []Horse {
	for i := range horses {
		horses[i].Chance = calcChance(horses[i].RaceHistory)
	}

	// result sort
	sort.Slice(horses, func(i, j int) bool {
		return horses[i].Chance > horses[j].Chance
	})

	// add race to horses history
	for i := range horses {
		horses[i].RaceHistory = append(horses[i].RaceHistory, i+1)
	}
	return horses
}

func ReadRace(raceResult []Horse) {
	fmt.Println("Race result:")
	for i, v := range raceResult {
		fmt.Printf("%d: name: %v | number: %v | owner: %v\n", i+1, v.Name, v.Number, v.OwnerName)
	}
	fmt.Println(raceResult)
}
