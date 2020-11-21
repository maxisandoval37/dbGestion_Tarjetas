package main

import (
	"fmt"
)

func main() {
	menuPrincipal()
}

var opcion int=0

func menuPrincipal() {

		
		fmt.Println("Digite una opcion")
		fmt.Println("1: Crear BD")
		

		fmt.Println("99: Salir ")

		fmt.Scanf("%d", &opcion)
	
	switch {
		case opcion == 1:
			fmt.Println("--------------------\n    ¡BD CREADA!\n--------------------")
		
		
		default:
			fmt.Println("Ingrese una opcion valida")
	}
	
}

