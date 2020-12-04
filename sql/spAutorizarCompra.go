package sql

import (
	"log"
	"strconv"
)

type consumoObj struct {
	nrotarjeta   string
	codseguridad string
	nrocomercio  int
	monto        float64
}

var nroperacionCompra int = 0
var nrorechazo int = 0

func spAutorizarCompra() {	
	
	nroperacionCompra=nroperacionCompra+1;
	nrorechazo=nrorechazo+1;
	
	_, err = db.Query(
		`create or replace function autorizar_compra(n_tarjeta tarjeta.nrotarjeta%type,
						codigo tarjeta.codseguridad%type,
						n_comercio comercio.nrocomercio%type,
						monto_abonado compra.monto%type) returns boolean as $$
	DECLARE
		tarjeta_encontrada record;  
		compras_pendientes_de_pago compra.monto%type;
		
	BEGIN
		SELECT * INTO tarjeta_encontrada from tarjeta t where n_tarjeta = t.nrotarjeta; 
		compras_pendientes_de_pago := (select sum (monto) from compra c WHERE c.nrotarjeta = n_tarjeta and c.pagado = false);

		if not found  then           
			INSERT INTO rechazo (nrorechazo,nrotarjeta, nrocomercio, fecha, monto, motivo) VALUES (`+strconv.Itoa(nrorechazo)+`,n_tarjeta, n_comercio, current_timestamp, monto_abonado, 'Tarjeta no valida');
			return false;	
		
		elsif tarjeta_encontrada.codseguridad != codigo then
			 INSERT INTO rechazo (nrorechazo,nrotarjeta, nrocomercio, fecha, monto, motivo) VALUES (`+strconv.Itoa(nrorechazo)+`,n_tarjeta, n_comercio, current_timestamp, monto_abonado, 'Codigo de seguridad no valido');    
			 return false;

		elsif tarjeta_encontrada.estado = 'suspendida' then
			 INSERT INTO rechazo (nrorechazo,nrotarjeta, nrocomercio, fecha, monto, motivo) VALUES (`+strconv.Itoa(nrorechazo)+`,n_tarjeta, n_comercio, current_timestamp, monto_abonado, 'Tarjeta suspendida');   
			 return false;
			 
		elsif tarjeta_encontrada.estado = 'anulada' then
			INSERT INTO rechazo (nrorechazo,nrotarjeta, nrocomercio, fecha, monto, motivo) VALUES (`+strconv.Itoa(nrorechazo)+`,n_tarjeta, n_comercio, current_timestamp, monto_abonado, 'Plazo de vigencia expirado');
			return false;
						
		elsif tarjeta_encontrada.limitecompra < (compras_pendientes_de_pago + monto_abonado)  then
			INSERT INTO rechazo (nrorechazo,nrotarjeta, nrocomercio, fecha, monto, motivo) VALUES (`+strconv.Itoa(nrorechazo)+`,n_tarjeta, n_comercio, current_timestamp, monto_abonado, 'Supera límite de tarjeta');
			return false;
	
		else
			INSERT INTO compra (nrooperacion, nrotarjeta, nrocomercio, fecha, monto, pagado) VALUES (`+strconv.Itoa(nroperacionCompra)+`,n_tarjeta, n_comercio, current_timestamp, monto_abonado, false);
			return true;
		
		end if;
	end;
$$ language plpgsql;`)
	
	if err != nil {
		log.Fatal(err)
	}

}

func AutorizarCompra() {
		var c consumoObj
		var r bool
	
		row, err := db.Query(`select * from consumo`)
		
		if err != nil {
			log.Fatal(err)
		}
		
		defer row.Close()
		

		for row.Next() {
			if err = row.Scan(&c.nrotarjeta, &c.codseguridad, &c.nrocomercio, &c.monto); err != nil {	
				log.Fatal(err)
			}
		}
		
		row, err = db.Query(`select autorizar_compra($1::char(16), $2::char(4), $3::int, $4::decimal(7,2));`, c.nrotarjeta, c.codseguridad, c.nrocomercio, c.monto)
		if err != nil {
			log.Fatal(err)
		}
		
		for row.Next() {
			if err = row.Scan(&r); err != nil {
				log.Fatal(err)
			}
		}
		
		if err != nil {
			log.Fatal(err)
		}
}

func TestConsumo(){
	spAutorizarCompra()
	AutorizarCompra()
}

