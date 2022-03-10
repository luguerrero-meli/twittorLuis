package routers

import (
	"encoding/json"
	"net/http"

	"github.com/luguerrero-meli/twittorLuis/bd"
	"github.com/luguerrero-meli/twittorLuis/models"
)

/*ModificarRegistro modifica perfil de usuario*/
func ModificarRegistro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificarRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro. Reintente nuevamente."+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modifica el registro del usuario.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
