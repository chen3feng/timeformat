package timeformat

// CFormat represent C language date time format, such as strftime.
type CFormat struct {
	baseFormat
}

// C returns a C language time format.
func C(format string) Format {
	f := CFormat{}
	f.init(format)
	return &f
}

func (f *CFormat) init(format string) {
	f.format = format
	start := 0
	for i := 0; i < len(format); i++ {
		c := format[i]
		if c == '%' && i+1 < len(format) {
			if start < i {
				s := format[start:i]
				f.appendStringFormat(s)
			}
			i++
			switch format[i] {
			case 'Y':
				f.fields = append(f.fields, fieldSpec{fieldKind_Year4, "year"})
			case 'y':
				f.fields = append(f.fields, fieldSpec{fieldKind_Year2, "year2"})
			case 'm':
				f.fields = append(f.fields, fieldSpec{fieldKind_Month2, "month2"})
			case 'b':
				f.fields = append(f.fields, fieldSpec{fieldKind_Month3, "month3"})
			case 'B':
				f.fields = append(f.fields, fieldSpec{fieldKind_Month, "month"})
			case 'd':
				f.fields = append(f.fields, fieldSpec{fieldKind_Day2, "day"})
			case 'H':
				f.fields = append(f.fields, fieldSpec{fieldKind_Hour24, "hour"})
			case 'h':
				f.fields = append(f.fields, fieldSpec{fieldKind_Hour12, "hour12"})
			case 'M':
				f.fields = append(f.fields, fieldSpec{fieldKind_Minute2, "minute2"})
			case 'S':
				f.fields = append(f.fields, fieldSpec{fieldKind_Second2, "second2"})
			case 'a':
				f.fields = append(f.fields, fieldSpec{fieldKind_Weekday3, "weekday"})
			case 'A':
				f.fields = append(f.fields, fieldSpec{fieldKind_Weekday, "weekday"})
			}
			start = i + 1
		}
	}
	if start < len(format) {
		s := format[start:]
		f.appendStringFormat(s)
	}
}

func (f *CFormat) appendStringFormat(s string) {
	f.fields = append(f.fields, fieldSpec{fieldKind_Literal, s})
}
