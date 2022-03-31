package main

import (
	"log"

	"krizz.org/sleepi/web"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}
}

func run() error {
	return web.Serve()
}
