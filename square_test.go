package square

import (
	"testing"

	"gotest.tools/assert"
)

func TestEnd(t *testing.T) {
	square := Square{Point{1, 2}, 3}
	end := square.End()
	assert.Equal(t, end, Point{4, 5},
		`End differs for 'Point{%v, %v}, %v': '%v' not '%v'`,
		1, 2, 3, end, Point{4, 5},
	)
}

func TestArea(t *testing.T) {
	square := Square{Point{1, 2}, 3}
	area := square.Area()
	assert.Equal(t, area, 9,
		`Area() differs for 'Point{%v, %v}, %v': '%v' not '%v'`,
		1, 2, 3, area, 9,
	)
}
