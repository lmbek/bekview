package bekview_test

import (
	"bekview"
	"testing"
)

func BenchmarkStart(benchmark *testing.B) {
	for i := 0; i < benchmark.N; i++ {
		_ = bekview.Start()
	}
}
