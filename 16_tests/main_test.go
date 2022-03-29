package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tests are cept in the same folder
// filenames have to end in '_test.go'
// Test functions have to start with 'Test' and take a  *testing.T argument as input

// use go test explorer, to get a similar panel to the vs test explorer
func TestNormal(t *testing.T) {
	result := 1 + 1
	assert.Equal(t, 2, result)
}

func TestTableDriven(t *testing.T) {
	testCases := []struct {
		desc          string
		firstOperand  int
		secondOperand int
		result        int
	}{
		{
			desc:          "1+1",
			firstOperand:  1,
			secondOperand: 1,
			result:        3,
		},
		{
			desc:          "2+2",
			firstOperand:  2,
			secondOperand: 2,
			result:        4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.result, tC.firstOperand+tC.secondOperand)
		})
	}
}
