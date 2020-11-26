package sql

import (
	"log"
	"fmt"
)

type clienteObj struct {
	nrocliente int64
	nombre string
	apellido string
	domicilio string
	telefono int64
}

type comercioObj struct {
	nrocomercio int
	nombre string
	domicilio string
	codigopostal string
	telefono int64
}

type tarjetaObj struct {
	nrotarjeta string
	nrocliente int
	validadesde string
	validahasta string
	codseguridad int
	limitecompra float64
	estado string
}

func MostrarTodosDatos() {
	fmt.Println("---------------------Datos Clientes:--------------------- \n")
	mostrarDatosClientes()
	fmt.Println("\n \n")
	fmt.Println("---------------------Datos Comercios:--------------------- \n")
	mostrarDatosComercios()
	fmt.Println("\n \n")
	fmt.Println("---------------------Datos Tarjetas:--------------------- \n")
	mostrarDatosTarjetas()
}

func mostrarDatosClientes(){
	row, err := db.Query(`SELECT * FROM cliente`)
	if err != nil {
			log.Fatal(err)
	}
	defer row.Close()
	
	
	for row.Next() {
		var c1 clienteObj
		
		if err := row.Scan(&c1.nrocliente, &c1.nombre, &c1.apellido, &c1.domicilio, &c1.telefono); err != nil {
			log.Fatal(err)
		}
		
		fmt.Println(c1)
	}
}

func mostrarDatosComercios(){
	row, err := db.Query(`SELECT * FROM comercio`)
	if err != nil {
			log.Fatal(err)
	}
	defer row.Close()
	
	
	for row.Next() {
		var c2 comercioObj
	
		if err := row.Scan(&c2.nrocomercio, &c2.nombre, &c2.domicilio, &c2.codigopostal, &c2.telefono); err != nil {
			log.Fatal(err)
		}

		fmt.Println(c2)
	}
}

func mostrarDatosTarjetas(){
	row, err := db.Query(`SELECT * FROM tarjeta`)
	if err != nil {
			log.Fatal(err)
	}
	defer row.Close()
	
	
	for row.Next() {
		var t1 tarjetaObj
	
		if err := row.Scan(&t1.nrotarjeta, &t1.nrocliente, &t1.validadesde, &t1.validahasta, &t1.codseguridad, &t1.limitecompra, &t1.estado); 
		err != nil {
			log.Fatal(err)
		}

		fmt.Println(t1)
	}
}
