package indicators

import (
	helpers "gallywix.elijahelrod.com/internal/helpers"
)

type VolumeTrend struct {
	trend              helpers.Trend
	averageDailyVolume uint64
	currentDailyVolume uint64
}
