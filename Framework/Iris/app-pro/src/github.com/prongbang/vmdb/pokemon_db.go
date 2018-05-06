package vmdb

import "github.com/prongbang/model"

// Virtual Database
func PokemonDB() []model.Pokemon {

	name := []string{
		"Bulbasaur",
		"Ivysaur",
		"Venusaur",
		"Charmander",
		"Charmeleon",
		"Charizard",
		"Squirtle",
		"Wartortle",
		"Blastoise",
		"Caterpie",
		"Metapod",
		"Butterfree",
		"Weedle",
		"Kakuna",
		"Beedrill",
		"Pidgey",
		"Pidgeotto",
		"Pidgeot",
		"Rattata",
		"Raticate",
		"Spearow",
		"Fearow",
		"Ekans",
		"Arbok",
		"Pikachu",
		"Raichu",
		"Sandshrew",
	}
	size := len(name)
	var pokemons = make([]model.Pokemon, size)

	for i := 0; i < size; i++ {
		pokemons[i].Id = i + 1
		pokemons[i].Name = name[i]
	}

	return pokemons
}
