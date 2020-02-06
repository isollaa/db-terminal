package helper

import (
	"testing"

	"github.com/isollaa/dbterm/config"
)

type testCase struct {
	Result interface{}
	valid  []bool
}

var testCases = testCase{
	Result: []string{"this", "is", "test", "case"},
	valid:  []bool{true, false},
}

func TestDoPrint(t *testing.T) {
	for _, v := range testCases.valid {
		conf := config.Config{config.FLAG_BEAUTY: v}
		err := DoPrint(conf, testCases.Result)
		if err != nil {
			t.Error(err)
		}
	}
}
