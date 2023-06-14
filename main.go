package main

import (
	"fmt"
	"src/src"
	"time"
)

/*
Apertura - Cierre Estaciones
Lunes a Viernes: 06:00 - 23:30
Sabados: 07:30 - 23:30
Domingos y Feriados: 7:30 - 23:30
*/
/*
Salida Trenes Puerto - Limache
Lunes a Viernes: 06:15 - 22:30
Sabados: 	08:30 - 22:24
Domingos y Festivos: 09:00 - 22:12
*/
/*
Salida Trenes Limache - Puerto
Lunes a Viernes: 06:15 - 22:15
Sabados: 07:30 - 22:06
Domingos y Festivos: 08:00 - 22:06
*/
/*
Frecuencias
Lunes a Viernes: Trenes desde Puerto a Sargento Aldea cada 6 minutos entre las 07:36:00 - 09:36:00, y 16:06:00 - 19:54:00
Lunes a Viernes: Trenes desde Sargento Aldea a Puerto cada 6 minutos entre las 06:49:00 - 09:00:00, y 15:36:00 - 19:00:00
Lunes a Viernes: Hora Alta y Media: 12 minutos
Lunes a Viernes: Hora Baja: 15 minutos
Lunes a Viernes: Primer Tren desde El Belloto a las 6:22:00 en ambos sentidos
Sabados: 12 minutos
Domingos y Festivos: 15 minutos
*/

/*
Hora Baja:	10:30:00 - 12:59:59, 14:00:00 – 15:59:59, 21:00:00 – 23:59:59
Hora Media:	06:00:00 – 06:29:59, 09:30:00 – 10:29:59, 13:00:00 – 13:59:59, 16:00:00 – 16:59:59, 20:00:00 – 20:59:59
Hora Alta:	06:30:00 – 09:29:59, 17:00.00 – 19:59:59
sabados y domingos es hora baja todo el dia
*/

/*
Tarifa:
Mismo tramo: 	Alta:	0430	Media:	0409	Baja:	0388
Tramo 1-2:		Alta:	0590	Media: 	0566	Baja:	0534
Tramo 1-3:		Alta:	0854	Media:	0811	Baja:	0769
Tramo 1-4		Alta:	0906	Media:	0860	Baja:	0815
Tramo 1-5		Alta:	1126	Media:	1069	Baja:	1012
*/

/*
Descuentos: Estudiante: 66% - 3ra Edad: 50%
*/

type station struct {
	name    string
	minutes int
	sector  int
}
type schedule struct {
	day        time.Weekday
	from       int
	firstTrain time.Time
	lastTrain  time.Time
}
type passenger struct {
	card card
	fee  fee
}
type card struct {
	name     string
	discount int
}
type fee struct {
	name  string
	value [5]float64
}

var horaBaja = fee{
	"Hora Baja",
	[5]float64{388.0, 534.0, 769.0, 815.0, 1012.0},
}

var horaMedia = fee{
	"Hora Media",
	[5]float64{409.0, 566.0, 811.0, 860.0, 1069.0},
}

var horaAlta = fee{
	"Hora Alta",
	[5]float64{430.0, 590.0, 854.0, 906.0, 1126.0},
}

var cards = [3]card{
	card{
		name:     "General",
		discount: 100,
	}, card{
		name:     "Estudiante",
		discount: 66,
	}, card{
		name:     "Adulto Mayor",
		discount: 50,
	},
}

