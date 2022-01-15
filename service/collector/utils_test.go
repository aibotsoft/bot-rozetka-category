package collector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_difference(t *testing.T) {
	got := difference([]int64{1, 2, 3, 4}, []int64{3, 4})
	assert.Equal(t, []int64{1, 2}, got)
}
