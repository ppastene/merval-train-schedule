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

func getStationTrainSchedule(origin, destination src.Station, userType int) {
	currentTime := time.Now()
	direction := origin.GetTravelDirection(destination)
	trainSchedule := src.GetDaySchedule(currentTime, direction)
	firstTrain, lastTrain := trainSchedule.FirstTrain, trainSchedule.LastTrain
	passenger := src.Cards[userType]
	currentTrain := firstTrain
	travelTime := time.Duration(src.Abs(origin.Minutes-destination.Minutes)) * time.Minute
	fmt.Printf("Train from %v to %v\n", origin.Name, destination.Name)
	for lastTrain.After(currentTrain) {
		trainFee := src.GetFeeFromTime(currentTrain)
		feePrice := trainFee.GetFeePrice(origin, destination) * passenger.Discount / 100 // REFACTOR
		departure := currentTrain.Format(time.Kitchen)                                   // REFACTOR
		arrival := currentTrain.Add(time.Minute + travelTime).Format(time.Kitchen)       // REFACTOR
		fmt.Printf("Departure %v - Arrival: %v - Fee: %v - Price: %v\n", departure, arrival, trainFee.Name, feePrice)
		currentTrain = currentTrain.Add(time.Minute * 12)
	}
}

func drawTable() {

}

func main() {
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
	//time := time.Date(2023, 7, 14, 8, 30, 0, 0, time.Local)
	originStation, originError := src.GetStation(origin)
	destinationStation, destinationError := src.GetStation(destination)
	if originError != nil {
		fmt.Println("An error happened:", originError)
	} else if destinationError != nil {
		fmt.Println("An error happened:", destinationError)
		return
	}
	getStationTrainSchedule(originStation, destinationStation, 0)
}
