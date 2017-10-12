package main

import (
	"faygo_about/router"
	"github.com/henrylee2cn/faygo"
)

func main() {
	router.Route(faygo.New("faygo_about"))
	faygo.Run()
}
