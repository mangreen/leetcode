package q362_DesignHitCounter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hitCounter(t *testing.T) {
	assert := assert.New(t)

	counter := Constructor()

	counter.Hit(1)
	counter.Hit(2)
	counter.Hit(3)
	
	assert.Equal(counter.GetHits(4), 3, "It should return 3")

	counter.Hit(300)

	assert.Equal(counter.GetHits(300), 4, "It should return 4")
	assert.Equal(counter.GetHits(301), 3, "It should return 3, as hit at timestamp 1 is now outside the window")
}
