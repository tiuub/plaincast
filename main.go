package main

import (
	"flag"

	"github.com/tiuub/plaincast/server"
)

func main() {
	flag.Parse()

	server.Serve()
}
