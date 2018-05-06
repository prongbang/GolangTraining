package main

import (
	"github.com/kataras/iris"
	"github.com/prongbang/enpoint"
)

func main() {

	iris.Get("/api/event", enpoint.EventHandle)

	iris.Get("/api/pokemon/:pageNo/:pageSize", enpoint.PokemonHandle)

	iris.Listen(":8080")
}
