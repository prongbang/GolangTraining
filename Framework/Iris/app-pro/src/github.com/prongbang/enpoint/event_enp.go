package enpoint

import (
	"github.com/kataras/iris"
	"fmt"
)

func EventHandle(ctx *iris.Context)  {

	token := ctx.RequestHeader("Authorization")

	fmt.Println("token", token)

	ctx.JSON(iris.StatusOK, struct { Name string }{ Name: "iris" })
}
