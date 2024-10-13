package main

import (
	"fmt"

	"github.com/allensg/dndhelpers/types"
)

func main() {
	fmt.Println("sup sluts")

	wyrmlingTable := types.NewWyrmlingArtTable()
	roll := types.Roll{
		Die:   types.Die{Size: types.D10},
		Count: 2,
	}
	rolls := []types.Roll{roll}

	rollTypes := types.Rolls{
		Components: rolls,
	}

	results, err := wyrmlingTable.HandleTableRolls(rollTypes)

	if err != nil {
		fmt.Println(err.Error())
	}

	results.PrintTableRolls()

	fmt.Println("blargh")
}
