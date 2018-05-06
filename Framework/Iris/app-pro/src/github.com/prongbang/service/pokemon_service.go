package service

import (
	"github.com/prongbang/dao"
	"github.com/prongbang/model"
)

func PokemonCountService() int {

	return dao.PokemonCountDao()
}

func PokemonService(startRow int, endRow int) []model.Pokemon {

	return dao.PokemonDao(startRow, endRow)
}
