package fuzzyLogic

import (
	"akim/internal/domain/model"
)

func CalculateMembership(building *model.FuzzyArchitecturalArtifact, intervalStart, intervalEnd int) float64 {
	if building.IntervalEnd == 0 {
		return 1 / float64(intervalEnd-intervalStart)
	}

	if intervalStart <= building.IntervalStart && building.IntervalStart <= building.IntervalEnd && building.IntervalEnd <= intervalEnd {
		return float64(building.IntervalEnd-building.IntervalStart) / float64(intervalEnd-intervalStart)
	} else if building.IntervalStart < intervalStart && intervalStart <= building.IntervalEnd && building.IntervalEnd <= intervalEnd {
		return float64(building.IntervalEnd-intervalStart) / float64(intervalEnd-building.IntervalStart)
	} else if intervalStart <= building.IntervalStart && building.IntervalStart <= intervalEnd && intervalEnd < building.IntervalEnd {
		return float64(intervalEnd-building.IntervalStart) / float64(building.IntervalEnd-intervalStart)
	} else if building.IntervalStart < intervalStart && intervalStart <= intervalEnd && intervalEnd < building.IntervalEnd {
		return float64(intervalEnd-intervalStart) / float64(building.IntervalEnd-building.IntervalStart)
	} else {
		return 0.0
	}
}
