package t24

import (
	"math"
	"testing"
)

func TestPoint_Distance(t *testing.T) {
	p1 := NewPoint(6, 2)
	p2 := NewPoint(-1, 3)
	if got := p1.Distance(p2); got != 5*math.Sqrt(2) {
		t.Errorf("got %f, expected %f", got, 5*math.Sqrt(2))
	}
}
