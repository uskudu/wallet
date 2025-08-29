package business

import (
	"fmt"
	"sort"
)

func MakeRace(horses []Horse) []Horse {
	// collect chances data
	for i := 0; i < len(horses); i++ {
		var ws float32
		ws = 1
		// each w in race history gives chance + 1
		for _, res := range horses[i].RaceHistory {
			if res == "w" {
				ws++
			} else {
				ws--
			}
		}
		if ws < 1 {
			ws = 1
		}
		horses[i].Chance = ws
	}
	// result
	sort.Slice(horses, func(i, j int) bool {
		return horses[i].Chance > horses[j].Chance
	})
	return horses
}

func ReadRace(raceResult []Horse) {
	fmt.Println("Race result:")
	for i, v := range raceResult {
		fmt.Printf("%d: name: %v | number: %v | owner: %v\n", i+1, v.Name, v.Number, v.OwnerName)
	}
}
