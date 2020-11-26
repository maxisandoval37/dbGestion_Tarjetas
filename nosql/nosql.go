package nosql

import (
	"encoding/json"
	"log"
	"strconv"
	bolt "github.com/coreos/bbolt"
)

type Cliente struct {
	Nrocliente int
	Nombre     string
	Apellido   string
	Domicilio  string
	Telefono   string
}

type Tarjeta struct {
	Nrotarjeta   string
	Nrocliente   int
	Validadesde  string
	Validahasta  string
	Codseguridad string
	Limitecompra int
	Estado       string
}

type Comercio struct {
	Nrocomercio  int
	Nombre       string
	Domicilio    string
	Codigopostal string
	Telefono     string
}

type Compra struct {
	Nrooperacion int
	Nrotarjeta   string
	Nrocomercio  int
	Fecha        string
	Monto        int
	Pagado       bool
}

var boltdb *bolt.DB
var err error

func dbConnection() {
	boltdb, err = bolt.Open("./no-sql/bolt_database/algo.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}