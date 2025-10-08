package main

import (
	"github.com/dabraga/goexpert-ratelimiter/config"
	"github.com/dabraga/goexpert-ratelimiter/router"
)

func main() {
	config.Init()
	router.Init()
}
