package sql

import (
	"log"
)

func Triggers() {
	//Compra rechazada
	_, err := db.Exec(
	`CREATE OR REPLACE FUNCTION func_trigger_compra_rechazada() RETURNS TRIGGER AS $$
	BEGIN
		INSERT INTO alerta VALUES(DEFAULT, new.nrotarjeta, now(), new.nrorechazo, 0, 'rechazo');
		RETURN new;
	END;
	$$ LANGUAGE plpgsql;

	CREATE TRIGGER trigger_compra_rechazada 
	AFTER 
	INSERT ON rechazo
	FOR EACH ROW
	EXECUTE PROCEDURE func_trigger_compra_rechazada(); `)
	
	if err != nil {
		log.Fatal(err)
	}

	// Caso compra un minuto
	_, err = db.Exec(
	`CREATE OR REPLACE FUNCTION func_trigger_compra_unminuto() RETURNS TRIGGER AS $$
	DECLARE
		anterior record;
		codpostal_nueva record;
		codpostal_anterior record;
		
	BEGIN

		SELECT * INTO anterior FROM compra WHERE new.nrotarjeta = compra.nrotarjeta and 
		compra.nrooperacion = (SELECT MAX (nrooperacion) FROM compra);

		SELECT codigopostal INTO codpostal_nueva FROM comercio WHERE new.nrocomercio = comercio.nrocomercio;
		SELECT codigopostal INTO codpostal_anterior FROM comercio WHERE anterior.nrocomercio = comercio.nrocomercio;

		IF FOUND THEN
			IF (((new.fecha - anterior.fecha) < INTERVAL '1 min')) and (codpostal_nueva = codpostal_anterior) THEN
				INSERT INTO alerta VALUES(DEFAULT, new.nrotarjeta, now(), NULL, 1, 'compra 1min');
			END IF;
		END IF;

		RETURN new;

	END;

	$$ LANGUAGE plpgsql;

	CREATE TRIGGER trigger_compra_unminuto
	BEFORE
	INSERT ON compra
	FOR EACH ROW
	EXECUTE PROCEDURE func_trigger_compra_unminuto();`)
	
	if err != nil {
		log.Fatal(err)
	}

	// Caso compra 5 minutos
	_, err = db.Exec(
	`CREATE OR REPLACE FUNCTION func_trigger_compra_cincominutos() RETURNS TRIGGER AS $$
	DECLARE

		anterior record;
		codpostal_nueva record;
		codpostal_anterior record;
		
	BEGIN

		SELECT * INTO anterior FROM compra WHERE new.nrotarjeta = compra.nrotarjeta and 
		compra.nrooperacion = (SELECT MAX (nrooperacion) FROM compra);

		SELECT codigopostal INTO codpostal_nueva FROM comercio WHERE new.nrocomercio = comercio.nrocomercio;
		SELECT codigopostal INTO codpostal_anterior FROM comercio WHERE anterior.nrocomercio = comercio.nrocomercio;

		IF FOUND THEN
			IF (((new.fecha - anterior.fecha) < INTERVAL '5 mins')) and (codpostal_nueva != codpostal_anterior) THEN
				INSERT INTO alerta VALUES(DEFAULT, new.nrotarjeta, now(), NULL, 5, 'compra 5min');
			END IF;
		END IF;

		RETURN new;

	END;

	$$ LANGUAGE plpgsql;

	CREATE TRIGGER trigger_compra_cincominutos
	BEFORE
	INSERT ON compra
	FOR EACH ROW
	EXECUTE PROCEDURE func_trigger_compra_cincominutos();`)
	
	if err != nil {
		log.Fatal(err)
	}

	//Caso dos rechazos exceso
	_, err = db.Exec(
	`CREATE OR REPLACE FUNCTION func_trigger_compra_dos_excesos() RETURNS TRIGGER AS $$
	DECLARE

		rechazos INT;

	BEGIN
		IF new.motivo = '?Supera límite de tarjeta' THEN
			SELECT COUNT (*) INTO rechazos FROM rechazo WHERE new.nrotarjeta = rechazo.nrotarjeta and 
															new.fecha::date = rechazo.fecha::date
															and rechazo.motivo = '?Supera límite de tarjeta';
			IF rechazos = 2 THEN
				UPDATE tarjeta SET estado = 'suspendida' WHERE new.nrotarjeta = tarjeta.nrotarjeta;
				
				INSERT INTO alerta VALUES(DEFAULT, new.nrotarjeta, now(), new.nrorechazo, 32, 'limite');
			END IF;
		END IF;

		RETURN new;
	END;

	$$ LANGUAGE plpgsql;

	CREATE TRIGGER trigger_compra_dosrechazos
	AFTER
	INSERT ON rechazo
	FOR EACH ROW
	EXECUTE PROCEDURE func_trigger_compra_dos_excesos();`)
	
	if err != nil {
		log.Fatal(err)
	}
}