var stations = [20]station{
	station{
		name:    "Puerto",
		minutes: 0,
		sector:  1,
	},
	station{
		name:    "Bellavista",
		minutes: 1,
		sector:  1,
	},
	station{
		name:    "Francia",
		minutes: 2,
		sector:  1,
	},
	station{
		name:    "Baron",
		minutes: 5,
		sector:  1,
	},
	station{
		name:    "Portales",
		minutes: 8,
		sector:  1,
	},
	station{
		name:    "Recreo",
		minutes: 11,
		sector:  2,
	},
	station{
		name:    "Miramar",
		minutes: 13,
		sector:  2,
	},
	station{
		name:    "Viña del Mar",
		minutes: 14,
		sector:  2,
	},
	station{
		name:    "Hospital",
		minutes: 16,
		sector:  2,
	},
	station{
		name:    "Chorrillos",
		minutes: 17,
		sector:  2,
	},
	station{
		name:    "El Salto",
		minutes: 19,
		sector:  2,
	},
	station{
		name:    "Quilpue",
		minutes: 40,
		sector:  3,
	},
	station{
		name:    "El Sol",
		minutes: 42,
		sector:  3,
	},
	station{
		name:    "El Belloto",
		minutes: 45,
		sector:  3,
	},
	station{
		name:    "Las Americas",
		minutes: 47,
		sector:  4,
	},
	station{
		name:    "La Concepcion",
		minutes: 49,
		sector:  4,
	},
	station{
		name:    "Villa Alemana",
		minutes: 51,
		sector:  4,
	},
	station{
		name:    "Sargento Aldea",
		minutes: 53,
		sector:  4,
	},
	station{
		name:    "Peñablanca",
		minutes: 56,
		sector:  4,
	},
	station{
		name:    "Limache",
		minutes: 69,
		sector:  5,
	},
}

func getFeeTime(t time.Time) fee {
	var f fee
	switch t.Weekday() {
	case 6, 0:
		f = horaBaja
	case 1, 2, 3, 4, 5:
		if src.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 10, 30, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 12, 59, 59, 0, t.Location())) ||
			src.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 14, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 15, 59, 59, 0, t.Location())) ||
			src.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 21, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())) {
			f = horaBaja
		} else if src.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 6, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 6, 29, 59, 0, t.Location())) ||
			src.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 9, 30, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 10, 29, 59, 0, t.Location())) ||
			src.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 13, 00, 00, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 13, 59, 59, 0, t.Location())) ||
			src.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 16, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 16, 59, 59, 0, t.Location())) ||
			src.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 20, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 20, 59, 59, 0, t.Location())) {
			f = horaMedia
		} else if src.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 6, 30, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 9, 29, 59, 0, t.Location())) ||
			src.IsTimeBetweenDates(t, time.Date(t.Year(), t.Month(), t.Day(), 17, 00, 0, 0, t.Location()), time.Date(t.Year(), t.Month(), t.Day(), 19, 59, 59, 0, t.Location())) {
			f = horaAlta
		}
	}
	return f
}

func getFeePrice(f fee, originIndex, destinationIndex int) int {
	originSector, destinationSector := stations[originIndex].sector, stations[destinationIndex].sector
	return int(f.value[src.Abs(destinationSector-originSector)])
}

func newSchedule(day time.Weekday, from int, firstTrain time.Time, lastTrain time.Time) schedule {
	return schedule{day, from, firstTrain, lastTrain}
}

func getDaySchedule(t time.Time, direction int) schedule {
	var daySchedule schedule
	weekday := t.Weekday()
	if int(weekday) > 6 {
		weekday = src.DayOverflow(weekday)
	}
	switch direction {
	case 0: // Desde Puerto
		switch weekday {
		case 0:
			daySchedule = schedule{
				day:        weekday,
				from:       0,
				firstTrain: time.Date(t.Year(), t.Month(), t.Day(), 9, 0, 0, 0, t.Location()),
				lastTrain:  time.Date(t.Year(), t.Month(), t.Day(), 22, 12, 0, 0, t.Location()),
			}
		case 6:
			daySchedule = schedule{
				day:        weekday,
				from:       0,
				firstTrain: time.Date(t.Year(), t.Month(), t.Day(), 8, 30, 0, 0, t.Location()),
				lastTrain:  time.Date(t.Year(), t.Month(), t.Day(), 22, 24, 0, 0, t.Location()),
			}
		case 1, 2, 3, 4, 5:
			daySchedule = schedule{
				day:        weekday,
				from:       0,
				firstTrain: time.Date(t.Year(), t.Month(), t.Day(), 6, 15, 0, 0, t.Location()),
				lastTrain:  time.Date(t.Year(), t.Month(), t.Day(), 22, 30, 0, 0, t.Location()),
			}
		}
	case 19: // Desde Limache
		switch weekday {
		case 0:
			daySchedule = schedule{
				day:        weekday,
				from:       19,
				firstTrain: time.Date(t.Year(), t.Month(), t.Day(), 8, 0, 0, 0, t.Location()),
				lastTrain:  time.Date(t.Year(), t.Month(), t.Day(), 22, 6, 0, 0, t.Location()),
			}
		case 6:
			daySchedule = schedule{
				day:        weekday,
				from:       19,
				firstTrain: time.Date(t.Year(), t.Month(), t.Day(), 7, 30, 0, 0, t.Location()),
				lastTrain:  time.Date(t.Year(), t.Month(), t.Day(), 22, 26, 0, 0, t.Location()),
			}
		case 1, 2, 3, 4, 5:

			daySchedule = schedule{
				day:        weekday,
				from:       19,
				firstTrain: time.Date(t.Year(), t.Month(), t.Day(), 6, 15, 0, 0, t.Location()),
				lastTrain:  time.Date(t.Year(), t.Month(), t.Day(), 22, 15, 0, 0, t.Location()),
			}
		}
	}
	return daySchedule
}

