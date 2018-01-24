package str

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex(t *testing.T) {
	var tests = []struct {
		s   string
		sep string
		out int
	}{
		{"", "", -1},
		{"", "a", -1},
		{"fo", "foo", -1},
	}
	for _, test := range tests {
		index := Index(test.s, test.sep)
		assert.Equal(t, test.out, index)
	}
}
