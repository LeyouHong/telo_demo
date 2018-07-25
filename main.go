package main

import (
	"github.com/LeyouHong/telo_demo/httpserver"
	"github.com/LeyouHong/telo_demo/cache"
)

func main() {
	cache.CacheBuild()
	httpserver.Start()
}
