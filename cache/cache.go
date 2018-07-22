package cache

import (
	"log"
	"github.com/telo_demo/config"
)

func CacheBuild() {
	log.Print("cache building ...")
	config.Client.Set("Name", "Leyou")
	if config.Client.Get("Name") != "Leyou" {
		log.Println("CacheBuild error")
	}
}
