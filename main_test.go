package main

import (
	"testing"
	"log"
	"os"
	"gitee.com/johng/gf/g/os/grpool"
	"bufio"
	"sync"
	"io"
	"strings"
	"github.com/telo_demo/redis"
)

func TestInsertData(t *testing.T) {
	log.Println("insertData ...")
	Client := redis.NewRedis("127.0.0.1:6379")
	insertData(Client)
}

func insertData(cli *redis.Redis) {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Println("Open data.txt failed")
	}

	defer file.Close()

	grpool.SetSize(100)
	br := bufio.NewReader(file)
	reading := true
	wg := sync.WaitGroup{}
	count := 0
	for reading {
		for i := 0; i < 100; i++ {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				reading = false
				break
			}
			kv := strings.Split(string(a), " ")

			if len(kv) == 2 {
				wg.Add(1)
				v := count
				go grpool.Add(func() {
					log.Println(v)
					cli.Set(kv[0], kv[1])
					wg.Done()
				})
			}
			count++
		}
		wg.Wait()
	}
}