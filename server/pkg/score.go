package pkg

import "math"

func CalculateScore(voteAverage float64) float64 {
	return math.Round(voteAverage) / 10.0
}
