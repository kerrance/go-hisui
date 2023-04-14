package main

import (
	"bufio"
	"fmt"
	"github.com/kerrance/go-hisui/app"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Search for a Pokémon by name, or National Pokédex number:")

	consoleInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal("An unexpected error occurred:", err)
	}
	enteredPokemonNameOrPokedexNumber := strings.TrimSuffix(consoleInput, "\n")

	sanitisedPokemonName := app.SanitisePokemonNameForSearch(enteredPokemonNameOrPokedexNumber)
	app.Show(app.FindPokemon(sanitisedPokemonName))
}
