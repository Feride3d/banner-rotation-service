package mab

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxCLickableBanner(t *testing.T) {
	clicks := make(map[BannerID]Mab)
	clicks[1] = Mab{Display: 10, Click: 0}
	clicks[2] = Mab{Display: 10, Click: 3}
	clicks[3] = Mab{Display: 10, Click: 5}
	clicks[4] = Mab{Display: 10, Click: 10}
	t.Run("banner 4", func(t *testing.T) {
		rotation := Ucb(clicks, 10)
		require.Equal(t, 4, int(rotation))
	})

}

func TestZeroValue(t *testing.T) {
	clicks := make(map[BannerID]Mab)
	clicks[1] = Mab{Display: 0, Click: 0}
	clicks[2] = Mab{Display: 0, Click: 0}
	clicks[3] = Mab{Display: 0, Click: 0}
	clicks[4] = Mab{Display: 0, Click: 0}
	t.Run("zero display and value", func(t *testing.T) {
		rotation := Ucb(clicks, 0)
		require.Equal(t, 0, int(rotation))
	})
}
