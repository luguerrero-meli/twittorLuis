package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/luguerrero-meli/twittorLuis/bd"
	"github.com/luguerrero-meli/twittorLuis/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var message models.GraboTweet
	err := json.NewDecoder(r.Body).Decode(&message)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Error al crear el tweet reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Error al crear el tweet.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
