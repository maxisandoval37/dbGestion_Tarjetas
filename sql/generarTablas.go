package sql

import (
	"fmt"
	"log"
)

var cliente=" cliente (nrocliente int,nombre text,apellido text,domicilio text,telefono char(12)); "
var tarjeta=" tarjeta (nrotarjeta char(16),nrocliente int,validadesde char(6),validahasta char(6),codseguridad char(4),limitecompra decimal(8,2),estado char(10)); "
var comercio=" comercio (nrocomercio int,nombre text,domicilio text,codigopostal char(8),telefono char(12)); "
var compra=" compra (nrooperacion serial,nrotarjeta char(16),nrocomercio int,fecha timestamp,monto decimal(7,2),pagado bool); "
var rechazo=" rechazo (nrorechazo serial,nrotarjeta char(16),nrocomercio int,fecha timestamp,monto decimal(7,2),motivo text); "
var cierre=" cierre (año int,mes int,terminacion int,fechainicio date,fechacierre date,fechavto date); "
var cabecera=" cabecera(nroresumen serial,nombre text,apellido text,domicilio text,nrotarjeta char(16),desde date,hasta date,vence date,total decimal(8,2)); "
var detalle=" detalle(nroresumen int,nrolinea int,fecha date,nombrecomercio text,monto decimal(7,2)); "
var alerta=" alerta (nroalerta serial,nrotarjeta char(16),fecha timestamp,nrorechazo int,codalerta int,descripcion text); "
var consum=" consumo (nrotarjeta char(16),codseguridad char(4),nrocomercio int,monto decimal(7,2))"


func CrearTablas() {
	
	auxResetearTablas()

	_, err = db.Exec(`CREATE TABLE` + cliente + `
											
					CREATE TABLE` + tarjeta + `
											
					CREATE TABLE` + comercio + `
											
					CREATE TABLE` + compra + `
											
					CREATE TABLE` + rechazo + `
											
					CREATE TABLE` + cierre + `
											
					CREATE TABLE` + cabecera + `
											
					CREATE TABLE` + detalle + `
											
					CREATE TABLE` + alerta + `
											
					CREATE TABLE` + consum)			// Esta tabla no es parte del modelo de datos, pero se incluye para
													// poder probar las funciones.
	if err != nil {
		log.Fatal(err)
	}
}

func auxResetearTablas(){//limpia las tablas para evitar error cuando creamos encima de las ya creadas
	_, err = db.Exec(`DROP SCHEMA public CASCADE`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE SCHEMA public`)
	if err != nil {
		log.Fatal(err)
	}
}

func MostrarTablas() {
	var todasTablas =cliente+"\n"+"\n"+tarjeta+"\n"+"\n"+comercio+"\n"+"\n"+compra+"\n"+"\n"+rechazo+"\n"+"\n"+cierre+"\n"+"\n"+cabecera+"\n"+"\n"+detalle+"\n"+"\n"+alerta+"\n"+"\n"+consum
	fmt.Println(todasTablas)
}
