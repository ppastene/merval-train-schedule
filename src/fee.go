package src

import "time"

type fee struct {
	Name  string
	Value [5]float64
}

var Fees = [3]fee{
	{"Hora Baja", [5]float64{388.0, 534.0, 769.0, 815.0, 1012.0}},
	{"Hora Media", [5]float64{409.0, 566.0, 811.0, 860.0, 1069.0}},
	{"Hora Media", [5]float64{409.0, 566.0, 811.0, 860.0, 1069.0}},
}

func newFee(Name string, Value [5]float64) fee {
	return fee{Name, Value}
}

func GetFeeFromTime(t time.Time) fee {
	var f fee
	switch t.Weekday() {
	case 6, 0:
		f = Fees[0]
	case 1, 2, 3, 4, 5:
		if IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 10, 30, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 12, 59, 59, 0, t.Location())) ||
			IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 14, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 15, 59, 59, 0, t.Location())) ||
			IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 21, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())) {
			f = Fees[0]
		} else if IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 6, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 6, 29, 59, 0, t.Location())) ||
			IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 9, 30, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 10, 29, 59, 0, t.Location())) ||
			IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 13, 00, 00, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 13, 59, 59, 0, t.Location())) ||
			IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 16, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 16, 59, 59, 0, t.Location())) ||
			IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 20, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 20, 59, 59, 0, t.Location())) {
			f = Fees[1]
		} else if IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 6, 30, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 9, 29, 59, 0, t.Location())) ||
			IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 17, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 19, 59, 59, 0, t.Location())) {
			f = Fees[2]
		}
	}
	return f
}

func (f fee) GetFeePrice(origin, destination Station) int {
	return int(f.Value[Abs(origin.Sector-destination.Sector)])
}
