# Hisui

Search and display information about a Pokémon directly from the terminal.

All Pokémon information is retrieved from [PokéAPI](https://pokeapi.co/).

For now, this application will only run locally on your device. In the future, I would like to provide an easier method
for running this application.

Firstly, you will need to [install Go](https://go.dev/doc/install) on your device. Secondly, you will need to download
the files from this repository to your device, I recommend you use Git by running
`git clone git@github.com:kerrance/go-hisui.git`.

# How to run the application

Once Go is installed, you can boot up your Terminal and navigate to the directory folder. You will then be able to run
`go run main.go` to start the application in your terminal window.

# How to test the application

Following the instructions above to run the application, from the root directory, you can enter `go test ./...` to run
all the automated tests bundled within the application.

# Credits

- [PokéAPI](https://pokeapi.co/)
- [MicroDex](https://github.com/404a10/MicroDex) for inspiration
- [Joe Sweeny](https://github.com/joesweeny), [Thomas Hockaday](https://github.com/thomashockaday),
[Frazer Lillie](https://github.com/fillie), [Neil Davies](https://github.com/NeilDavies92) and
[Sam Langley](https://github.com/samlangley1) for their guidance whilst I have been learning Go
