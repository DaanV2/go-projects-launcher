package slicesx_test

import (
	"slices"
	"testing"

	"github.com/DaanV2/go-projects-launcher/slicesx"
	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	type TC struct {
		ID string
		V  int
	}

	data := []TC{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
	}

	mappedID := slicesx.Map(data, func(item TC) string {
		return item.ID
	})
	mappedV := slicesx.Map(data, func(item TC) int {
		return item.V
	})

	for _, i := range data {
		assert.True(t, slices.Contains(mappedID, i.ID))
		assert.True(t, slices.Contains(mappedV, i.V))
	}
}
