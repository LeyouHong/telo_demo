package rest_api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/telo_demo/config"
	"log"
)

func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m config.Inputs
	b, _ := ioutil.ReadAll(r.Body)
	log.Println(string(b))
	json.Unmarshal(b, &m)

	res := process(m)
	w.Write(res)
}

func process(m config.Inputs) []byte {
	for _, input := range m.Names {
		val := config.Client.Get(input)
		config.Outputs.WriteMap(input, val)
	}

	jsonString, err := json.Marshal(config.Outputs.MAP)
	if err != nil {
		return []byte("can't find anything")
	}

	return jsonString
}
