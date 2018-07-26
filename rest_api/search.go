package rest_api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/LeyouHong/telo_demo/config"
	"log"
	"gitee.com/johng/gf/g/os/grpool"
	"sync"
)

func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m config.Inputs
	b, _ := ioutil.ReadAll(r.Body)
	//log.Println(string(b))
	json.Unmarshal(b, &m)

	res := process2(m)
	log.Println(string(res))
	w.Write(res)
}

func process1(m config.Inputs) []byte {
	for _, input := range m.Names {
			val, _ := config.Client.Classify(input)
			log.Println(val)
			config.Outputs.WriteMap(input, val)

	}


	jsonString, err := json.Marshal(config.Outputs.MAP)
	if err != nil {
		return []byte("can't find anything")
	}

	return jsonString
}


func process2(m config.Inputs) []byte {
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
			val, err := config.Client.Classify(v)
			if err != nil {
				log.Println(err)
			}

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
