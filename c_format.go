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
			case 'm':
				f.fields = append(f.fields, fieldSpec{fieldKind_Month2, "month"})
			case 'd':
				f.fields = append(f.fields, fieldSpec{fieldKind_Day, "day"})
			case 'H':
				f.fields = append(f.fields, fieldSpec{fieldKind_Hour, "hour"})
			case 'M':
				f.fields = append(f.fields, fieldSpec{fieldKind_Minute, "minute"})
			case 'S':
				f.fields = append(f.fields, fieldSpec{fieldKind_Second, "second"})
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
	f.fields = append(f.fields, fieldSpec{fieldKind_Str, s})
}
