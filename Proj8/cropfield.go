package main

type Crop struct {
	SeedType int
	Growth   int
	Value    int
	Watered  bool
}

type CropField struct {
	Plots [FarmSize][FarmSize]Crop
}

func NewCropField() CropField {
	return CropField{}
}

func (cf *CropField) Plant(x, y, seedType int, bonus int) {
	cf.Plots[y][x] = Crop{
		SeedType: seedType,
		Growth:   0,
		Value:    getBaseValue(seedType) + bonus,
		Watered:  false,
	}
}

func (cf *CropField) GrowAll() {
	for y := range cf.Plots {
		for x := range cf.Plots[y] {
			if cf.Plots[y][x].Watered && cf.Plots[y][x].Growth < 3 {
				cf.Plots[y][x].Growth++
			}
			cf.Plots[y][x].Watered = false
		}
	}
}

func getBaseValue(seedType int) int {
	switch seedType {
	case 1:
		return 5
	case 2:
		return 20
	case 3:
		return 40
	}
	return 0
}
