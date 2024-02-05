package typeutl

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncMap(t *testing.T) {
	m := NewSyncMap[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)

	v, ok := m.Get("a")
	assert.Equal(t, 1, v)
	assert.Equal(t, true, ok)

	v, ok = m.Get("b")
	assert.Equal(t, 2, v)

	v, ok = m.Get("c")
	assert.Equal(t, 3, v)

	v, ok = m.Get("d")
	assert.False(t, ok)

	values := m.Values()
	wantValues := []int{1, 2, 3}
	for _, v := range values {
		assert.True(t, slices.Contains(wantValues, v))
	}

	m.Delete("a")
	v, ok = m.Get("a")
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}
