package bd

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/luguerrero-meli/twittorLuis/models"
)

/*IntentoLogin es utilizado para la validación del usuario*/
func IntentoLogin(email, password string) (models.Usuario, bool) {
	usr, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usr, false
	}

	passwordByte := []byte(password)
	passwordBD := []byte(usr.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordByte)
	if err != nil {
		return usr, false
	}

	return usr, true

}
