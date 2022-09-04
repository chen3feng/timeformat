package timeformat_test

import (
	"testing"
	"time"

	"github.com/chen3feng/timeformat"
)

const (
	testTimeStr = "2006-01-02 15:04:05"
)

var (
	testTime = time.Date(2016, 1, 2, 15, 4, 5, 0, time.UTC)
)

func Test_CFormat_Print(t *testing.T) {
	expectEq(t, timeformat.C("%Y-%m-%d").Print(testTime), "2016-01-02")
	expectEq(t, timeformat.C("%Y-%m-%d %H:%M:%S").Print(testTime), "2016-01-02 15:04:05")
}

func TestTimeParse(t *testing.T) {
	f := timeformat.C("%Y-%m-%d %H:%M:%S")
	tm1 := time.Date(2021, 9, 3, 0, 0, 0, 0, time.Local)
	tm, err := f.Parse("2021-09-03 00:00:00")
	expectNoError(t, err)
	expectTrue(t, tm1.Equal(tm))
}
