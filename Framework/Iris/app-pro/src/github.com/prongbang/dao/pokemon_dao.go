package dao

import (
	"github.com/prongbang/model"
	"github.com/prongbang/vmdb"
)

func PokemonCountDao() int {

	return len(vmdb.PokemonDB())
}

func PokemonDao(startRow int, endRow int) []model.Pokemon {

	var pokemons = make([]model.Pokemon, endRow - startRow + 1)

	data := vmdb.PokemonDB()
	count := PokemonCountDao();
	f := 0;
	for i := startRow - 1; i < endRow; i++ {
		if (i < count) {
			pokemons[f] = data[i];
			f++
		}
	}

	return pokemons
}

