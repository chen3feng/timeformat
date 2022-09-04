package timeformat

import (
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

type Format interface {
	Print(time.Time) string
	Parse(value string) (time.Time, error)
}

type brokenTime struct {
	year                 int
	month                time.Month
	day                  int
	hour, minute, second int
	nsec                 int
	loc                  *time.Location
}

type fieldKind int

const (
	fieldKind_Str fieldKind = iota
	fieldKind_Year4
	fieldKind_Month2
	fieldKind_Day
	fieldKind_Hour
	fieldKind_Minute
	fieldKind_Second
)

type fieldSpec struct {
	kind fieldKind
	name string
}

type baseFormat struct {
	format string
	fields []fieldSpec
}

func (f *baseFormat) Print(t time.Time) string {
	const bufSize = 64
	var bs []byte
	max := len(f.format) + 10
	if max < bufSize {
		var buf [bufSize]byte
		bs = buf[:0]
	} else {
		bs = make([]byte, 0, max)
	}
	bs = f.Append(bs, t)
	return string(bs)
}

func (f *baseFormat) Append(bs []byte, t time.Time) []byte {
	for _, field := range f.fields {
		switch field.kind {
		case fieldKind_Str:
			bs = append(bs, field.name...)
		case fieldKind_Year4:
			bs = appendYear(bs, t)
		case fieldKind_Month2:
			bs = appendMonth2(bs, t)
		case fieldKind_Day:
			bs = appendMonthDay(bs, t)
		case fieldKind_Hour:
			bs = appendHour(bs, t)
		case fieldKind_Minute:
			bs = appendMinute(bs, t)
		case fieldKind_Second:
			bs = appendSecond(bs, t)
		}
	}
	return bs
}

//go:linkname noescape runtime.noescape
//go:noescape
func noescape(p unsafe.Pointer) unsafe.Pointer

// //go:nosplit
// func noescape(p unsafe.Pointer) unsafe.Pointer {
// 	x := uintptr(p)
// 	return unsafe.Pointer(x ^ 0)
// }

//go:nosplit
func noEscape[T any](p *T) *T {
	return (*T)(noescape(unsafe.Pointer(p)))
}

func (f baseFormat) Parse(value string) (t time.Time, err error) {
	var bt brokenTime
	bt.loc = time.Local
	err = f.parse(value, noEscape(&bt))
	return time.Date(bt.year, bt.month, bt.day, bt.hour, bt.minute, bt.second, bt.nsec, bt.loc), nil
}

func (f baseFormat) parse(value string, bt *brokenTime) error {
	// for _, parser := range f.parsers {
	// 	var err error
	// 	value, err = parser(value, bt)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	var err error
	for _, field := range f.fields {
		switch field.kind {
		case fieldKind_Str:
			value, err = parseString(field.name, value)
		case fieldKind_Year4:
			value, err = parseYear(value, bt)
		case fieldKind_Month2:
			value, err = parseMonth2(value, bt)
		case fieldKind_Day:
			value, err = parseMonthDay(value, bt)
		case fieldKind_Hour:
			value, err = parseHour(value, bt)
		case fieldKind_Minute:
			value, err = parseMinute(value, bt)
		case fieldKind_Second:
			value, err = parseSecond(value, bt)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// small number to string
var numBytes [100][2]byte

func init() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			numBytes[i*10+j][0] = byte(i + '0')
			numBytes[i*10+j][1] = byte(j + '0')
		}
	}
}

func parseString(expected, str string) (string, error) {
	if str[:len(expected)] != expected {
		return str, fmt.Errorf("mismatch, expect %v, got %v", expected, str)
	}
	return str[len(expected):], nil
}

func appendSignedInt(bs []byte, n int, width int) []byte {
	u := uint(n)
	if n < 0 {
		bs = append(bs, '-')
		u = uint(-n)
	}

	// Assemble decimal in reverse order.
	var buf [20]byte
	i := len(buf)
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)

	// Add 0-padding.
	for w := len(buf) - i; w < width; w++ {
		bs = append(bs, '0')
	}

	return append(bs, buf[i:]...)
}

func appendInt(bs []byte, n int, width int) []byte {
	u := uint(n)

	// Assemble decimal in reverse order.
	var buf [8]byte
	i := len(buf)
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)

	// Add 0-padding.
	for w := len(buf) - i; w < width; w++ {
		bs = append(bs, '0')
	}

	return append(bs, buf[i:]...)
}

func appendSignedInt4(bs []byte, n int) []byte {
	return appendSignedInt(bs, n, 4)
}

func appendInt2(bs []byte, n int) []byte {
	ns := numBytes[n]
	return append(bs, ns[0], ns[1])
}

func parseYear(str string, bt *brokenTime) (string, error) {
	n, err := strconv.ParseInt(str[:4], 10, 16)
	bt.year = int(n)
	return str[4:], err
}

func appendYear(bs []byte, t time.Time) []byte {
	return appendSignedInt4(bs, t.Year())
}

func parseYear2(str string, bt *brokenTime) (string, error) {
	return str, nil
}

func parseMonth(str string, bt *brokenTime) (string, error) {
	return str, nil
}

func parseMonth3(str string, bt *brokenTime) (string, error) {
	return str, nil
}

func parseString2(field, str string, value *int) (string, error) {
	if len(str) < 2 {
		return str, fmt.Errorf("invalid %v %v", field, str)
	}
	*value = int((str[0]-'0')*10 + str[1] - '0')
	return str[2:], nil
}

func parseMonth2(str string, bt *brokenTime) (tail string, err error) {
	month := 0
	tail, err = parseString2("month", str, &month)
	bt.month = time.Month(month)
	return
}

func appendMonth2(bs []byte, t time.Time) []byte {
	return appendInt2(bs, int(t.Month()))
}

func parseMonthDay(str string, bt *brokenTime) (string, error) {
	return parseString2("day", str, &bt.day)
}

func appendMonthDay(bs []byte, t time.Time) []byte {
	return appendInt2(bs, int(t.Day()))
}

func parseYearDay(str string, bt *brokenTime) (string, error) {
	return str, nil
}

func parseWeekDay(str string, bt *brokenTime) (string, error) {
	return str, nil
}

func parseWeekDay3(str string, bt *brokenTime) (string, error) {
	return str, nil
}

func parseWeekDay1(str string, bt *brokenTime) (string, error) {
	return str, nil
}

func parseHour(str string, bt *brokenTime) (string, error) {
	return parseString2("hour", str, &bt.hour)
}

func appendHour(bs []byte, t time.Time) []byte {
	return appendInt2(bs, t.Hour())
}

func parseMinute(str string, bt *brokenTime) (string, error) {
	return parseString2("minute", str, &bt.minute)
}

func appendMinute(bs []byte, t time.Time) []byte {
	return appendInt2(bs, t.Minute())
}

func parseSecond(str string, bt *brokenTime) (string, error) {
	return parseString2("second", str, &bt.second)
}

func appendSecond(bs []byte, t time.Time) []byte {
	return appendInt2(bs, t.Second())
}
