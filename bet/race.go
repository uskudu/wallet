package bet

import "time"

type Horse struct {
	Number    int
	Name      string
	OwnerName string
}

type Race struct {
	Horses     []Horse
	TotalLaps  int
	CurrentLap int
	PrizePool  int
}

type RaceResult struct {
	Race         Race
	Duration     time.Duration
	GoldWinner   Horse
	SilverWinner Horse
	BronzeWinner Horse
}
