package types

import (
	"errors"
	"fmt"
	"math/rand"
)

type TableType string

const (
	D2   = 2
	D4   = 4
	D6   = 6
	D8   = 8
	D10  = 10
	D12  = 12
	D20  = 20
	D100 = 100
	// table types
	TableGoldValue  TableType = "TableGoldValue"  // gp value of item
	TableValueLink  TableType = "TableValueLink"  // hyperlink to an item
	TableOtherTable TableType = "TableOtherTable" // links to another table
)

type Die struct {
	Size int
}

type Roll struct {
	Die   Die
	Count int
}

type Rolls struct {
	Components []Roll
}

// the idea here is that you roll a table, check the bounds and the index of the found bound is the index of the object to check
type Table struct { // ex
	Name    string    // 'art objects', gems, magic items, magic item tables
	Bounds  []Bounds  //
	Content TableRows // 100gp, Magic item table A, https://www.dndbeyond.com/magic-items/4581-bag-of-holding, etc
}

type TableRows struct {
	Type          TableType // gp, item links, table links
	Links         []string
	ItemValues    []string
	LinkingTables []Table
}

type TableResult struct { // like table but for resultsd
	Type         TableType
	Link         string
	ItemValue    string
	LinkingTable Table
}

type TableResults struct {
	Results []*TableResult
}

type Bounds struct {
	Lower int
	Upper int
}

type TableInterface interface {
	HandleTableRolls(rolls Rolls) (*TableResults, error)
}

func NewArtTableRows() TableRows {
	rows := TableRows{
		Type:       TableGoldValue,
		ItemValues: []string{"25", "250", "750", "2500", "7500"},
	}

	return rows
}

func RollOnTable(table *Table, roll Roll) (*TableResult, error) {
	result := &TableResult{
		Type: table.Content.Type,
	}

	fmt.Println("inside roll on table")

	sum := 0
	for i := 0; i < roll.Count; i++ {
		rollResult := roll.Die.RollDie()
		sum = sum + rollResult
	}

	fmt.Println("roll result")
	fmt.Println(sum)
	for boundIndex, bound := range table.Bounds {
		if bound.InBounds(sum) {
			switch table.Content.Type {
			case TableValueLink:
				result.Link = table.Content.Links[boundIndex]
			case TableGoldValue:
				result.ItemValue = table.Content.ItemValues[boundIndex]
			case TableOtherTable:
				result.LinkingTable = table.Content.LinkingTables[boundIndex]
			default:
				return nil, errors.New("invalid table type")
			}
		}
	}

	return result, nil
}

func (r *TableResults) PrintTableRolls() {
	for _, result := range r.Results {
		fmt.Println(result.Type)
		switch result.Type {
		case TableValueLink:
			fmt.Println(result.Link)
		case TableGoldValue:
			fmt.Println(result.ItemValue)
		case TableOtherTable:
			fmt.Println(result.LinkingTable.Name)
		default:
			fmt.Println("invalid table type sluts")
		}
	}
}

func (b Bounds) InBounds(roll int) bool {
	if b.Upper >= roll && b.Lower <= roll {
		return true
	}
	return false
}

func (d *Die) RollDie() int {
	return rand.Intn(d.Size)
}

func HandleTableRolls(table *Table, rolls Rolls) (*TableResults, error) {
	tableResults := []*TableResult{}
	for _, roll := range rolls.Components {
		fmt.Println("inside handle table rolls")
		result, err := RollOnTable(table, roll)
		if err != nil {
			return nil, errors.Join(err, errors.New("Failed to roll"))
		}
		fmt.Println()
		tableResults = append(tableResults, result)
	}
	return &TableResults{
		Results: tableResults,
	}, nil
}
