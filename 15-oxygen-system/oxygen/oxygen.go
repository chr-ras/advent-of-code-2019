package oxygen

import (
	"github.com/chr-ras/advent-of-code-2019/15-oxygen-system/repairdroid"
	"github.com/chr-ras/advent-of-code-2019/util/geometry"
	"github.com/gosuri/uilive"
)

// SimulateOxygenExpansion simulates the expansion of oxyugen from the oxygen station through the space ship.
func SimulateOxygenExpansion(shipMap map[geometry.Vector]repairdroid.Position, oxygenStation repairdroid.OxygenStation) int {
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	oxygenStationPositionInfo := shipMap[oxygenStation.Position]
	oxygenStationPositionInfo.Status = repairdroid.Oxygen
	shipMap[oxygenStation.Position] = oxygenStationPositionInfo

	currentlyExpandingPositions := []geometry.Vector{oxygenStation.Position}

	for minutes := 0; true; minutes++ {
		var newExpandingPositions []geometry.Vector

		for _, expandingPosition := range currentlyExpandingPositions {
			expandingPositionInfo := shipMap[expandingPosition]

			for _, adjacentPosition := range expandingPositionInfo.AdjacentPositions {
				adjacentPositionInfo := shipMap[adjacentPosition]

				if adjacentPositionInfo.Status == repairdroid.Oxygen || adjacentPositionInfo.Status == repairdroid.HitWall {
					continue
				}

				adjacentPositionInfo.Status = repairdroid.Oxygen
				shipMap[adjacentPosition] = adjacentPositionInfo

				newExpandingPositions = append(newExpandingPositions, adjacentPosition)
			}
		}

		repairdroid.PrettyPrint(geometry.Vector{}, shipMap, writer, 100)

		currentlyExpandingPositions = newExpandingPositions

		if len(currentlyExpandingPositions) == 0 {
			return minutes
		}
	}

	return 0
}
