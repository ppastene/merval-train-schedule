package src

import "time"

type schedule struct {
	Day        time.Weekday
	From       int
	FirstTrain time.Time
	LastTrain  time.Time
}

func newSchedule(day time.Weekday, from int, firstTrain time.Time, lastTrain time.Time) schedule {
	return schedule{day, from, firstTrain, lastTrain}
}

func GetDaySchedule(t time.Time, direction int) schedule {
	var daySchedule schedule
	weekday := t.Weekday()
	if int(weekday) > 6 {
		weekday = DayOverflow(weekday)
	}
	if direction == 0 {
		switch weekday {
		case 0:
			daySchedule = schedule{
				Day:        weekday,
				From:       0,
				FirstTrain: time.Date(t.Year(), t.Month(), t.Day(), 9, 0, 0, 0, t.Location()),
				LastTrain:  time.Date(t.Year(), t.Month(), t.Day(), 22, 12, 0, 0, t.Location()),
			}
		case 6:
			daySchedule = schedule{
				Day:        weekday,
				From:       0,
				FirstTrain: time.Date(t.Year(), t.Month(), t.Day(), 8, 30, 0, 0, t.Location()),
				LastTrain:  time.Date(t.Year(), t.Month(), t.Day(), 22, 24, 0, 0, t.Location()),
			}
		case 1, 2, 3, 4, 5:
			daySchedule = schedule{
				Day:        weekday,
				From:       0,
				FirstTrain: time.Date(t.Year(), t.Month(), t.Day(), 6, 15, 0, 0, t.Location()),
				LastTrain:  time.Date(t.Year(), t.Month(), t.Day(), 22, 30, 0, 0, t.Location()),
			}
		}
	} else if direction == 19 {
		switch weekday {
		case 0:
			daySchedule = schedule{
				Day:        weekday,
				From:       19,
				FirstTrain: time.Date(t.Year(), t.Month(), t.Day(), 8, 0, 0, 0, t.Location()),
				LastTrain:  time.Date(t.Year(), t.Month(), t.Day(), 22, 6, 0, 0, t.Location()),
			}
		case 6:
			daySchedule = schedule{
				Day:        weekday,
				From:       19,
				FirstTrain: time.Date(t.Year(), t.Month(), t.Day(), 7, 30, 0, 0, t.Location()),
				LastTrain:  time.Date(t.Year(), t.Month(), t.Day(), 22, 26, 0, 0, t.Location()),
			}
		case 1, 2, 3, 4, 5:

			daySchedule = schedule{
				Day:        weekday,
				From:       19,
				FirstTrain: time.Date(t.Year(), t.Month(), t.Day(), 6, 15, 0, 0, t.Location()),
				LastTrain:  time.Date(t.Year(), t.Month(), t.Day(), 22, 15, 0, 0, t.Location()),
			}
		}
	}
	return daySchedule
}
