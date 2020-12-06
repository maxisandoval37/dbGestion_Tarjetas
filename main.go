package main

import (
	"fmt"
	"./sql"
	nosql "./nosql"
)

func main() {
	sql.OpenDbConnection()
	menuPrincipal()
}

var opcion int = -1

func menuPrincipal() {
	
	for opcion != 0 {
		fmt.Println("-------------------------------------")
		fmt.Println("------- - DIGITE UNA OPCION - -------")
		fmt.Println("-------------------------------------")
		fmt.Println("-  0:  .    .    .    .    .  SALIR -")
		fmt.Println("-  1:   .    .    .    Crear BD SQL -")
		fmt.Println("-  2:    .    .    .   Crear Tablas -")
		fmt.Println("-  3:.    .    .  Generar PKs y FKs -")
		fmt.Println("-  4: .    .    .    .    . Alertas -")
		fmt.Println("-  5:  .    .Generar Cierres (2020) -")
		fmt.Println("-  6:   .    .    .    Cargar datos -")
		fmt.Println("-  7:  .    .    . Agregar Consumos -")
		fmt.Println("-  8: .    .    . Generar Resumenes -")
		fmt.Println("-------------------------------------")
		fmt.Println("-  9: .     Mostrar Tablas de la BD -")
		fmt.Println("- 10:  Mostrar Todos Datos Cargados -")	
		fmt.Println("- 11:    .    .    Borrar PKs y FKs -")		
		fmt.Println("- 12:.   CargarDatosNoSQL en BoltDB -")
		fmt.Println("-------------------------------------")		
		fmt.Scanf("%d", &opcion)
		
	switch {
		case opcion == 0:
			fmt.Println("-------------------------------------")
			fmt.Println("------------ ¡HA SALIDO! ------------")
			fmt.Println("-------------------------------------")
		case opcion == 1:
			sql.CrearDB()
			fmt.Println("-------------------------------------")
			fmt.Println("------------ ¡BD CREADA! ------------")
			fmt.Println("-------------------------------------")			
		case opcion == 2:
			sql.CrearTablas()
			fmt.Println("-------------------------------------")
			fmt.Println("--------- ¡TABLAS  CREADAS! ---------")
			fmt.Println("-------------------------------------")		
		case opcion == 3:
			sql.GenerarPKs()
			sql.GenerarFKs()
			fmt.Println("--------- ¡KEYS  GENERADAS! ---------")
			fmt.Println("-------------------------------------")		
		case opcion == 4:
			sql.Triggers()
			fmt.Println("-------------------------------------")
			fmt.Println("--------- ¡GENERAR ALERTAS! ---------")
			fmt.Println("-------------------------------------")
		case opcion == 5:
			sql.Cierres()
			fmt.Println("-------------------------------------")
			fmt.Println("-------- ¡CIERRES GENERADOS! --------")
			fmt.Println("-------------------------------------")
		case opcion == 6:
			sql.CargarDatos()
			fmt.Println("-------------------------------------")
			fmt.Println("--------- ¡DATOS  CARGADOS! ---------")
			fmt.Println("-------------------------------------")		
		case opcion == 7:
			sql.TestConsumo()
			fmt.Println("-------------------------------------")
			fmt.Println("------- ¡CONSUMOS  AGREGADOS! -------")
			fmt.Println("-------------------------------------")
		case opcion == 8:
			sql.GenerarResumenesPrincipal()
			fmt.Println("-------------------------------------")
			fmt.Println("------- ¡RESUMENES GENERADOS! -------")
			fmt.Println("-------------------------------------")
		
		case opcion == 9:
			sql.MostrarTablas()
			fmt.Println("-------------------------------------")
			fmt.Println("-------- ¡TABLAS  MOSTRADAS! --------")
			fmt.Println("-------------------------------------")
		case opcion == 10:
			sql.MostrarTodosDatos()
			fmt.Println("-------------------------------------")
			fmt.Println("--------- ¡DATOS MOSTRADOS! ---------")
			fmt.Println("-------------------------------------")
		case opcion == 11:
			sql.BorrarFKs()
			sql.BorrarPKs()
			fmt.Println("-------------------------------------")
			fmt.Println("---------- ¡KEYS BORRADAS! ----------")
			fmt.Println("-------------------------------------")
		case opcion == 12:
			nosql.CargaDatosNoDB()
			fmt.Println("-------------------------------------")
			fmt.Println("------ ¡DATOS NO-SQL CARGADOS! ------")
			fmt.Println("-------------------------------------")
		default:
			fmt.Println("Ingrese una opcion valida")
		}
	}
}
