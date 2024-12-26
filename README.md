# TuiCraft

TuiCraft is a terminal-based game built using the Go programming language. It leverages the power of the Bubble Tea framework to create a rich, interactive text user interface (TUI). This project is currently a work in progress.

## Overview

The game allows players to navigate through menus, create and load game profiles, and interact with various game elements such as items and combat encounters. The game state is managed and saved to JSON files to allow for persistent gameplay.

## Project Structure

```.gitignore
go.mod
go.sum
main.go
model_save.json
player_save.json
pkg/
    model.go
    styles.go
    update.go
    views.go
```

### Package: [`pkg`](pkg)

- [`pkg/model.go`](pkg/model.go): Contains the core game logic, including the definition of game types, initialization functions, and state management.
- [`pkg/styles.go`](pkg/styles.go): Defines the styles used in the TUI, leveraging the Lip Gloss library for styling.
- [`pkg/update.go`](pkg/update.go): Handles the update logic for the game, processing user inputs and updating the game state accordingly.
- [`pkg/views.go`](pkg/views.go): Contains the view logic, rendering the different screens and menus of the game.

## Dependencies

The project uses several external packages to build the TUI and manage game logic:

- [Bubble Tea](https://github.com/charmbracelet/bubbletea): A powerful, elegant, and simple framework for building terminal user interfaces.
- [Lip Gloss](https://github.com/charmbracelet/lipgloss): A library for styling terminal applications.
- [Ease](https://github.com/fogleman/ease): A library for easing functions, used for animations.
- [Colorful](https://github.com/lucasb-eyer/go-colorful): A library for color manipulation.

## Getting Started

To run the project, clone the repository and run the following commands:

```sh
go mod tidy
go run main.go
```

## Contributing
This project is a work in progress. It will probably never see updates but it was fun to work on.