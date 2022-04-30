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
		rotation := Ucb(clicks, 10) // check formula
		require.Equal(t, 4, int(rotation))
	})

}

func TestConcurrency(t *testing.T) {
	// добавить тесты на пограничные и отрицательные значения, конкурентную обработку запросов
}
