package src

import "time"

type schedule struct {
	day        time.Weekday
	from       int
	firstTrain time.Time
	lastTrain  time.Time
}

func newSchedule(day time.Weekday, from int, firstTrain time.Time, lastTrain time.Time) schedule {
	return schedule{day, from, firstTrain, lastTrain}
}
