package timeformat_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/chen3feng/timeformat"
)

func Test_GoFormat_Print(t *testing.T) {
	f := timeformat.Go(time.RFC822Z)
	//tm1 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	tm2 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	s := f.Print(tm2)
	expectEq(t, s, "01 Jan 00 00:00 +0000")
}

func Test_GoFormat_Parse(t *testing.T) {
	//f := timeformat.Go(time.RFC1123Z)
	tm1 := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
	tm, err := time.ParseInLocation(time.RFC1123Z, "Mon, 02 Jan 2006 15:04:05 -0000", time.UTC)
	fmt.Println(tm.Format(time.RFC850))
	expectNoError(t, err)
	expectTrue(t, tm1.Equal(tm))
	name, offset := tm.Zone()
	_, _ = name, offset
}
