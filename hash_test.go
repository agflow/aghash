package aghash

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	type ints struct{ ints []int }
	someInt := 2
	s := struct {
		a map[string]int
		b struct{}
		c *int
		d []float64
		e ints
		f bool
	}{
		a: map[string]int{"key1": 1},
		b: struct{}{},
		c: &someInt,
		d: []float64{math.Pi},
		e: ints{[]int{1, 2, 3}},
		f: true,
	}

	initial, err := Hash(s)
	require.Nil(t, err)
	second, err := Hash(s)
	require.Nil(t, err)

	require.Equal(t, initial, second)

	s.f = false
	after, err := Hash(s)
	require.Nil(t, err)
	require.NotEqual(t, initial, after)
}
