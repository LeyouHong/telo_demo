package rest_api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/LeyouHong/telo_demo/config"
	"sync"
	"gitee.com/johng/gf/g/os/grpool"
	"log"
)

func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m config.Inputs
	b, _ := ioutil.ReadAll(r.Body)
	//log.Println(string(b))
	json.Unmarshal(b, &m)

	res := process(m)
	log.Println(string(res))
	w.Write(res)
}

func process(m config.Inputs) []byte {
	size := len(m.Names)
	if size > 10000 {
		size = 10000
	}

	grpool.SetSize(size)

	wg := sync.WaitGroup{}
	for _, input := range m.Names {
		wg.Add(1)
		v := input
		go grpool.Add(func() {
			//log.Println(v)
			val := config.Client.Get(v)
			config.Outputs.WriteMap(v, val)
			wg.Done()
		})
	}
	wg.Wait()

	jsonString, err := json.Marshal(config.Outputs.MAP)
	if err != nil {
		return []byte("can't find anything")
	}

	return jsonString
}
