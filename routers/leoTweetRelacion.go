package routers

import (
	"encoding/json"
	"github.com/luguerrero-meli/twittorLuis/bd"

	"net/http"
	"strconv"
)

func LeoTweetRelacion (w http.ResponseWriter, r *http.Request){

	if len (r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro page", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro page mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, status := bd.LeoTweetSeguidores(IDUsuario, page)
	if status == false {
		http.Error(w,"Error al leer los tweets de los seguidores", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
