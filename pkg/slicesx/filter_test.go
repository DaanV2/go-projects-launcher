package slicesx_test

import (
	"slices"
	"testing"

	"github.com/DaanV2/go-projects-launcher/slicesx"
	"github.com/stretchr/testify/assert"
)

func Test_Filter(t *testing.T) {
	data := []string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}

	for _, i := range data {
		filtered := slicesx.Filter(data, func(item string) bool {
			return item != i
		})

		assert.Len(t, filtered, len(data)-1)
		assert.False(t, slices.Contains(filtered, i))
	}
}
