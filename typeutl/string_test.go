package typeutl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyToStr(t *testing.T) {
	tests := []struct {
		name string
		arg  any
		want string
	}{
		{"string", "string", "string"},
		{"int", 1, "1"},
		{"bytes", []byte("bytes"), "bytes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, AnyToStr(tt.arg))
		})
	}
}
