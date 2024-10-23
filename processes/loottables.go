package processes

import (
	"fmt"

	"github.com/allensg/dndhelpers/types"
)

func artTables(rolls types.Rolls) *types.TableResults {
	results := &types.TableResults{}

	return results
}

func gemTables(rolls types.Rolls) *types.TableResults {
	results := &types.TableResults{}

	return results
}

func generateBossLoot(cr types.CrRange) *types.TableResults {
	results := &types.TableResults{}

	return results
}

// Ancient Dragon loot suggestions
// 4,200 (12d6 × 100) cp
// 14,000 (4d6 × 1,000) sp
// 210,000 (6d6 × 10,000) gp
// 42,000 (12d6 × 1,000) pp
// 9 (2d8) mundane items
// 21 (6d6) gems
// 11 (2d10) art objects
// 7 (2d6) magic items
func GenerateAncientDragonHoard() *types.TableResults {
	//resultArr := []*types.TableResult{}
	// roll xdx times on the x table
	artTimesDie := &types.Die{
		Size: 10,
	}
	artTimesResult := 0
	for i := 0; i < 2; i++ {
		artTimesResult += artTimesDie.RollDie()
	}
	fmt.Println("artTimesResult")
	fmt.Println(artTimesResult)

	AncientArtTable := types.NewAncientDragonArtTable()
	artRoll := types.Roll{
		Die:   types.Die{Size: types.D100},
		Count: 1,
	}

	artRolls := []types.Roll{}
	for i := 0; i < artTimesResult; i++ {

	}

	artRollTypes := types.Rolls{Components: artRolls}
	artResults, err := types.HandleTableRolls(AncientArtTable, artRollTypes)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// resultArr = append(resultArr, artResults.Results...)

	// return &types.TableResults{
	// 	Results: resultArr,
	// }
	return nil
}
