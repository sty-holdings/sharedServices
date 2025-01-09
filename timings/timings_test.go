package sharedServices

import (
	"math/rand"
	"testing"
	"time"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
)

func TestRecordFunctionTimings(tPtr *testing.T) {

	var (
		testMode = true
	)

	tests := []struct {
		name string
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Test 0 ",
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Test 1 ",
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Test 2 ",
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Test 3 ",
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				fakeElapsedSeconds := float64(rand.Intn(100))
				RecordFunctionTimings(time.Duration(fakeElapsedSeconds), nil, testMode)
			},
		)
	}
}
