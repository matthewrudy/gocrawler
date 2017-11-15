package main

import (
	"flag"
	"fmt"

	"github.com/matthewrudy/gocrawler"
)

var entrypoint = flag.String("entrypoint", "http://tomblomfield.com/", "entrypoint to crawl from")

func init() {
	flag.Parse()
}

func main() {
	crawler := gocrawler.New(*entrypoint)
	crawler.Crawl()
	fmt.Println(crawler)
}
