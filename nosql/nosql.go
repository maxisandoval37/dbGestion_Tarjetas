package nosql

import (
	"encoding/json"
	"log"
	"strconv"
	bolt "github.com/coreos/bbolt"
)

type cliente struct {
	nrocliente int
	nombre     string
	apellido   string
	domicilio  string
	telefono   string
}

type tarjeta struct {
	nrotarjeta   string
	nrocliente   int
	validadesde  string
	validahasta  string
	codseguridad string
	limitecompra int
	estado       string
}

type comercio struct {
	nrocomercio  int
	nombre       string
	domicilio    string
	codigopostal string
	telefono     string
}

type compra struct {
	nrooperacion int
	nrotarjeta   string
	nrocomercio  int
	fecha        string
	monto        int
	pagado       bool
}