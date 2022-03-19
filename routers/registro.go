package routers

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"github.com/luguerrero-meli/twittorLuis/bd"
	"github.com/luguerrero-meli/twittorLuis/models"
)

/*Registro es la funcion de creacion en bd el registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	fmt.Println(t.Email)
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido.", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos seis caracteres.", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "El usuario ya existe.", 400)
		return
	}
	t.ID = primitive.NewObjectID()

	fmt.Print(t)
	_, status, err := bd.InsertaRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrrio un error al registrar el usuario."+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
