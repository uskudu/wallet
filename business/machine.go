package business

import (
	"fmt"
	"sort"
)

func MakeRace(horses []Horse) []Horse {
	chances := make([]int, len(horses))

	// collect chances data
	for i := 0; i < len(horses); i++ {
		// if no race history change = 1
		if len((horses)[i].RaceHistory) < 1 {
			chances[i] = 1
		} else {
			// each w in race history gives chance + 1
			ws := 1
			for _, res := range (horses)[i].RaceHistory {
				if res == "w" {
					ws++
				} else {
					ws--
				}
			}
			if ws < 1 {
				ws = 1
			}
			chances[i] = ws
		}
	}
	// win probability adjustment

	// result

	sort.Slice(horses, func(i, j int) bool {
		return chances[i] > chances[j]
	})
	return horses
}

func ReadRace(raceResult []Horse) {
	fmt.Println("Race result:")
	for i, v := range raceResult {
		fmt.Printf("%d: number: %v | name: %v | owner: %v\n", i+1, v.Number, v.Name, v.OwnerName)
	}
}
