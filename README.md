# Discord Bot

Discord Bot is a modular boilerplate for a simple command-based discord bot that I have open-sourced from my private use.

## Features

Discord Bot is:
* Modular (easy to implement and remove "modules" or commands)
* Adjustable prefix
* Containerized (provided Dockerfile)

## Installation

`go mod download`

## How To Run

Create a `.env` file in the root directory which contains your Bot Token and optionally, a prefix

Default Webhook Embed formatting can be edited in the `pkg/helpers` package.

Run it in:
* A Container
* `go run cmd/main.go`

## Adding New Modules

New modular modules can be created and binded to a specific command.

### Creating the module

Create a `<module>.go` file in `pkg/modules` for your new module. It must fulfill the `Modules` interface defined in `pkg/modules/modules.go`. 

Note: Modules are acted only upon by text commands. For custom handlers, see below.



### Custom Handlers

Custom Handlers can be added to the bot in `cmd/main.go` via the `AddHandler()` method which is a wrapper for [bwmarrin's discordgo](https://github.com/bwmarrin/discordgo) implementation.

[AddHandler documentation](https://godoc.org/github.com/bwmarrin/discordgo#Session.AddHandler)
