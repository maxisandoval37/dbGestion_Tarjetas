package main

import (
	"fmt"
	"./sql"
)

func main() {
	sql.OpenDbConnection()
	menuPrincipal()
}

var opcion int=0

func menuPrincipal() {

		for opcion != 99 {
			fmt.Println("Digite una opcion")
			fmt.Println("1: Crear BD SQL")
			fmt.Println("2: Crear Tablas")
			

			fmt.Println("99: Salir ")

			fmt.Scanf("%d", &opcion)
		
		switch {
			case opcion == 1:
				sql.CrearDB()
				fmt.Println("--------------------\n    ¡BD CREADA!\n--------------------")
			case opcion == 2:
				sql.CrearTablas()
				fmt.Println("------------------------\n    ¡TABLAS CREADAS!\n------------------------")
			
			
			default:
				fmt.Println("Ingrese una opcion valida")
			}
	}
}