func getDayFromDate(t time.Time) time.Weekday {
	weekday := t.Weekday()
	return weekday
}

func getTimeTravelDuration(originIndex, destinationIndex int) time.Duration {
	minutes := src.Abs(stations[destinationIndex].minutes - stations[originIndex].minutes)
	return time.Duration(time.Duration(minutes) * time.Minute)
}

func getTrainDirection(originIndex, destinationIndex int) int {
	var direction int
	if originIndex < destinationIndex {
		direction = 0 // Tren parte de Estacion Puerto
	} else if originIndex > destinationIndex {
		direction = 19 // Tren parte de Estacion Limache
	}
	return direction
}

func getStationTrainSchedule(originIndex, destinationIndex, userType int) {
	currentTime := time.Now()
	direction := getTrainDirection(originIndex, destinationIndex)
	trainSchedule := getDaySchedule(currentTime, direction)
	firstTrain := trainSchedule.firstTrain
	lastTrain := trainSchedule.lastTrain
	passenger := cards[userType]
	currentTrain := firstTrain
	travelTime := time.Duration(src.Abs(stations[originIndex].minutes-stations[destinationIndex].minutes)) * time.Minute
	fmt.Printf("Train from %v to %v\n", stations[originIndex].name, stations[destinationIndex].name)
	for lastTrain.After(currentTrain) {
		trainFee := getFeeTime(currentTrain)
		feePrice := getFeePrice(trainFee, originIndex, destinationIndex) * passenger.discount / 100
		departure := currentTrain.Format(time.Kitchen)
		arrival := currentTrain.Add(time.Minute + travelTime).Format(time.Kitchen)
		fmt.Printf("Departure %v - Arrival: %v - Fee: %v - Price: %v\n", departure, arrival, trainFee.name, feePrice)
		currentTrain = currentTrain.Add(time.Minute * 12)
	}
}

func drawTable() {

}

func main() {
	//getStationTrainSchedule(19, 0)
	var origin, destination, user int
	fmt.Println("Tramo 1: 1: Puerto - 2: Bellavista - 3: Francia - 3: Baron - 4: Portales")
	fmt.Println("Tramo 2: 5: Recreo - 6 Miramar - 7: Viña del Mar - 8: Hospital - 9: Chorrillos - 10: El Salto")
	fmt.Println("Tramo 3: 11: Quilpue - 12: El Sol - 13: El Belloto")
	fmt.Println("Tramo 4: 14: Las Americas - 15: La Concepcion - 16: Villa Alemana - 17: Sargento Aldea - 18: Peñablanca")
	fmt.Println("Tramo 5: 19: Limache")
	fmt.Println("Tipos de usuario: 0: General (sin %) - 1: Estudiante (66%) - 2: Tercera Edad (50%)")
	fmt.Println("Porfavor ingrese los datos separados por un espacio y en el siguiente orden: ")
	fmt.Println("  Estacion Origen  Estacion Destino  Tipo de Usuario  ")
	_, err := fmt.Scanf("%v %v %v", &origin, &destination, &user)
	if err != nil {
		fmt.Printf("Your input could not be read%v", err)
	}
	getStationTrainSchedule(origin, destination, user)
}
