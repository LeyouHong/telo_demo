package httpserver

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/telo_demo/rest_api"
)

func Start() {
	log.Println("Http server starting ...")
	router := mux.NewRouter()
	router.HandleFunc("/search/{name}", rest_api.Search).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
