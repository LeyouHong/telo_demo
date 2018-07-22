package rest_api

import (
	"net/http"
	"github.com/gorilla/mux"
)

func Search(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	word := params["name"]

	w.Write([]byte(word))
}
