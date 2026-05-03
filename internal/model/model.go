package model

import "time"

type Player struct {
	Name  string `json:"name"`
	HP    int    `json:"hp"`
	MaxHP int    `json:"max_hp"`
}

type Plot struct {
	CropID    string    `json:"crop_id"`
	PlantedAt time.Time `json:"planted_at"`
}

type Farm struct {
	Plots []Plot `json:"plots"`
}

type BreakState struct {
	Active        bool      `json:"active"`
	LastStartedAt time.Time `json:"last_started_at"`
	LastEndedAt   time.Time `json:"last_ended_at"`
	LockedUntil   time.Time `json:"locked_until"`
}

type Summary struct {
	Harvested     int `json:"harvested"`
	Crafted       int `json:"crafted"`
	BugsFixed     int `json:"bugs_fixed"`
	TestsRestored int `json:"tests_restored"`
}

type Unlocked struct {
	Crops  []string `json:"crops"`
	Spells []string `json:"spells"`
}

type Save struct {
	Player        Player         `json:"player"`
	Inventory     map[string]int `json:"inventory"`
	Farm          Farm           `json:"farm"`
	Break         BreakState     `json:"break"`
	Summary       Summary        `json:"summary"`
	Unlocked      Unlocked       `json:"unlocked"`
	NextRaidBonus int            `json:"next_raid_bonus"`
}

func DefaultSave() Save {
	return Save{
		Player: Player{
			Name:  "Sir GREPalot",
			HP:    10,
			MaxHP: 10,
		},
		Inventory: map[string]int{
			"coffee":        2,
			"regex_root":    1,
			"stack_crystal": 1,
		},
		Farm: Farm{
			Plots: []Plot{{}, {}, {}, {}},
		},
		Unlocked: Unlocked{
			Crops: []string{
				"coffee",
				"regex_root",
				"stack_crystal",
				"heap_mushroom",
				"mutex_mint",
			},
			Spells: []string{},
		},
	}
}
