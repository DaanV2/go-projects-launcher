package slicesx_test

import (
	"testing"

	"github.com/DaanV2/go-projects-launcher/slicesx"
	"github.com/stretchr/testify/assert"
)

func Test_Unigue(t *testing.T) {
	for i := range 10 {
		data := make([]int, 0)

		for i := range i {
			data = append(data, i)
		}

		filtered := slicesx.Unique(data)
		assert.Len(t, filtered, i)
	}
}