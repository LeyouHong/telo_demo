package cache

import (
	"log"
	"io/ioutil"
	"fmt"
	"github.com/LeyouHong/telo_demo/config"
)

func CacheBuild() {
	log.Print("cache building ...")
	business, err := ioutil.ReadFile("/tmp/business.txt")
	if err != nil {
		fmt.Print(err)
	}

	location, err := ioutil.ReadFile("/tmp/location.txt")
	if err != nil {
		fmt.Print(err)
	}

	person, err := ioutil.ReadFile("/tmp/person.txt")
	if err != nil {
		fmt.Print(err)
	}

	err = config.Client.Shd.Learn("business", string(business))
	if err == nil {
		log.Println("learn business data successful")
	}
	err = config.Client.Shd.Learn("location", string(location))
	if err == nil {
		log.Println("learn location data successful")
	}
	err = config.Client.Shd.Learn("person", string(person))
	if err == nil {
		log.Println("learn person data successful")
	}
}

