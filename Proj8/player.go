package main

type Player struct {
	GridX        int         `json:"gridX"`
	GridY        int         `json:"gridY"`
	IsOutside    bool        `json:"isOutside"`
	Seeds        map[int]int `json:"seeds"`
	WaterCanUses int         `json:"waterCanUses"`
}
