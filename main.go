package main

import (
	"fmt"
	"./sql"
	"./nosql"
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
			fmt.Println("3: Mostrar Tablas de la BD")
			fmt.Println("4: Generar PKs y FKs")
			fmt.Println("5: Borrar PKs y FKs")
			fmt.Println("6: Cargar datos")
			fmt.Println("7: Mostrar Todos Datos Cargados")
			
			fmt.Println("9: CargarDatosNoSQL en BoltDB")

			fmt.Println("99: Salir ")

			fmt.Scanf("%d", &opcion)
		
		switch {
			case opcion == 1:
				sql.CrearDB()
				fmt.Println("--------------------\n    ¡BD CREADA!\n--------------------")
			case opcion == 2:
				sql.CrearTablas()
				fmt.Println("------------------------\n    ¡TABLAS CREADAS!\n------------------------")
			case opcion == 3:
				sql.MostrarTablas()
				fmt.Println("---------------------------\n    ¡TABLAS MOSTRADAS!\n-------------------------")
			case opcion == 4:
				sql.GenerarPKs()
				sql.GenerarFKs()
				fmt.Println("---------------------------\n    ¡KEYS GENERADAS!\n-------------------------")
			case opcion == 5:
				sql.BorrarFKs()
				sql.BorrarPKs()
				fmt.Println("---------------------------\n    ¡KEYS BORRADAS!\n-------------------------")
			case opcion == 6:
				sql.CargarDatos()
				fmt.Println("---------------------------\n    ¡DATOS CARGADOS!\n-------------------------")
			case opcion == 7:
				sql.MostrarTodosDatos()
				fmt.Println("---------------------------\n    ¡DATOS MOSTRADOS!\n-------------------------")
			case opcion == 9:
				nosql.CargaDatosNoDB()
				fmt.Println("---------------------------\n    ¡DATOS NOSQL CARGADOS!\n-------------------------")
			default:
				fmt.Println("Ingrese una opcion valida")
			}
	}
}
