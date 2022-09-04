package timeformat

import (
	"time"
)

// Go represent C language date time format, such as strftime.
type GoFormat struct {
	layout string
}

// Go returns a go time format.
func Go(layout string) Format {
	return &GoFormat{layout}
}

func (f GoFormat) Print(t time.Time) string {
	return t.Format(f.layout)
}

func (f GoFormat) Parse(value string) (time.Time, error) {
	return time.Parse(f.layout, value)
}
