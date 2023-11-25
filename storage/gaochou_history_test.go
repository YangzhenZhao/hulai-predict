package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetGaochouHistory(t *testing.T) {
	initGaochouHisotyMap()
	historys := GetGaochouHistory("wei")
	assert.Equal(t, true, len(historys) >= 10)
	for _, history := range historys {
		if history.Date.Equal(newDate(2023, 5, 12)) {
			assert.Equal(t, history.FirstHero, "曹丕")
			assert.Equal(t, history.SecondHero, "夏侯惇")
		}
	}
}
