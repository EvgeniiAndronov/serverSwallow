package server

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func Test_CheckHit(t *testing.T) {
	testCases := []struct {
		name     string
		bx       float32
		by       float32
		tx       float32
		ty       float32
		expected bool
	}{
		{
			name:     "in range",
			bx:       0,
			by:       0,
			tx:       0,
			ty:       0,
			expected: true,
		},
		{
			name:     "in range",
			bx:       5,
			by:       5,
			tx:       0,
			ty:       0,
			expected: true,
		},
		{
			name:     "out of range",
			bx:       0,
			by:       0,
			tx:       -120,
			ty:       -120,
			expected: false,
		},
		{
			name:     "out of range",
			bx:       -120,
			by:       120,
			tx:       0,
			ty:       0,
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := CheckHit(testCase.bx, testCase.by, testCase.tx, testCase.ty)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
