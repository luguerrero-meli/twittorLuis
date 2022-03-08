package main

import (
	"log"

	"github.com/luguerrero-meli/twittorLuis/bd"
	"github.com/luguerrero-meli/twittorLuis/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
