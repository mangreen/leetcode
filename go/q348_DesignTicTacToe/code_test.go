package q348_DesignTicTacToe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_encodeAndDecode(t *testing.T) {
	assert := assert.New(t)

	obj := Constructor(3)

 	move1 := obj.Move(0, 0, 1)
	assert.Equal(0, move1, `[[x], [], []] Player 1 makes a move at (0, 0)`)

	move2 := obj.Move(0, 2, 2)
	assert.Equal(0, move2, `[[x, , o], [], []] Player 2 makes a move at (0, 2)`)

	move3 := obj.Move(2, 2, 1)
	assert.Equal(0, move3, `[[x, , o], [], [ , , x]] Player 1 makes a move at (2, 2)`)

	move4 := obj.Move(1, 1, 2)
	assert.Equal(0, move4, `[[x, , o], [, o, ], [, , x]] Player 2 makes a move at (1, 1)`)

	move5 := obj.Move(2, 0, 1)
	assert.Equal(0, move5, `[[x, , o], [, o, ], [x, , x]] Player 1 makes a move at (2, 0)`)

	move6 := obj.Move(1, 0, 2)
	assert.Equal(0, move6, `[[x, , o], [o, o, ], [x, , x]] Player 2 makes a move at (1, 0)`)

	move7 := obj.Move(2, 1, 1)
	assert.Equal(1, move7, `[[x, , o], [o, o, ], [x, x, x]] Player 1 makes a move at (2, 1) and won`)
}