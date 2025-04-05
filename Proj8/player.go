package main

type Player struct {
	GridX        int
	GridY        int
	IsOutside    bool
	Seeds        map[int]int
	WaterCanUses int
}
