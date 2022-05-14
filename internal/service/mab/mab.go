package mab

import (
	"math"
)

type BannerID int

type Mab struct { // MultiArmedBandit
	Display int // amount of displays banners to the group
	Click   int // amount of clicks on banners
}

// Ucb represents the upper confidence bound algorithm
func Ucb(amount map[BannerID]Mab, displays int) BannerID {
	rotations := make(map[float64]BannerID)
	for k, v := range amount {
		doneCLick := float64(v.Click) / float64(v.Display)
		rotation := doneCLick + math.Sqrt(2*math.Log(float64(displays))/float64(v.Display))
		rotations[rotation] = k
	}
	max := maxClick(rotations)

	return rotations[max]
}

func maxClick(rotations map[float64]BannerID) float64 {
	max := 0.0
	for k := range rotations {
		if k > max {
			max = k
		}
	}

	return max
}
