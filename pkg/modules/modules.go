package modules

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// Module Interface
type Module interface {
	// Run takes in a string of arguments of dynamic length.
	// Each module responsible for sending their own discord Objects.
	// Returns:
	// *discordgo.MessageEmbed: Empty if nothing to return
	// interface{}: any state to be saved with the messageID
	Run(m *discordgo.MessageCreate, s *discordgo.Session, args []string) error
}

// Modules gives access to the global store for all commands
type Modules struct {
	// messageID -> []Structs
	MessageStore map[string]interface{}
	// CommandName -> CommandModule
	ModuleStore  map[string]Module
	HandlerStore []interface{}
}

// NewModules returns a Modules struct
func NewModules() *Modules {
	moduleStore := make(map[string]Module)
	msgStore := make(map[string]interface{})
	return &Modules{
		MessageStore: msgStore,
		ModuleStore:  moduleStore,
	}
}

// AddModule registers a module to the discord bot
func (m *Modules) AddModule(commandName string, commandModule Module) {
	if _, ok := m.ModuleStore[commandName]; ok {
		log.Fatalf("Module with name: %s already exists", commandName)
	}

	m.ModuleStore[commandName] = commandModule
}

// AddHandler registers a handler / listener to the discord bot.
func (m *Modules) AddHandler(handler interface{}) {
	m.HandlerStore = append(m.HandlerStore, handler)
}

const (
	nameStockx  = "stockx"
	nameHelp    = "help"
	nameConvert = "exchange"
)
