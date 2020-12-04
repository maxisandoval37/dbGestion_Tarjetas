package sql

import (
	"log"
)

func spAutorizarCompra() {	
	_, err = db.Query(
		`CREATE OR REPLACE FUNCTION autorizarcompra(_nrotarjeta char(16), _codseguridad char(4), _nrocomercio int, _monto decimal(7,2)) returns bool as $$
		 DECLARE
			totalPendiente decimal(8,2);
			montoMaximo decimal(8,2);
			fechaVenceTarjeta int;
			fechaVence date;
		 BEGIN
			INSERT INTO compra(nrotarjeta, nrocomercio, fecha, monto, pagado) VALUES(_nrotarjeta, _nrocomercio, current_timestamp, _monto,False);
			return True;
		 END;
	$$ LANGUAGE PLPGSQL;`)
	
	if err != nil {
		log.Fatal(err)
	}
}

func spAgregarConsumoTESTaux(nrotarjeta string, codseg string, nrocomer string, monto string) {
	_, err = db.Query(
		`CREATE OR REPLACE FUNCTION agregar_consumo() returns void as $$
		 BEGIN
		
		 PERFORM autorizarcompra( CAST(`+nrotarjeta+` as char(16)), CAST(`+codseg+` as char(4)), CAST(`+nrocomer+` as int), CAST(`+monto+` as decimal(7,2)) );
		 --return new;
		 END;
	$$ LANGUAGE PLPGSQL;`)
	if err != nil {
		log.Fatal(err)
	}
}

func AgregarConsumo(){
	spAutorizarCompra()
	spAgregarRechazo()
	spAgregarConsumoTESTaux("4382420954476737","9184","95169","10")//valido
}

func spAgregarRechazo() {
	_, err = db.Query(
		`
		CREATE OR REPLACE FUNCTION agregarrechazo(_nrotarjeta char(16),_nrocomercio int, _fecha timestamp,_monto decimal(7,2),_motivo text) returns void as $$
		DECLARE
			numerorechazo int;
		BEGIN
			INSERT INTO rechazo(nrotarjeta, nrocomercio, fecha, monto, motivo) VALUES( _nrotarjeta, _nrocomercio, current_timestamp, _monto, _motivo)
			RETURNING nrorechazo INTO numerorechazo;
		END;

	$$ LANGUAGE PLPGSQL;`)
	if err != nil {
		log.Fatal(err)
	}
}
