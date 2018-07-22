package main

import (
	"github.com/telo_demo/httpserver"
	"github.com/telo_demo/cache"
)

func main() {
	cache.CacheBuild()
	httpserver.Start()
}
