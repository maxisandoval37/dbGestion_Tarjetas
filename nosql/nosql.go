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
	//defer db.Close()
}

func cargarCliente(nro_cliente int, nombre string, apellido string, domicilio string, telefono string) {
	cliente := Cliente{nro_cliente, nombre, apellido, domicilio, telefono}
	data, err := json.Marshal(cliente)
	if err != nil {
		log.Fatal(err)
	}

	CreateUpdate(db, "Cliente", []byte(strconv.Itoa(cliente.Nrocliente)), data)
}

func cargarTarjeta(nro_tarjeta string, nro_cliente int, valida_desde string, valida_hasta string, codigo_seguridad string, limite_compra int, estado string) {
	tarjeta := Tarjeta{nro_tarjeta, nro_cliente, valida_desde, valida_hasta, codigo_seguridad, limite_compra, estado}
	data, err := json.Marshal(tarjeta)
	if err != nil {
		log.Fatal(err)
	}

	CreateUpdate(db, "Tarjeta", []byte(strconv.Itoa(tarjeta.Nrocliente)), data)
}

func cargarComercio(nro_comercio int, nombre string, domicilio string, codigo_postal string, telefono string) {
	comercio := Comercio{nro_comercio, nombre, domicilio, codigo_postal, telefono}
	data, err := json.Marshal(comercio)
	if err != nil {
		log.Fatal(err)
	}

	CreateUpdate(db, "Comercio", []byte(strconv.Itoa(comercio.Nrocomercio)), data)
}

func cargarCompra(nro_operacion int, nro_tarjeta string, nro_comercio int, fecha string, monto int, pagado bool) {
	compra := Compra{nro_operacion, nro_tarjeta, nro_comercio, fecha, monto, pagado}
	data, err := json.Marshal(compra)
	if err != nil {
		log.Fatal(err)
	}

	CreateUpdate(db, "Compra", []byte(strconv.Itoa(compra.Nrooperacion)), data)
}

func CreateUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) error {
    // abre transacción de escritura
    tx, err := db.Begin(true)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))

    err = b.Put(key, val)
    if err != nil {
        return err
    }

    // cierra transacción
    if err := tx.Commit(); err != nil {
        return err
    }

    return nil
}

func CargarDatos_nosql() {
	dbConnection()

	//cargarCliente()
	//cargarTarjeta()
	//cargarComercio()
	//cargarCompra()

	// tres de cada entidad segun el enunciado
}




