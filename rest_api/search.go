package rest_api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/telo_demo/config"
	"sync"
	"gitee.com/johng/gf/g/os/grpool"
)

func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m config.Inputs
	b, _ := ioutil.ReadAll(r.Body)
	//log.Println(string(b))
	json.Unmarshal(b, &m)

	res := process(m)
	w.Write(res)
}

func process(m config.Inputs) []byte {
	size := len(m.Names)
	if size > 1000 {
		size = 1000
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
