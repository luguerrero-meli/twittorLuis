package middlew

import (
	"net/http"

	"github.com/luguerrero-meli/twittorLuis/bd"
)

/*ChequeoBD es el middlewar que hace la validacion de conexion hacia la bd y envio de control al routers*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdido con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
