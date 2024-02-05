package typeutl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	cases := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{
			name: "same",
			a:    []int{1, 2, 3},
			b:    []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "first empty",
			a:    []int{},
			b:    []int{1, 2, 3},
			want: nil,
		},
		{
			name: "second empty",
			a:    []int{1, 2, 3},
			b:    []int{},
			want: nil,
		},
		{
			name: "simple 1",
			a:    []int{1, 2, 3, 4, 5},
			b:    []int{1, 3, 5, 7, 9},
			want: []int{1, 3, 5},
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			got := Intersect(cc.a, cc.b)
			assert.Equal(t, cc.want, got)
		})
	}
}

func TestIntersectStruct(t *testing.T) {
	type S struct {
		id   int
		name string
	}

	cases := []struct {
		name string
		a    []S
		b    []S
		want []S
	}{
		{
			name: "compare structs",
			a:    []S{{0, "A"}, {1, "B"}, {2, "C"}, {3, "D"}},
			b:    []S{{0, "A"}, {10, "B"}, {2, "C"}, {3, "DD"}},
			want: []S{{0, "A"}, {2, "C"}},
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			got := Intersect(cc.a, cc.b)
			assert.Equal(t, cc.want, got)
		})
	}
}
