package main

import (
	"fmt"

	"github.com/allensg/dndhelpers/types"
)

func main() {
	fmt.Println("sup sluts")

	wyrmlingTable := types.NewWyrmlingArtTable()

	rolls := []types.Roll{}

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
