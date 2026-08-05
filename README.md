# 🧢 Pokedex CLI

A command-line Pokédex built with Go that interacts with the [PokéAPI](https://pokeapi.co/). Explore the Pokémon world, discover Pokémon in different areas, catch them, inspect their stats, and build your own Pokédex—all from your terminal.

## Features

- 🌍 Browse Pokémon location areas
- ⏪ Navigate forward and backward through map pages
- 🔎 Explore Pokémon found in a specific area
- ⚾ Catch Pokémon with a randomized catch mechanic
- 📖 Inspect Pokémon you've caught
- 📚 View your personal Pokédex
- ⚡ In-memory API caching for faster requests
- ❄️ Reproducible development environment using Nix Flakes

---

## Requirements

- Go 1.22+ (or the version specified in `go.mod`)
- Internet connection (uses the public PokéAPI)

### Optional

- Nix with Flakes enabled

---

## Installation

### Clone the repository

```bash
git clone https://github.com/NinjaCrusader/pokedexcli.git
cd pokedexcli
```

### Run with Go

```bash
go run .
```

### Build

```bash
go build -o pokedex
./pokedex
```

---

## Using Nix

If you have Nix installed, enter the development shell:

```bash
nix develop
```

Then run:

```bash
go run .
```

---

# Commands

| Command | Description |
|----------|-------------|
| `help` | Display available commands |
| `exit` | Exit the application |
| `map` | View the next 20 Pokémon location areas |
| `mapb` | View the previous 20 location areas |
| `explore <area>` | Show Pokémon that can be found in a location area |
| `catch <pokemon>` | Attempt to catch a Pokémon |
| `inspect <pokemon>` | Display information about a caught Pokémon |
| `pokedex` | Display all Pokémon you've caught |

---

## Example Session

```text
Pokedex > map

canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore canalave-city-area

Found Pokemon:
 - tentacool
 - wingull
 - pelipper

Pokedex > catch tentacool

Throwing a Pokeball at tentacool...
tentacool was caught!

Pokedex > inspect tentacool

Name: tentacool
Height: 9
Weight: 455

Stats:
  - hp: 40
  - attack: 40
  - defense: 35
  ...

Types:
  - water
  - poison

Pokedex > pokedex

 - tentacool
```

---

## Project Structure

```
.
├── internal/        # API helpers, cache, and models
├── main.go          # Application entry point
├── repl.go          # Interactive REPL and commands
├── go.mod
├── flake.nix        # Nix development environment
└── README.md
```

---

## How It Works

The application launches an interactive REPL (Read-Eval-Print Loop) where users can enter commands. Data is fetched from the PokéAPI and cached in memory to reduce repeated API requests. Successfully caught Pokémon are stored in memory for the duration of the session, allowing players to inspect them or view their growing Pokédex.

---

## Technologies

- Go
- PokéAPI
- Nix Flakes
- Standard Library

---

## Future Improvements

- Save Pokédex between sessions
- Better catch mechanics
- Search Pokémon by type
- Battle simulator
- Colored terminal output
- Command history and autocomplete

---

## Contributors

- [Jerold Pemberton (NinjaCrusader)](https://github.com/NinjaCrusader)
- [Tai Fong](https://github.com/TaiFong)

---

## License

This project is intended for learning and educational purposes.
