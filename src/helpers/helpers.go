package helpers

import "time"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func DayOverflow(w time.Weekday) time.Weekday {
	if w > 6 {
		w -= 7
		return DayOverflow(w)
	}
	return w
}

func IsTimeBetweenDates(t, min, max time.Time) bool {
	if min.After(max) {
		min, max = max, min
	}
	return (t.Equal(min) || t.After(min)) && (t.Equal(max) || t.Before(max))
}

func GetDayFromDateTime(t time.Time) time.Weekday {
	weekday := t.Weekday()
	return weekday
}
