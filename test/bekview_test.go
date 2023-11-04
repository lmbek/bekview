package bekview_test

import (
	"bekview"
	"testing"
)

func TestStart(test *testing.T) {
	err := bekview.Start()
	if err != nil {
		test.Errorf(err.Error())
	}
}
