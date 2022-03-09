package routers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/luguerrero-meli/twittorLuis/bd"
	"github.com/luguerrero-meli/twittorLuis/models"
)

/*Email valor del Email usado en todos los EndPoints. */
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usara en todos los Endpoints.*/
var IDUsuario string

/**ProcesoToken proceso token para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupoFacebook")
	claim := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claim, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claim, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claim.Email)
		if encontrado == true {
			Email = claim.Email
			IDUsuario = claim.ID.Hex()
		}
		return claim, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claim, false, string(""), errors.New("token invalido")
	}

	return claim, false, string(""), err

}
