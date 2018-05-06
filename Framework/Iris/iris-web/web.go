package main

import (
"net/http"
"github.com/kataras/iris"
"fmt"
)

func main() {
    iris.Get("/hi", hi)
    iris.Listen(":8888")
}

func hi(ctx *iris.Context) {

	if origin := r.Header.Get("Origin"); origin != "" {
        w.Header().Set("Access-Control-Allow-Origin", origin)
        fmt.Println(origin)
    }
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")


	token := ctx.RequestHeader("Authorization")

	fmt.Println("token", token)

	ctx.JSON(http.StatusOK, iris.Map{
		"data": "iris",
	})

}