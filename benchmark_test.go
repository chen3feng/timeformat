package timeformat_test

import (
	"testing"
	"time"

	"github.com/chen3feng/timeformat"
)

func BenchmarkPrint(b *testing.B) {
	f := timeformat.C("%Y-%m-%d %H:%M:%S")
	b.Run("C", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f.Print(testTime)
		}
	})
	b.Run("Go", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = testTime.Format("2006-01-02 15:04:05")
		}
	})
}

func BenchmarkParse(b *testing.B) {
	f := timeformat.C("%Y-%m-%d %H:%M:%S")
	b.Run("C", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f.Parse("2006-01-02 15:04:05")
		}
	})
	b.Run("Go", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
		}
	})
}
