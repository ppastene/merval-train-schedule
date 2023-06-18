package structs

import (
	"time"
)

type TarjetaEspecial struct {
	Tarjeta
	Descuento  int
	Expiracion time.Time
}

func NewTarjetaEspecial(t Tarjeta, d int, e time.Time) *TarjetaEspecial {
	return &TarjetaEspecial{t, d, e}
}

/*
type Beneficio interface {
	esVigente() bool
	aplicarDescuento(valor int) int
}
*/

func (te TarjetaEspecial) esVigente() bool {
	if time.Now().Before(te.Expiracion) {
		return false
	}
	return true
}

func (te TarjetaEspecial) AplicarDescuento(valor int) int {
	vigente := te.esVigente()
	if !vigente {
		return valor
	}
	return valor * te.Descuento / 100
}

var Estudiante = NewTarjetaEspecial(NewTarjeta("Estudiante"), 66, time.Now().AddDate(0, -6, 0))
var TerceraEdad = NewTarjetaEspecial(NewTarjeta("Tercera Edad"), 50, time.Now().AddDate(0, -6, 0))
