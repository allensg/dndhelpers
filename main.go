package main

import (
	"fmt"

	"github.com/allensg/dndhelpers/types"
)

func main() {
	fmt.Println("sup sluts")

	wyrmlingTable := types.NewAncientDragonArtTable()
	roll := types.Roll{
		Die:   types.Die{Size: types.D100},
		Count: 1,
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
