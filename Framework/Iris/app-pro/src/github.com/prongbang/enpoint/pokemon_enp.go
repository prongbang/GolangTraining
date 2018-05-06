package enpoint

import (
	"net/http"
	"strconv"
	"math"
	"github.com/kataras/iris"
	"github.com/prongbang/service"
	"github.com/prongbang/utils"
	"fmt"
)

func PokemonHandle(ctx *iris.Context) {

	token := ctx.RequestHeader("Authorization")


	fmt.Println("token", token)

	pageNo, _ := strconv.Atoi(ctx.Param("pageNo"))
	pageSize, _ := strconv.Atoi(ctx.Param("pageSize"))

	var pageCount = utils.Round(math.Ceil(float64(service.PokemonCountService() / pageSize)))

	pageNo = utils.CalcPageNo(pageNo, pageCount)
	var startRow = utils.CalcStartRow(pageNo, pageSize)
	var endRow = utils.CalcEndRow(startRow, pageSize)

	ctx.JSON(http.StatusOK, iris.Map{
		"pokemons": service.PokemonService(startRow, endRow),
		"pageNo": pageNo,
		"pageCount": pageCount,
	})

}
