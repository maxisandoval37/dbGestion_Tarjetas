package nosql

import (
	"encoding/json"
	"log"
	"strconv"
	bolt "github.com/coreos/bbolt"
	"fmt"
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

func CargaDatosNoDB() {
	dbConnection()

	cargarCliente(81635249, "Juan", "Perez", "Corrientes 159", "271542570123")
	cargarCliente(51362874, "Nahuel", "Rodriguez", "Calle 24 1235", "284529192033")
	cargarCliente(64723591, "Matias Adrian", "Perez", "Río Cuarto 2020", "221651370676")

	cargarTarjeta("4382420954476737", 81635249, "201205", "201306", "9184", 50, "vigente")
	cargarTarjeta("7229894669781604", 51362874, "200510", "200701", "4894", 10, "vigente")
	cargarTarjeta("6924033286851784", 64723591, "201703", "202001", "8850", 14, "vigente")

	cargarComercio(61351, "Fiambreria luna", "Av L Quaranta 7091", "T6863CSD", "339534886414")
	cargarComercio(79751, "Panaderia maxi", "Av R S Ortiz 160", "M3138MMM", "267532588765")
	cargarComercio(51249, "Grido Helado", "H Yrigoyen 1659", "E0179VAG", "361518160053")

	cargarCompra(1, "4382420954476737", 61351, "2020-04-25 17:50:42", 1230.00, true)
	cargarCompra(2, "4382420954476737", 79751, "2020-04-25 18:30:12", 800.00, true)
	cargarCompra(3, "4382420954476737", 51249, "2020-05-01 16:03:33", 500.00, true)
	cargarCompra(4, "7229894669781604", 61351, "2020-05-02 16:30:02", 700.00, true)
	cargarCompra(5, "7229894669781604", 79751, "2020-05-02 17:22:29", 5000.00, true)
	cargarCompra(6, "7229894669781604", 51249, "2020-05-02 17:59:13", 950.00, true)
	cargarCompra(7, "6924033286851784", 61351, "2020-05-03 09:42:59", 68.00, true)
	cargarCompra(8, "6924033286851784", 51249, "2020-05-04 03:10:01", 8000.00, false)
	cargarCompra(9, "6924033286851784", 51249, "2020-05-04 03:25:34", 8000.00, false)
}

func cargarCliente(nro_cliente int, nombre string, apellido string, domicilio string, telefono string) {
	cliente := Cliente{nro_cliente, nombre, apellido, domicilio, telefono}
	data, err := json.Marshal(cliente)
	if err != nil {
		log.Fatal(err)
	}

	CreateUpdate(boltdb, "Cliente", []byte(strconv.Itoa(cliente.Nrocliente)), data)

	resultado, err := ReadUnique(boltdb, "Cliente", []byte(strconv.Itoa(cliente.Nrocliente)))
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("%s\n", resultado)
}

func cargarTarjeta(nro_tarjeta string, nro_cliente int, valida_desde string, valida_hasta string, codigo_seguridad string, limite_compra int, estado string) {
	tarjeta := Tarjeta{nro_tarjeta, nro_cliente, valida_desde, valida_hasta, codigo_seguridad, limite_compra, estado}
	data, err := json.Marshal(tarjeta)
	if err != nil {
		log.Fatal(err)
	}

	CreateUpdate(boltdb, "Tarjeta", []byte(strconv.Itoa(tarjeta.Nrocliente)), data)

	resultado, err := ReadUnique(boltdb, "Tarjeta", []byte(strconv.Itoa(tarjeta.Nrocliente)))
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("%s\n", resultado)
}

func cargarComercio(nro_comercio int, nombre string, domicilio string, codigo_postal string, telefono string) {
	comercio := Comercio{nro_comercio, nombre, domicilio, codigo_postal, telefono}
	data, err := json.Marshal(comercio)
	if err != nil {
		log.Fatal(err)
	}

	CreateUpdate(boltdb, "Comercio", []byte(strconv.Itoa(comercio.Nrocomercio)), data)
	
	resultado, err := ReadUnique(boltdb, "Comercio", []byte(strconv.Itoa(comercio.Nrocomercio)))
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("%s\n", resultado)
}

func cargarCompra(nro_operacion int, nro_tarjeta string, nro_comercio int, fecha string, monto int, pagado bool) {
	compra := Compra{nro_operacion, nro_tarjeta, nro_comercio, fecha, monto, pagado}
	data, err := json.Marshal(compra)
	if err != nil {
		log.Fatal(err)
	}

	CreateUpdate(boltdb, "Compra", []byte(strconv.Itoa(compra.Nrooperacion)), data)
	
	resultado, err := ReadUnique(boltdb, "Compra", []byte(strconv.Itoa(compra.Nrooperacion)))
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("%s\n", resultado)
}

func CreateUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) error {
    // abre transacción de escritura
    tx, err := boltdb.Begin(true)
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

func ReadUnique(db *bolt.DB, bucketName string, key []byte) ([]byte, error) {
    var buf []byte

    // abre una transacción de lectura
    err := db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucketName))
        buf = b.Get(key)
        return nil
    })

    return buf, err
}
