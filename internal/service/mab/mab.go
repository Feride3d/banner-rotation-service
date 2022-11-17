package mab

import (
	"fmt"
	"math"
	"math/rand"
)

var ErrZeroDisplays = fmt.Errorf("there are no displays for banners")

type Mab struct { // MultiArmedBandit
	Display int // amount of displays banners to the group
	Click   int // amount of clicks on banners
}

func NewMab() *Mab {
	return &Mab{}
}

// Ucb represents the upper confidence bound algorithm
func (m *Mab) Ucb(banners []string, clicks map[string]int, displays map[string]int) (string, error) {
	amounts := make(map[string]float64)

	if err := m.ZeroDisplays(banners, displays); err != nil {
		return "", err
	}

	for _, banner := range banners {
		amounts[banner] = m.GetAmount(float64(displays[banner]), float64(clicks[banner]), float64(len(displays)))
	}
	maxClick := m.MaxClick(amounts)
	maxClickedBanners := m.SelectMaxClickedBanners(amounts, maxClick)
	bannerID := m.SelectOne(maxClickedBanners)

	return bannerID, nil
}

func (m *Mab) GetAmount(displays float64, clicks float64, total float64) float64 {
	doneClick := clicks / displays
	rotation := math.Sqrt(((2 * math.Log(total)) / displays))

	return doneClick + rotation
}

func (m *Mab) MaxClick(rotations map[string]float64) float64 {
	max := 0.0
	for _, k := range rotations {
		if k > max {
			max = k
		}
	}

	return max
}

func (m *Mab) SelectMaxClickedBanners(amounts map[string]float64, maxAmount float64) []string {
	maxClickedBanners := []string{}
	for k, amount := range amounts {
		if amount == maxAmount {
			maxClickedBanners = append(maxClickedBanners, k)
		}
	}
	return maxClickedBanners
}

func (m *Mab) SelectOne(banners []string) string {
	if len(banners) == 0 {
		return ""
	}
	rand.Shuffle(len(banners), func(i, j int) { banners[i], banners[j] = banners[j], banners[i] })
	return banners[0]
}

func (m *Mab) ZeroDisplays(banners []string, displays map[string]int) error {
	for _, banner := range banners {
		if displays[banner] == 0 {
			return ErrZeroDisplays
		}
	}
	return nil
}
