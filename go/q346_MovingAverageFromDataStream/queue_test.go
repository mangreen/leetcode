package q346_MovingAverageFromDataStream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_movingAverage(t *testing.T) {
	assert := assert.New(t)

	ma := Constructor(3)

	assert.Equal(ma.Next(1), 1.0, "1/1 = 1.0")
	assert.Equal(ma.Next(10), 5.5, "(1+10)/2 = 5.5")
	assert.Equal(ma.Next(3), 4.666666666666667, "(1+10+3) = 4.666666666666667")
	assert.Equal(ma.Next(5), 6.0, "(10+3+5) = 6.6")
}
