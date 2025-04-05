package main

type Crop struct {
	SeedType int  `json:"seedType"`
	Growth   int  `json:"growth"`
	Value    int  `json:"value"`
	Watered  bool `json:"watered"`
}

type CropField struct {
	Plots [FarmSize][FarmSize]Crop `json:"plots"`
}

func NewCropField() CropField {
	return CropField{}
}

func (cf *CropField) Plant(x, y, seedType int, bonus int) {
	if x < 0 || x >= FarmSize || y < 0 || y >= FarmSize {
		return
	}
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
			crop := &cf.Plots[y][x]
			if crop.SeedType == 0 {
				continue
			}

			growthIncrement := 1
			if crop.Watered {
				growthIncrement++
			}

			crop.Growth += growthIncrement

			switch crop.SeedType {
			case 1: //Radish
				if crop.Growth > 2 {
					crop.Growth = 2
				}
			case 2: //wheat
				if crop.Growth > 3 {
					crop.Growth = 3
				}
			default: //cotton
				if crop.Growth > 4 {
					crop.Growth = 4
				}
			}

			crop.Watered = false
		}
	}
}

func (cf *CropField) TryHarvest(x, y int) int {
	if x < 0 || x >= FarmSize || y < 0 || y >= FarmSize {
		return 0
	}

	crop := cf.Plots[y][x]

	var requiredGrowth int
	switch crop.SeedType {
	case 1:
		requiredGrowth = 2
	case 2:
		requiredGrowth = 3
	default:
		requiredGrowth = 4
	}

	if crop.SeedType > 0 && crop.Growth >= requiredGrowth {
		cf.Plots[y][x] = Crop{}
		return crop.Value
	}
	return 0
}

func getBaseValue(seedType int) int {
	switch seedType {
	case 1:
		return 5 // Radish
	case 2:
		return 20 // Wheat
	case 3:
		return 40 // Cotton
	}
	return 0
}
