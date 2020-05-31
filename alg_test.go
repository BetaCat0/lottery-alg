package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPickWinners(t *testing.T) {
	min, max, num := 0, 5000, 3
	key := "0000000000000000000eaee9dda170aa083087d0dad1197a3109cdcfdceaa54a"

	assert.True(t, equals([]int{1110, 1388, 1351}, PickWinners(min, max, num, key)))

	min, max, num = 0, 500, 3
	key = "00000000000000000005a13a533eac084091170c5e21a41d00494ce9a7aabc9f"

	assert.True(t, equals([]int{237, 39, 315}, PickWinners(min, max, num, key)))

	min, max, num = 0, 50, 3
	key = "000000000000000000110638fd09e10f662c850fd144d31d9b188b0cc96fd1af"

	assert.True(t, equals([]int{18, 1, 5}, PickWinners(min, max, num, key)))
}

func equals(expected, actual []int) bool {
	for i := range expected {
		hit := func(i int, actual []int) bool {
			for ii := range actual {
				if i == actual[ii] {
					return true
				}
			}
			return false
		}(expected[i], actual)
		if !hit {
			return false
		}
	}
	return true
}

func BenchmarkPickWinners(b *testing.B) {
	min, max, num := 1, 50, 10
	key := "0000000000000000000eaee9dda170aa083087d0dad1197a3109cdcfdceaa54a"

	PickWinners(min, max, num, key)
}
