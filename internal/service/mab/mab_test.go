package mab

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxCLickedBanners(t *testing.T) {
	mab := NewMab()
	t.Run("test banner with max click", func(t *testing.T) {
		amounts := map[string]float64{"banner1": 0.5, "banner2": 0.15, "banner3": 0.2}
		maxClickedBanners := mab.SelectMaxClickedBanners(amounts, 0.5)
		require.Len(t, maxClickedBanners, 1)
		require.Contains(t, maxClickedBanners, "banner1")
	})
}

func TestMaxClick(t *testing.T) {
	mab := NewMab()
	t.Run("test getting max click amount", func(t *testing.T) {
		amounts := map[string]float64{"banner1": 0.5, "banner2": 0.15, "banner3": 0.2}
		maxClick := mab.MaxClick(amounts)
		require.Equal(t, 0.5, maxClick)
	})
}

func TestZeroDisplays(t *testing.T) {
	mab := NewMab()
	t.Run("test zero displays", func(t *testing.T) {
		banner, err := mab.Ucb([]string{"banner1", "banner2", "banner3"}, map[string]int{}, map[string]int{})
		require.ErrorIs(t, err, ErrZeroDisplays)
		require.Empty(t, banner)
	})
}
