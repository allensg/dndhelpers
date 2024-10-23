package types

func NewWyrmlingGemTable() *Table {
	content := NewArtTableRows()

	bounds := []Bounds{
		{Lower: 1, Upper: 43},
		{Lower: 44, Upper: 99},
		{Lower: 100, Upper: 100},
	}

	return &Table{
		Name:    "Wyrmling Gem Table",
		Bounds:  bounds,
		Content: content,
	}
}

func NewYoungDragonGemTable() *Table {
	content := NewArtTableRows()

	bounds := []Bounds{
		{Lower: 1, Upper: 51},
		{Lower: 52, Upper: 75},
		{Lower: 76, Upper: 99},
		{Lower: 100, Upper: 100},
	}

	return &Table{
		Name:    "Young Dragon Gem Table",
		Bounds:  bounds,
		Content: content,
	}
}

func NewAdultDragonGemTable() *Table {
	content := NewArtTableRows()

	bounds := []Bounds{
		{Lower: 1, Upper: 18},
		{Lower: 19, Upper: 36},
		{Lower: 37, Upper: 54},
		{Lower: 55, Upper: 77},
		{Lower: 77, Upper: 99},
		{Lower: 100, Upper: 100},
	}

	return &Table{
		Name:    "Adult Dragon Gem Table",
		Bounds:  bounds,
		Content: content,
	}
}

func NewAncientDragonGemTable() *Table {
	content := NewArtTableRows()

	bounds := []Bounds{
		{Lower: 1, Upper: 14},
		{Lower: 15, Upper: 28},
		{Lower: 29, Upper: 42},
		{Lower: 43, Upper: 58},
		{Lower: 59, Upper: 93},
		{Lower: 94, Upper: 100},
	}

	return &Table{
		Name:    "Ancient Dragon Gem Table",
		Bounds:  bounds,
		Content: content,
	}

}
