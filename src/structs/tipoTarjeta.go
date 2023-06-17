/*
------------------------------
--------- DEPRECATED ---------
------------------------------
*/
package structs

type TipoTarjeta struct {
	Nombre    string
	Descuento int
}

func NewTipoTarjeta(n string, d int) TipoTarjeta {
	return TipoTarjeta{n, d}
}

var general TipoTarjeta = NewTipoTarjeta("General", 100)
var estudiante TipoTarjeta = NewTipoTarjeta("Estudiante", 66)
var terceraEdad TipoTarjeta = NewTipoTarjeta("Tercera Edad", 50)
