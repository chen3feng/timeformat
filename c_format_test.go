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
	expectEq(t, timeformat.C("%Y %b %d %H:%M:%S").Print(testTime), "2016 Jan 02 15:04:05")
	expectEq(t, timeformat.C("%Y %B %d %H:%M:%S").Print(testTime), "2016 January 02 15:04:05")
}

func TestTimeParse(t *testing.T) {
	tm1 := time.Date(2021, 9, 3, 0, 0, 0, 0, time.Local)
	test := func(format, str string) {
		t.Helper()
		f := timeformat.C(format)
		tm, err := f.Parse(str)
		expectNoError(t, err)
		expectTrue(t, tm1.Equal(tm))
	}
	test("%Y-%m-%d %H:%M:%S", "2021-09-03 00:00:00")
	test("%Y-%b-%d %H:%M:%S", "2021-Sep-03 00:00:00")
	test("%Y-%B-%d %H:%M:%S", "2021-September-03 00:00:00")
}
