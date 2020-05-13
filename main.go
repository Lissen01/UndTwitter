package main

import (
	"log"

	"github.com/Lissen01/UndTwitter/bd"
	"github.com/Lissen01/UndTwitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin Conexion a la BD")
		return
	}
	handlers.Manejadores()
}
