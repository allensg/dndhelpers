package types

func NewWyrmlingArtTable() *Table {
	content := NewArtTableRows()

	bounds := []Bounds{
		{Lower: 1, Upper: 95},
		{Lower: 96, Upper: 100},
	}

	wyrmlingArtTable := &Table{
		Name:    "Wyrmling Art Objects Table",
		Bounds:  bounds,
		Content: content,
	}

	return wyrmlingArtTable
}

func NewAncientDragonArtTable() *Table {
	content := NewArtTableRows()

	bounds := []Bounds{
		{Lower: 1, Upper: 22},
		{Lower: 23, Upper: 42},
		{Lower: 43, Upper: 58},
		{Lower: 59, Upper: 93},
		{Lower: 94, Upper: 100},
	}

	wyrmlingArtTable := &Table{
		Name:    "Wyrmling Art Objects Table",
		Bounds:  bounds,
		Content: content,
	}

	return wyrmlingArtTable
}
